package number

import "math"

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Mul1(a, b, c float64) float64 {
	return roundFloat(a*b*c, 2)
}

func Mul2(b, c float64, as [3]float64) float64 {
	r := 0.0
	for _, a := range as {
		r += a * b * c
	}
	return roundFloat(r, 2)
}

func Mul3(b, c float64, as [3]float64) float64 {
	// (a*b*c*a1/a) +
	// ((a*b*c - a*b*c*a1/a)*a2/(a-a1)) +
	// ((a*b*c - a*b*c*a1/a) - (a*b*c - a*b*c*a1/a)*a2/(a-a1))*a3/(a-a1-a2)
	a := 0.0
	for _, s := range as {
		a += s
	}

	r := a*b*c*as[0]/a +
		(a*b*c-a*b*c*as[0]/a)*as[1]/(as[1]+as[2]) +
		((a*b*c-a*b*c*as[0]/a)-(a*b*c-a*b*c*as[0]/a)*as[1]/(as[1]+as[2]))*as[2]/as[2]
	return roundFloat(r, 2)
}
