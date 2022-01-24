package generic

type A interface {
	int|string
}

type AI interface {
	I()
}

// type B = interface {
// 	int | string 
// }

type BI = interface {
	I()
}

// type C = int | string
