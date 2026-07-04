package syml

import (
	"errors"
	"fmt"
	"reflect"
)

var file_extension = "syml"

/*
 * This Program is to extend the YAML instead of creating something entirely new
 * This programs goal is to create a perser for markup language where it will perse existing the yaml file but add more features
 * fix annoying issues in existing yaml file
 */

func Unmarshal(data []byte, v any) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("v must be a non-nil pointer")
	}

	elem := rv.Elem()

	switch elem.Kind() {
	case reflect.Struct:
		fmt.Println("Processing struct")
		return UnmarshalToStruct(data, rv)
	case reflect.Map:
		// Decode into a map
	case reflect.Slice:
		// Decode into a slice
	default:
		return fmt.Errorf("unsupported type: %s", elem.Kind())
	}

	return nil
}

// It will only work with struct
func UnmarshalToStruct(data []byte, v reflect.Value) error {
	token := Lexer(data, DefaultSettings())

	for _, t := range token {
		fmt.Println(t.String())
	}
	return nil
}
