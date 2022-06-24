package new_make

type A struct {
	Name string
	Age  uint
}

func (a A) String() string {
	return a.Name
}
