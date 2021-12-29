package service

import (
	"github.com/donnol/blog/demo/go/ddd/domain"
	"github.com/donnol/blog/demo/go/ddd/store"
)

type Service struct {
}

func (srv *Service) Buy(userID uint, productID uint, amount int) error {
	// 读库并初始化domain
	user := &store.User{}
	user.ById(userID)
	userDomain := domain.NewUser(user.Name, user.Phone, domain.Money(user.Balance))
	product := &store.Product{}
	product.ById(productID)
	productDomain := domain.NewProduct(product.Name, domain.Money(product.Price), product.Stock, nil)
	shop := &store.Shop{}
	shop.ById(product.ShopId)
	shopDomain := domain.NewShop(shop.Name, shop.Addr)
	productDomain.SetShopIfNil(shopDomain)

	// 执行domain
	orderDomain := domain.NewOrder(userDomain, productDomain, amount)

	// 状态存储
	{ // tx
		// 添加订单记录
		order := &store.Order{
			Name:      orderDomain.Name(),
			UserId:    userID,
			ProductId: productID,
			Price:     product.Price,
			Amount:    amount,
		}
		order.Add()
		// 减少用户余额
		(&store.User{
			Balance: int(orderDomain.User().Balance()),
		}).UpdateBalance(userID)
		// 减少产品库存
		(&store.Product{
			Stock: orderDomain.Product().Stock(),
		}).UpdateStock(productID)
	}

	return nil
}
