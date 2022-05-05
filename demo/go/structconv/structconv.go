package structconv

import (
	"fmt"
	"reflect"
)

// ConvByFieldName fill to with from by name
// like to.Name = from.Name
// to must be a struct pointer
func ConvByFieldName[F, T any](from F, to T) {
	toByFieldNameReflect(from, to)
}

var (
	emptyValue = reflect.Value{}
)

func toByFieldNameReflect[F, T any](from F, to T) {
	fromValue := reflect.ValueOf(from)
	if fromValue.Type().Kind() == reflect.Pointer {
		fromValue = fromValue.Elem()
	}

	toValue := reflect.ValueOf(to)
	toType := toValue.Type()
	if toType.Kind() != reflect.Pointer {
		panic(fmt.Errorf("to is not a pointer"))
	}
	toElemValue := toValue.Elem()
	toElemType := toType.Elem()
	if toElemType.Kind() != reflect.Struct {
		panic(fmt.Errorf("to is not a struct"))
	}

	for i := 0; i < toElemType.NumField(); i++ {
		field := toElemType.Field(i)
		fieldValue := toElemValue.Field(i)

		fromFieldValue := fromValue.FieldByName(field.Name)
		if fromFieldValue == emptyValue {
			continue
		}
		if fromFieldValue.Type().Kind() != fieldValue.Type().Kind() {
			continue
		}

		fieldValue.Set(fromFieldValue)
	}
}
