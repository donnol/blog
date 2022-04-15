package lo

import (
	"testing"

	"github.com/samber/lo"
)

func TestLo(t *testing.T) {
	r := lo.Map([]int{1, 2, 3}, func(t int, i int) int {
		return t + i // i是遍历时的下标，t是元素
	})
	t.Logf("r: %+v\n", r) // [1, 3, 5]

	re := lo.Reduce([]int{1, 2, 3}, func(r int, t int, i int) int {
		return r + t // r是初始值（也就是Reduce的最后一个参数），t是元素值，i是下标
	}, 0)
	t.Logf("re: %+v\n", re)

	// 级联
	re = lo.Reduce(lo.Map([]int{1, 2, 3}, func(t int, i int) int {
		return t + i // i是遍历时的下标，t是元素
	}), func(r int, t int, i int) int {
		return r + t
	}, 0)
	t.Logf("re: %+v\n", re)
}

func TestObjectArray(t *testing.T) {
	type M struct {
		Id     uint
		Name   string
		Person struct {
			Age   int
			Phone string
		}
		Addrs []struct {
			Company string
			Home    string
		}
	}
	data := []M{
		{
			Id:   1,
			Name: "jd",
			Person: struct {
				Age   int
				Phone string
			}{
				Age:   30,
				Phone: "12345678901",
			},
			Addrs: []struct {
				Company string
				Home    string
			}{
				{Company: "GZ", Home: "QY"},
			},
		},
		{
			Id:   2,
			Name: "jq",
			Person: struct {
				Age   int
				Phone string
			}{
				Age:   20,
				Phone: "22345678901",
			},
			Addrs: []struct {
				Company string
				Home    string
			}{
				{Company: "FS", Home: "QX"},
			},
		},
	}

	// 收集id
	ids := lo.Map(data, func(t M, i int) uint {
		return t.Id
	})
	t.Logf("ids: %v\n", ids) // [1, 2]

	// 做映射
	m := lo.KeyBy(data, func(v M) uint {
		return v.Id // 使用Id作key
	})
	t.Logf("m: %+v\n", m) // map[1:{Id:1 Name:jd Person:{Age:30 Phone:12345678901} Addrs:[{Company:GZ Home:QY}]} 2:{Id:2 Name:jq Person:{Age:20 Phone:22345678901} Addrs:[{Company:FS Home:QX}]}]

	m2 := lo.KeyBy(data, func(v M) string {
		return v.Name // 使用Name作key
	})
	t.Logf("m2: %+v\n", m2) // map[jd:{Id:1 Name:jd Person:{Age:30 Phone:12345678901} Addrs:[{Company:GZ Home:QY}]} jq:{Id:2 Name:jq Person:{Age:20 Phone:22345678901} Addrs:[{Company:FS Home:QX}]}]
}
