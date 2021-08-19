// Package env is a lightweight package for loading OS environment variables into structs in Go projects.
package env

import (
	"errors"
	"github.com/golobby/cast"
	"os"
	"reflect"
	"unsafe"
)

// Load retrieves OS environment variables into the given struct.
// It gets the pointer of a struct that is going to holds the variables.
// The struct fields must have an `env` tag that determines the related OS environment variable.
// Tag example: `env:"APP_NAME"`
func Load(structure interface{}) error {
	inputType := reflect.TypeOf(structure)
	if inputType != nil {
		if inputType.Kind() == reflect.Ptr {
			if inputType.Elem().Kind() == reflect.Struct {
				return loadStruct(reflect.ValueOf(structure).Elem())
			}
		}
	}

	return errors.New("env: invalid structure")
}

// loadStruct retrieves OS environment variables into the given struct reflected value.
func loadStruct(s reflect.Value) error {
	for i := 0; i < s.NumField(); i++ {
		if t, exist := s.Type().Field(i).Tag.Lookup("env"); exist {
			v, err := cast.FromString(os.Getenv(t), s.Type().Field(i).Type.Name())
			if err != nil {
				return err
			}

			ptr := reflect.NewAt(s.Field(i).Type(), unsafe.Pointer(s.Field(i).UnsafeAddr())).Elem()
			ptr.Set(reflect.ValueOf(v))
		} else if s.Type().Field(i).Type.Kind() == reflect.Struct {
			if err := loadStruct(s.Field(i)); err != nil {
				return err
			}
		} else if s.Type().Field(i).Type.Kind() == reflect.Ptr {
			if s.Field(i).IsZero() == false && s.Field(i).Elem().Type().Kind() == reflect.Struct {
				if err := loadStruct(s.Field(i).Elem()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
