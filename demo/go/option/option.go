package option

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/samber/mo"
)

var (
	ErrFieldNotExist = errors.New("field not exist")
)

func OptionFromRawMessage[T any](m json.RawMessage) (mo.Option[T], error) {
	var dst mo.Option[T]
	if len(m) == 0 {
		fmt.Printf("%s is not exist '%s'\n", "field", m)

		return dst, ErrFieldNotExist
	}

	if string(m) == "null" {
		fmt.Printf("%s is there but null\n", "field")
	} else {
		fmt.Printf("%s is there and not null: %s\n", "field", m)
	}
	if err := json.Unmarshal([]byte(m), &dst); err != nil {
		return dst, err
	}

	return dst, nil
}

type Option[T any] struct {
	present int8 // 0 不存在；1 为null；2 常值；
	value   T
}

// Exist 表明在json字符串里该字段是否存在
func (o Option[T]) Exist() bool {
	return o.present != 0
}

// Exist 表明在json字符串里该字段是否值为null
func (o Option[T]) IsNull() bool {
	return o.present == 1
}

// Exist 表明在json字符串里该字段是否为常值
func (o Option[T]) HasValue() bool {
	return o.present == 2
}

func (o Option[T]) Get() (T, bool) {
	if o.HasValue() {
		return o.value, true
	}
	return o.value, false
}

func (o Option[T]) MustGet() T {
	v, ok := o.Get()
	if !ok {
		panic(fmt.Errorf("value not exist"))
	}
	return v
}

// MarshalJSON encodes Option into json.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.present == 2 {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON decodes Option from json.
func (o *Option[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		o.present = 1
		return nil
	}

	err := json.Unmarshal(b, &o.value)
	if err != nil {
		return err
	}

	o.present = 2
	return nil
}
