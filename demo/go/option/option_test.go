package option

import (
	"encoding/json"
	"testing"

	"github.com/samber/mo"
)

func TestOption(t *testing.T) {
	// 推荐
	{
		oi := mo.Option[int]{}
		data := `null`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if v, ok := oi.Get(); ok {
			t.Errorf("get %v, %v", ok, v)
		}
	}

	{
		oi := mo.Option[int]{}
		data := `0`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if oi.MustGet() != 0 {
			t.Errorf("bad case, %v != %v", oi.MustGet(), 0)
		}
	}

	{
		oi := mo.Option[int]{}
		data := `1`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if oi.MustGet() != 1 {
			t.Errorf("bad case, %v != %v", oi.MustGet(), 1)
		}
	}

	// 不推荐
	{
		oi := mo.Option[*int]{}
		data := `null`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if v, ok := oi.Get(); ok {
			t.Errorf("get %v, %v", ok, v)
		}
	}

	{
		oi := mo.Option[*int]{}
		data := `0`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if v, ok := oi.Get(); !ok || *v != 0 {
			t.Errorf("get %v, %v", ok, v)
		}
	}

	{
		oi := mo.Option[*int]{}
		data := `1`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if v, ok := oi.Get(); !ok || *v != 1 {
			t.Errorf("get %v, %v", ok, v)
		}
	}

	// 作为字段，null、零值、常值
	type M struct {
		Name mo.Option[string] `json:"name"`
		Age  mo.Option[int]    `json:"age"`
	}
	{
		oi := M{}
		data := `{}` // 键不存在 -- 没传
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if v, ok := oi.Name.Get(); ok {
			t.Errorf("get %v, %v", ok, v)
		}
		if v, ok := oi.Age.Get(); ok {
			t.Errorf("get %v, %v", ok, v)
		}
	}
	{
		oi := M{}
		data := `{"name": null, "age": null}` // 键存在，值为null -- 传了但null
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if v, ok := oi.Name.Get(); ok {
			t.Errorf("get %v, %v", ok, v)
		}
		if v, ok := oi.Age.Get(); ok {
			t.Errorf("get %v, %v", ok, v)
		}
	}
	{
		oi := M{}
		data := `{"name": "", "age": 0}`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if v, ok := oi.Name.Get(); !ok || v != "" {
			t.Errorf("bad case, %v != %v", v, "")
		}
		if v, ok := oi.Age.Get(); !ok || v != 0 {
			t.Errorf("bad case, %v != %v", v, 0)
		}
	}
	{
		oi := M{}
		data := `{"name": "jd", "age": 16}`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if oi.Name.MustGet() != "jd" {
			t.Errorf("bad case, %v != %v", oi.Name.MustGet(), "jd")
		}
		if oi.Age.MustGet() != 16 {
			t.Errorf("bad case, %v != %v", oi.Age.MustGet(), 16)
		}
	}

	// 加一层指针，能区分'没传'和'值为null'吗？
	type MP struct {
		Name *mo.Option[string] `json:"name"`
		Age  *mo.Option[int]    `json:"age"`
	}
	{
		oi := MP{}
		data := `{}` // 键不存在 -- 没传
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		// 字段值均为nil
		if oi.Name != nil {
			t.Errorf("get %v", oi.Name)
		}
		if oi.Age != nil {
			t.Errorf("get %v", oi.Age)
		}
	}
	{
		oi := MP{}
		data := `{"name": null, "age": null}` // 键存在，值为null -- 传了但null
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		// 字段值均为nil
		if oi.Name != nil {
			t.Errorf("get %v", oi.Name)
		}
		if oi.Age != nil {
			t.Errorf("get %v", oi.Age)
		}
	}
	// 显然，还是不行，因为当字段类型为指针，而传入json值为null时，字段值也只会是nil，不会给mo.Option赋值

	// 那怎么区分'没传'和'值为null'呢？
	// use json.RawMessage. -- https://stackoverflow.com/questions/36601367/json-field-set-to-null-vs-field-not-there
	type MR struct {
		Name json.RawMessage `json:"name"`
	}
	{
		oi := MR{}
		data := `{}` // 键不存在 -- 没传
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		if _, err := OptionFromRawMessage[any](oi.Name); err != ErrFieldNotExist {
			t.Errorf("bad case, got %v", err)
		}
	}
	{
		oi := MR{}
		data := `{"name": null}` // 键存在，值为null -- 传了但null
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		name, err := OptionFromRawMessage[string](oi.Name)
		if err != nil {
			t.Error(err)
		}
		if v, ok := name.Get(); ok {
			t.Errorf("got %v, %v", ok, v)
		}
	}
	{
		oi := MR{}
		data := `{"name": "jd"}`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		name, err := OptionFromRawMessage[string](oi.Name)
		if err != nil {
			t.Error(err)
		}
		if v, ok := name.Get(); !ok || v != "jd" {
			t.Errorf("got %v, %v", ok, v)
		}
	}
	{
		oi := MR{}
		data := `{"name": 1}`
		if err := json.Unmarshal([]byte(data), &oi); err != nil {
			t.Error(err)
		}
		_, err := OptionFromRawMessage[string](oi.Name)
		if err != nil && err.Error() != "json: cannot unmarshal number into Go value of type string" {
			t.Error(err)
		}
	}
	// 不过也不好用，先要定义一个字段类型都是json.RawMessage的结构体来解析，然后解析得到之后根据长度判断完是否存在之后还要再次执行json解析，得不偿失
	// 希望官方添加这方面的支持吧
}
