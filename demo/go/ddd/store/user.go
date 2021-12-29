package store

type User struct {
	Id      uint
	Name    string
	Phone   string
	Addr    string
	Balance int
}

func (u *User) ById(id uint) error {
	// select * from user where id = ?
	// scan data to u
	return nil
}

func (u *User) UpdateBalance(id uint) error {
	// update user set balance = ? where id = ?
	return nil
}
