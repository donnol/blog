package rangestring

import "fmt"

func RangeString(s string) {
	var _ byte // type byte = uint8
	var _ rune // type rune = int32

	fmt.Printf("len(s) is %d, len([]rune(s)) is %d\n", len(s), len([]rune(s)))

	for i := 0; i < len(s); i++ {
		k := s[i]
		_ = k

		fmt.Printf("byte is %v\n", k)
	}

	// 当用range遍历字符串时，是在遍历它的[]rune
	for i, v := range s {
		v := v
		_ = v

		k := s[i]
		_ = k

		fmt.Printf("rune is %v, rune bin is %b; byte is %v, byte bin is %b\n", v, v, k, k)
	}
}
