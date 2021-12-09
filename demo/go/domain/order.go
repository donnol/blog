package domain

import (
	"crypto/rand"
	"math/big"

	"github.com/pkg/errors"
)

// 关键词：用户、店铺、商品、订单
//
// 场景描述：店铺展示商品，其价格为P、库存为N，用户（余额为Y）看到商品觉得合适，于是下单购买B个；购买前，用户余额Y必须不小于P，商品库存N不小于B；购买后，用户余额减少P，库存减少B；
//
// 先不考虑并发情况，建立此时的领域模型

type User struct {
	name  string // 名称
	phone string // 电话

	balance Money // 余额
}

type Shop struct {
	name string // 名称
	addr string // 地址
}

type Product struct {
	name  string // 名称
	price Money  // 价格
	stock int    // 库存

	ownShop *Shop // 所属商铺
}

type Order struct {
	name string // 名称

	user    *User    // 用户
	product *Product // 商品
}

type Money int

func NewUser(name, phone string, bal Money) *User {
	return &User{
		name:    name,
		phone:   phone,
		balance: bal,
	}
}
func (u *User) Balance() Money {
	return u.balance
}
func (u *User) DeductBalance(amount Money) {
	if u.balance < amount {
		panic("not enough money")
	}
	u.balance -= amount
}

func NewShop(name, addr string) *Shop {
	return &Shop{
		name: name,
		addr: addr,
	}
}
func NewProduct(name string, price Money, stock int, shop *Shop) *Product {
	return &Product{
		name:    name,
		price:   price,
		stock:   stock,
		ownShop: shop,
	}
}
func (p *Product) Stock() int {
	return p.stock
}
func (p *Product) DeductStock(c int) {
	if p.stock < c {
		panic("not enough stock")
	}
	p.stock -= c
}

// NewOrder 用户对商品下单c个
func NewOrder(user *User, product *Product, c int) *Order {
	name, err := GenerateRandomString(12)
	if err != nil {
		panic(err)
	}

	user.DeductBalance(product.price * Money(c))
	product.DeductStock(c)

	return &Order{
		name:    name,
		user:    user,
		product: product,
	}
}

func (o *Order) User() *User {
	return o.user
}

func (o *Order) Product() *Product {
	return o.product
}

// GenerateRandomString 随机字符串包含有数字和大小写字母
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	return generate(n, letters)
}

func generate(n int, letters string) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", errors.WithStack(err)
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
