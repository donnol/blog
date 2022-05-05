package structconv

import "testing"

// 虽然通过名字来赋值，省下了功夫，但现实中字段对应往往不是简单的通过名字就能实现，还有其它因素，所以实用性并不太高
type (
	UserReq struct {
		Phone string
	}

	UserResp struct {
		Name string
		Age  uint
	}

	UserTable struct {
		Id    uint
		Name  string
		Age   uint
		Phone string
	}

	Article struct {
		Name     string
		UserName string
	}
)

func TestConvByFieldName(t *testing.T) {
	from := UserReq{
		Phone: "12345678901",
	}
	to := &UserTable{}
	ConvByFieldName(from, to)

	if to.Phone != from.Phone {
		t.Fatalf("converse failed: %s != %s\n", to.Phone, from.Phone)
	}

	to.Id = 1
	to.Name = "jd"
	to.Age = 18

	to2 := &UserResp{}
	ConvByFieldName(to, to2)

	if to2.Name != to.Name {
		t.Fatalf("converse failed: %s != %s\n", to2.Name, to.Name)
	}
	if to2.Age != to.Age {
		t.Fatalf("converse failed: %d != %d\n", to2.Age, to.Age)
	}
}
