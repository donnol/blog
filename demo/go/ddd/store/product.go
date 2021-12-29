package store

type Product struct {
	Id     uint
	Name   string
	Price  int
	Stock  int
	ShopId uint
}

func (product *Product) ById(id uint) error {
	// select * from product where id = ?
	// scan data to product
	return nil
}

func (product *Product) UpdateStock(id uint) error {
	// update product set stock = ? where id = ?
	return nil
}
