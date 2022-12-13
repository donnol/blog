package interfacex

type I interface {
	Method()
}

type IunexportMethod interface {
	I
	unexportMethod()
}
