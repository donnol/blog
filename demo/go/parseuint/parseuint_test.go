package parseuint

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestParseUint(t *testing.T) {
	var a uint = 1769282207731417088
	m := make(map[string]interface{})
	m["id"] = a
	fmt.Println("Hello, 世界", fmt.Sprintf("%v", m["id"]))
	r, err := strconv.ParseUint(fmt.Sprintf("%v", m["id"]), 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("r:", r)
}

func TestParseUintAfterJsonDecode(t *testing.T) {
	var a uint = 1769282207731417088

	rm := make(map[string]interface{})
	rm["id"] = a
	data, err := json.Marshal(rm)
	if err != nil {
		t.Fatal(err)
	}
	m := make(map[string]interface{})
	// json decode后，长整型数被截断，并使用科学计数法表示
	err = json.Unmarshal(data, &m)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(m["id"], "become:", fmt.Sprintf("%v", m["id"]))
	r, err := strconv.ParseUint(fmt.Sprintf("%v", m["id"]), 10, 64)
	if err != nil {
		t.Fatal(err) // strconv.ParseUint: parsing "1.769282207731417e+18": invalid syntax
		// 因为json没有区分整形和浮点，一律将数字认为是浮点，而它支持的范围又小了一点，导致长整形数被截断，数值变小，变为科学计数法表示，不符合语法格式
	}
	fmt.Println("r:", r)

	// 使用结构体，并声明字段类型为uint，则正常接收
	var u struct {
		Id uint `json:"id"`
	}
	err = json.Unmarshal(data, &u)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("u:", u)
}
