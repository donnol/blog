package store

type Order struct {
	Id        uint
	Name      string
	UserId    uint
	ProductId uint
	Price     int
	Amount    int
}

func (order *Order) Add() error {
	// insert into order (...) values (?...)
	return nil
}
