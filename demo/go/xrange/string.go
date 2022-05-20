package xrange

import "fmt"

func RangeString(s string) {
	for i, one := range s {
		// one的类型是rune(int32: 4字节整形，表示unicode字符)
		e := one
		fmt.Printf("%d, %v, %s\n", i, e, string(e))
	}
}

func ForString(s string) {
	for i := 0; i < len(s); i++ {
		// b的类型是byte(uint8: 1字节整形)
		b := s[i]
		fmt.Printf("%d, %v, %s\n", i, b, string(b))
	}
}
