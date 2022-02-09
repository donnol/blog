package map_plus_plus

import "fmt"

func mapPlusPlus() {
	m := make(map[string]int)
	m["count"]++
	fmt.Printf("m: %+v\n", m)

	m["count"]++
	fmt.Printf("m: %+v\n", m)
}

type M struct {
	count int
}

func mapStructPlusPlus() {
	m := make(map[string]M)

	// cannot assign to struct field m["count"].count in map
	// m["count"].count++

	{
		v, ok := m["count"]
		if ok {
			v.count++
			m["count"] = v
		} else {
			m["count"] = M{count: 1}
		}
	}
	fmt.Printf("m: %+v\n", m)

	{
		v, ok := m["count"]
		if ok {
			v.count++
			m["count"] = v
		} else {
			m["count"] = M{count: 1}
		}
	}
	fmt.Printf("m: %+v\n", m)
}
