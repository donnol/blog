package inner

type Type struct {
	Name  string
	slice []int
}

type I interface {
	Name() string
}

type EI interface {
	I

	e()
}

type EI1 struct {
	name string
}

func (ei1 EI1) Name() string {
	return ei1.name
}

func (ei1 EI1) e() {

}
