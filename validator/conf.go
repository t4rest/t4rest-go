package validator

import (
	"reflect"
)

// Validator is interface for all validatable structures
type Validator interface {
	Validate() error
}

// Validate validates additional structures (which implements BaseConfig)
func Validate(cfg interface{}) error {
	baseConfigType := reflect.TypeOf((*Validator)(nil)).Elem()

	v := reflect.ValueOf(cfg).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Type.Implements(baseConfigType) {
			err := v.Field(i).Interface().(Validator).Validate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
