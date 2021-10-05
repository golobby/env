// Package env is a lightweight library for loading OS environment variables into structs.
package env

import (
	"errors"
	"fmt"
	"github.com/golobby/cast"
	"os"
	"reflect"
	"unsafe"
)

// Feed loads OS environment variables into the given struct.
// It gets the pointer of a struct that is going to holds the variables.
// The struct fields must have an `env` tag that determines the related OS environment variable.
// `env:"APP_NAME"` is a tag example.
func Feed(structure interface{}) error {
	inputType := reflect.TypeOf(structure)
	if inputType != nil {
		if inputType.Kind() == reflect.Ptr {
			if inputType.Elem().Kind() == reflect.Struct {
				return fillStruct(reflect.ValueOf(structure).Elem())
			}
		}
	}

	return errors.New("env: invalid structure")
}

// fillStruct sets a reflected struct fields with the equivalent OS environment variables.
func fillStruct(s reflect.Value) error {
	for i := 0; i < s.NumField(); i++ {
		if t, exist := s.Type().Field(i).Tag.Lookup("env"); exist {
			if osv := os.Getenv(t); osv != "" {
				v, err := cast.FromType(osv, s.Type().Field(i).Type)
				if err != nil {
					return fmt.Errorf("env: cannot set `%v` field; err: %v", s.Type().Field(i).Name, err)
				}

				ptr := reflect.NewAt(s.Field(i).Type(), unsafe.Pointer(s.Field(i).UnsafeAddr())).Elem()
				ptr.Set(reflect.ValueOf(v))
			}
		} else if s.Type().Field(i).Type.Kind() == reflect.Struct {
			if err := fillStruct(s.Field(i)); err != nil {
				return err
			}
		} else if s.Type().Field(i).Type.Kind() == reflect.Ptr {
			if s.Field(i).IsZero() == false && s.Field(i).Elem().Type().Kind() == reflect.Struct {
				if err := fillStruct(s.Field(i).Elem()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
