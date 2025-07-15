package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ReportIfZeroValue reporta error si el campo es el valor cero para su tipo
func ReportIfZeroValue(sl validator.StructLevel, field interface{}, fieldName, structFieldName, tag string) {
	if isEmpty(field) {
		sl.ReportError(field, fieldName, structFieldName, tag, "")
	}
}

// isEmpty verifica si un valor es "vacío" según su tipo
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	// Manejar punteros
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		return strings.TrimSpace(v.String()) == ""
	case reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool() // Opcional: considerar false como "vacío"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Struct:
		// Para time.Time y otras estructuras
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	default:
		return false
	}
}
