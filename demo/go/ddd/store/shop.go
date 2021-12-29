package store

type Shop struct {
	Id   uint
	Name string
	Addr string
}

func (shop *Shop) ById(id uint) error {
	// select * from shop where id = ?
	// scan data to shop
	return nil
}
