package validation

import (
	"fmt"
	"reflect"
	"strconv"
)

func ValidateStructFields(v interface{}, path string) ([]string, error) {
	var emptyFields []string
	var requiredErrors []string

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("CheckStructFields expects a struct, got %s", val.Kind())
	}

	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := val.Field(i)
		yamlTag := field.Tag.Get("yaml")
		requiredTag := field.Tag.Get("required")
		fieldPath := path + yamlTag

		if field.Type.Kind() == reflect.Struct && (requiredTag == "" || requiredTag == "true") {
			nestedEmpty, err := ValidateStructFields(fieldValue.Interface(), fieldPath+".")
			if err != nil {
				return nil, err
			}
			emptyFields = append(emptyFields, nestedEmpty...)
		} else if IsStructFieldEmpty(fieldValue) && (requiredTag == "" || requiredTag == "true") {
			emptyFields = append(emptyFields, fieldPath)
			requiredErrors = append(requiredErrors, fieldPath)
		}
	}

	if len(requiredErrors) > 0 {
		return nil, fmt.Errorf("required fields are empty: %v", requiredErrors)
	}

	return emptyFields, nil
}

func IsStructFieldEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, err := strconv.ParseInt(v.String(), 10, 64)
		return err != nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, err := strconv.ParseUint(v.String(), 10, 64)
		return err != nil
	case reflect.Float32, reflect.Float64:
		_, err := strconv.ParseFloat(v.String(), 64)
		return err != nil
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Struct:
		return v.NumField() == 0
	}
	return false
}
