package option

import (
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
