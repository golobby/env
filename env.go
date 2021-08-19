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
// Field example: IsAdmin bool `env:"IS_ADMIN"`
func Load(structure interface{}) error {
	receiverType := reflect.TypeOf(structure)
	if receiverType != nil && receiverType.Kind() == reflect.Ptr {
		elem := receiverType.Elem()
		if elem.Kind() == reflect.Struct {
			s := reflect.ValueOf(structure).Elem()
			return loadStruct(structure, s)
		}
	}

	return errors.New("env: invalid structure")
}

func loadStruct(structure interface{}, s reflect.Value) error {
	for i := 0; i < s.NumField(); i++ {
		if t, exist := s.Type().Field(i).Tag.Lookup("env"); exist {
			v, err := cast.FromString(os.Getenv(t), s.Type().Field(i).Type.Name())
			if err != nil {
				return err
			}

			ptr := reflect.NewAt(s.Field(i).Type(), unsafe.Pointer(s.Field(i).UnsafeAddr())).Elem()
			ptr.Set(reflect.ValueOf(v))
		} else if s.Type().Field(i).Type.Kind() == reflect.Struct {
			if err := loadStruct(structure, s.Field(i)); err != nil {
				return err
			}
		} else if s.Type().Field(i).Type.Kind() == reflect.Ptr {
			if s.Field(i).IsZero() == false && s.Field(i).Elem().Type().Kind() == reflect.Struct {
				if err := loadStruct(structure, s.Field(i).Elem()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
