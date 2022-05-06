package enum

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type EnumField[T comparable] struct {
	key    string // 结构体的字段名
	value  T
	name   string
	zhName string
}

func (e EnumField[T]) Value() T {
	return e.value
}

func (e EnumField[T]) Name() string {
	return e.name
}

func (e EnumField[T]) ZhName() string {
	return e.zhName
}

func Convert[T comparable, R any](ef EnumField[T], f func(EnumField[T]) R) R {
	return f(ef)
}

type Enum[T comparable] struct {
	fields []EnumField[T]

	values   []T
	fieldMap map[T]EnumField[T]
}

func (e Enum[T]) Values() []T {
	return e.values
}

func (e Enum[T]) Fields() []EnumField[T] {
	return e.fields
}

func (e Enum[T]) Map() map[T]EnumField[T] {
	return e.fieldMap
}

func (e Enum[T]) NameByValue(value T) string {
	return e.fieldMap[value].name
}

func (e Enum[T]) ZhNameByValue(value T) string {
	return e.fieldMap[value].zhName
}

// Init 需要传入结构体指针实例，否则报错；将根据它的enum tag的值来初始化该实例
// 因为枚举值一般只需一套固定值，所以全局初始化一个实例即可
// 获取tag，为字段赋值，没有则报错
func Init[E comparable](e any) error {
	v := reflect.ValueOf(e)
	t := v.Type()
	if t.Kind() != reflect.Ptr {
		return (fmt.Errorf("input value is not a pointer"))
	}
	ve := v.Elem()
	te := t.Elem()
	if te.Kind() != reflect.Struct {
		return (fmt.Errorf("input value is not a pointer of struct"))
	}
	if te.NumField() == 0 {
		return (fmt.Errorf("input value don't have field"))
	}
	var fieldValueType reflect.Type
	for i := 0; i < te.NumField(); i++ {
		field := te.Field(i)

		if field.Anonymous {
			continue
		}

		if field.Type.Kind() != reflect.Struct {
			continue
		}

		for j := 0; j < field.Type.NumField(); j++ {
			ff := field.Type.Field(j)
			if ff.Name == "value" {
				fieldValueType = ff.Type
				break
			}
		}

		if fieldValueType != nil {
			break
		}
	}

	var enumObj = Enum[E]{
		fields:   make([]EnumField[E], 0, te.NumField()),
		values:   make([]E, 0, te.NumField()),
		fieldMap: make(map[E]EnumField[E], te.NumField()),
	}
	var enumValue reflect.Value
	for i := 0; i < te.NumField(); i++ {
		field := te.Field(i)

		if field.Anonymous {
			if field.Name == "Enum" {
				enumValue = ve.Field(i)
			}
			continue
		}

		ev, ok := field.Tag.Lookup("enum")
		if !ok {
			return (fmt.Errorf("field %s don't have an enum tag", field.Name))
		}

		vf := ve.Field(i)

		ef := EnumField[E]{
			key: field.Name,
		}
		evs := strings.Split(ev, ",")
		if len(evs) == 0 {
			return (fmt.Errorf("field tag don't have enough value set"))
		}
		var firstValue reflect.Value
		var fv E
		var isE bool
		switch fieldValueType.Kind() {
		case reflect.Int:
			v, err := strconv.Atoi(evs[0])
			if err != nil {
				return (err)
			}
			firstValue = reflect.ValueOf(v)
		case reflect.String:
			first := reflect.ValueOf(evs[0])
			firstValue = first
		default:
			return fmt.Errorf("not support %v yet", fieldValueType.Kind())
		}
		fv, isE = firstValue.Interface().(E)
		if !isE {
			eType := reflect.TypeOf(fv)
			fv = firstValue.Convert(eType).Interface().(E)
		}
		enumObj.values = append(enumObj.values, fv)
		switch len(evs) {
		case 1:
			ef.value = fv
		case 2:
			ef.value = fv
			ef.name = firstValue.Interface().(string)
			ef.zhName = evs[1]
		case 3:
			ef.value = fv
			ef.name = evs[1]
			ef.zhName = evs[2]
		}
		enumObj.fields = append(enumObj.fields, ef)

		// 枚举值不能重复
		if _, ok := enumObj.fieldMap[ef.value]; ok {
			return fmt.Errorf("enum value %v already exist, please use another value", ef.value)
		}
		enumObj.fieldMap[ef.value] = ef

		vf.Set(reflect.ValueOf(ef))
	}

	enumValue.Set(reflect.ValueOf(enumObj))

	return nil
}
