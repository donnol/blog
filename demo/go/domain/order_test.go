package domain_test

import (
	"testing"

	"github.com/donnol/blog/demo/go/domain"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		user    *domain.User
		product *domain.Product
		c       int
	}
	tests := []struct {
		name string
		args args
		want *domain.Order
	}{
		{name: "", args: args{
			user:    domain.NewUser("jd", "123", 10000),
			product: domain.NewProduct("树莓派", 1000, 10, domain.NewShop("a shop", "zhongshan")),
			c:       1,
		}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.NewOrder(tt.args.user, tt.args.product, tt.args.c); got.User().Balance() != 9000 || got.Product().Stock() != 9 {
				t.Logf("user: %+v, product: %+v\n", got.User(), got.Product())
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
