package enum

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// 定义结构的同时，通过struct tag设置值
// 字段，字段类型，[字段整形值、]字段英文值、字段中文值(第一个值将作为map的key)
// 如果不内嵌Enum[T]，要怎么拿到枚举的集合呢？
type (
	Color = EnumField[int]
)

var (
	ColorEnumObj struct {
		Enum[int]
		Blue  Color `enum:"1,blue,蓝色"`
		Green Color `enum:"2,green,绿色"`
		Red   Color `enum:"3,red,红色"`
	}
)

// 考虑到结构体只需出现一次，直接使用匿名结构体，并且在定义结构体的同时声明变量
var (
	WeekdayEnumObj struct {
		Enum[string]
		Monday    EnumField[string] `enum:"monday,星期一"`
		Tuesday   EnumField[string] `enum:"tuesday,星期二"`
		Wednesday EnumField[string] `enum:"wednesday,星期三"`
		Thursday  EnumField[string] `enum:"thursday,星期四"`
		Friday    EnumField[string] `enum:"friday,星期五"`
		Saturday  EnumField[string] `enum:"saturday,星期六"`
		Sunday    EnumField[string] `enum:"sunday,星期日"`
	}
)

// 初始化
func init() {
	panicIf(Init[int](&ColorEnumObj))

	panicIf(Init[string](&WeekdayEnumObj))
}

func TestEnum(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		handlers []handler
	}{
		{"color", colorHandlers},
		{"weekday", weekdayHandlers},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			for _, h := range testCase.handlers {
				r := h.handler()
				if !reflect.DeepEqual(r, h.want) {
					t.Fatalf("bad case: %v != %v\n", r, h.want)
				}
			}
		})
	}
}

type (
	handler struct {
		handler func() any
		want    any
	}
)

var (
	colorHandlers = []handler{
		{func() any { return ColorEnumObj.Blue.Value() }, 1},
		{func() any { return ColorEnumObj.Green.Value() }, 2},
		{func() any { return ColorEnumObj.Red.Value() }, 3},
		{func() any { return ColorEnumObj.Blue.Name() }, "blue"},
		{func() any { return ColorEnumObj.Green.Name() }, "green"},
		{func() any { return ColorEnumObj.Red.Name() }, "red"},
		{func() any { return ColorEnumObj.Blue.ZhName() }, "蓝色"},
		{func() any { return ColorEnumObj.Green.ZhName() }, "绿色"},
		{func() any { return ColorEnumObj.Red.ZhName() }, "红色"},

		{func() any { return ColorEnumObj.Values() }, []int{1, 2, 3}},
		{func() any { return ColorEnumObj.Fields() }, []EnumField[int]{
			{"Blue", 1, "blue", "蓝色"},
			{"Green", 2, "green", "绿色"},
			{"Red", 3, "red", "红色"},
		}},
		{func() any { return ColorEnumObj.Map() }, map[int]EnumField[int]{
			1: {"Blue", 1, "blue", "蓝色"},
			2: {"Green", 2, "green", "绿色"},
			3: {"Red", 3, "red", "红色"},
		}},
		{func() any { return ColorEnumObj.NameByValue(ColorEnumObj.Blue.Value()) }, "blue"},
		{func() any { return ColorEnumObj.NameByValue(ColorEnumObj.Green.Value()) }, "green"},
		{func() any { return ColorEnumObj.NameByValue(ColorEnumObj.Red.Value()) }, "red"},
		{func() any { return ColorEnumObj.ZhNameByValue(ColorEnumObj.Blue.Value()) }, "蓝色"},
		{func() any { return ColorEnumObj.ZhNameByValue(ColorEnumObj.Green.Value()) }, "绿色"},
		{func() any { return ColorEnumObj.ZhNameByValue(ColorEnumObj.Red.Value()) }, "红色"},
	}

	weekdayHandlers = []handler{
		{func() any { return WeekdayEnumObj.Monday.Value() }, "monday"},
		{func() any { return WeekdayEnumObj.Tuesday.Value() }, "tuesday"},
		{func() any { return WeekdayEnumObj.Wednesday.Value() }, "wednesday"},
		{func() any { return WeekdayEnumObj.Thursday.Value() }, "thursday"},
		{func() any { return WeekdayEnumObj.Friday.Value() }, "friday"},
		{func() any { return WeekdayEnumObj.Saturday.Value() }, "saturday"},
		{func() any { return WeekdayEnumObj.Sunday.Value() }, "sunday"},
		{func() any { return WeekdayEnumObj.Monday.Name() }, "monday"},
		{func() any { return WeekdayEnumObj.Tuesday.Name() }, "tuesday"},
		{func() any { return WeekdayEnumObj.Wednesday.Name() }, "wednesday"},
		{func() any { return WeekdayEnumObj.Thursday.Name() }, "thursday"},
		{func() any { return WeekdayEnumObj.Friday.Name() }, "friday"},
		{func() any { return WeekdayEnumObj.Saturday.Name() }, "saturday"},
		{func() any { return WeekdayEnumObj.Sunday.Name() }, "sunday"},
		{func() any { return WeekdayEnumObj.Monday.ZhName() }, "星期一"},
		{func() any { return WeekdayEnumObj.Tuesday.ZhName() }, "星期二"},
		{func() any { return WeekdayEnumObj.Wednesday.ZhName() }, "星期三"},
		{func() any { return WeekdayEnumObj.Thursday.ZhName() }, "星期四"},
		{func() any { return WeekdayEnumObj.Friday.ZhName() }, "星期五"},
		{func() any { return WeekdayEnumObj.Saturday.ZhName() }, "星期六"},
		{func() any { return WeekdayEnumObj.Sunday.ZhName() }, "星期日"},

		{func() any { return WeekdayEnumObj.Values() }, []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}},
		{func() any { return WeekdayEnumObj.Fields() }, []EnumField[string]{
			{"Monday", "monday", "monday", "星期一"},
			{"Tuesday", "tuesday", "tuesday", "星期二"},
			{"Wednesday", "wednesday", "wednesday", "星期三"},
			{"Thursday", "thursday", "thursday", "星期四"},
			{"Friday", "friday", "friday", "星期五"},
			{"Saturday", "saturday", "saturday", "星期六"},
			{"Sunday", "sunday", "sunday", "星期日"},
		}},
		{func() any { return WeekdayEnumObj.Map() }, map[string]EnumField[string]{
			"monday":    {"Monday", "monday", "monday", "星期一"},
			"tuesday":   {"Tuesday", "tuesday", "tuesday", "星期二"},
			"wednesday": {"Wednesday", "wednesday", "wednesday", "星期三"},
			"thursday":  {"Thursday", "thursday", "thursday", "星期四"},
			"friday":    {"Friday", "friday", "friday", "星期五"},
			"saturday":  {"Saturday", "saturday", "saturday", "星期六"},
			"sunday":    {"Sunday", "sunday", "sunday", "星期日"},
		}},
		{func() any {
			return WeekdayEnumObj.NameByValue(WeekdayEnumObj.Monday.Value())
		}, "monday"},
		{func() any {
			return WeekdayEnumObj.NameByValue(WeekdayEnumObj.Tuesday.Value())
		}, "tuesday"},
		{func() any {
			return WeekdayEnumObj.NameByValue(WeekdayEnumObj.Wednesday.Value())
		}, "wednesday"},
		{func() any {
			return WeekdayEnumObj.NameByValue(WeekdayEnumObj.Thursday.Value())
		}, "thursday"},
		{func() any {
			return WeekdayEnumObj.NameByValue(WeekdayEnumObj.Friday.Value())
		}, "friday"},
		{func() any {
			return WeekdayEnumObj.NameByValue(WeekdayEnumObj.Saturday.Value())
		}, "saturday"},
		{func() any {
			return WeekdayEnumObj.NameByValue(WeekdayEnumObj.Sunday.Value())
		}, "sunday"},
		{func() any {
			return WeekdayEnumObj.ZhNameByValue(WeekdayEnumObj.Monday.Value())
		}, "星期一"},
		{func() any {
			return WeekdayEnumObj.ZhNameByValue(WeekdayEnumObj.Tuesday.Value())
		}, "星期二"},
		{func() any {
			return WeekdayEnumObj.ZhNameByValue(WeekdayEnumObj.Wednesday.Value())
		}, "星期三"},
		{func() any {
			return WeekdayEnumObj.ZhNameByValue(WeekdayEnumObj.Thursday.Value())
		}, "星期四"},
		{func() any {
			return WeekdayEnumObj.ZhNameByValue(WeekdayEnumObj.Friday.Value())
		}, "星期五"},
		{func() any {
			return WeekdayEnumObj.ZhNameByValue(WeekdayEnumObj.Saturday.Value())
		}, "星期六"},
		{func() any {
			return WeekdayEnumObj.ZhNameByValue(WeekdayEnumObj.Sunday.Value())
		}, "星期日"},
	}
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	ErrorEnumObj struct {
		Enum[int]

		Ok             EnumField[int] `enum:"0,ok,正常"`
		BadParam       EnumField[int] `enum:"1,BadParam,参数错误"`
		BusinessFailed EnumField[int] `enum:"2,BusinessFailed,业务执行失败"`
		EncodeFailed   EnumField[int] `enum:"3,EncodeFailed,编码失败"`
		DecodeFailed   EnumField[int] `enum:"4,DecodeFailed,解码失败"`
	}
)

type (
	Result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}
)

func TestEnumError(t *testing.T) {
	err := Init[int](&ErrorEnumObj)
	if err != nil {
		t.Fatal(err)
	}

	badParam := Convert(ErrorEnumObj.BadParam, func(ef EnumField[int]) error {
		return fmt.Errorf("code: %v, msg: %s-%s", ef.Value(), ef.Name(), ef.ZhName())
	})
	if badParam.Error() != "code: 1, msg: BadParam-参数错误" {
		t.Fatalf("badParam: %+v\n", badParam)
	}

	res := Convert(ErrorEnumObj.Ok, func(ef EnumField[int]) Result {
		return Result{
			Code: ef.value,
			Msg:  ef.ZhName(),
		}
	})
	data, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"code":0,"msg":"正常","data":null}` {
		t.Fatalf("bad result: %s != %s\n", data, `{"code":0,"msg":"正常","data":null}`)
	}
}

var (
	repeatedEnumValue struct {
		Enum[int]
		First  EnumField[int] `enum:"1"`
		Repeat EnumField[int] `enum:"1"`
	}
)

func TestRepeatedEnum(t *testing.T) {
	err := Init[int](&repeatedEnumValue)
	if err.Error() != "enum value 1 already exist, please use another value" {
		t.Fatal(err)
	}
}
