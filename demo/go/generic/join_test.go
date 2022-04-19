package generic

import (
	"reflect"
	"testing"
)

func Test_leftJoinById(t *testing.T) {
	type args[L, R IdConstraint[uint], LR any] struct {
		left  []L
		right []R
		f     func(l L, r R) LR
	}
	type testCase[L, R IdConstraint[uint], LR any] struct {
		name string
		args args[L, R, LR]
		want []LR
	}
	type User struct {
		Id uint
		// Name string

		Data any // 那不是还要从Data里提取具体数据？
	}
	type Article struct {
		Id uint
		// UserId uint
		// Name   string

		Data any
	}
	type ArticleUser struct {
		ArticleName string
		UserName    string
	}
	tests := []testCase[User, Article, any]{
		{
			name: "1",
			args: args[User, Article, any]{
				left: []User{
					{
						Id:   1,
						Data: "jd",
					},
				},
				right: []Article{
					{
						Id:   1,
						Data: "go generic",
					},
				},
				f: func(l User, r Article) any {
					return ArticleUser{
						ArticleName: r.Data.(string),
						UserName:    l.Data.(string),
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftJoinById(tt.args.left, tt.args.right, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftJoinById() = %v, want %v", got, tt.want)
			}
		})
	}
}
