package strconv

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var testString = "Albert Einstein: Logic will get you from A to B. Imagination will take you everywhere."
var testBytes = []byte(testString)

func rawBytesToStr(b []byte) string {
	return string(b)
}

func rawStrToBytes(s string) []byte {
	return []byte(s)
}

// go test -v

func TestBytesToString(t *testing.T) {
	data := make([]byte, 1024)
	for i := 0; i < 100; i++ {
		rand.Read(data)
		if rawBytesToStr(data) != BytesToString(data) {
			t.Fatal("don't match")
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func TestStringToBytes(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := RandStringBytesMaskImprSrcSB(64)
		if !bytes.Equal(rawStrToBytes(s), StringToBytes(s)) {
			t.Fatal("don't match")
		}
	}
}

func TestStringToBytes0(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := RandStringBytesMaskImprSrcSB(64)
		if !bytes.Equal(rawStrToBytes(s), StringToBytes0(s)) {
			t.Fatal("don't match")
		}
	}
}

// func TestStringToBytes1(t *testing.T) {
// 	for i := 0; i < 100; i++ {
// 		s := RandStringBytesMaskImprSrcSB(64)
// 		if !bytes.Equal(rawStrToBytes(s), StringToBytes1(s)) {
// 			t.Fatal("don't match")
// 		}
// 	}
// }

// func TestStringToBytes2(t *testing.T) {
// 	for i := 0; i < 100; i++ {
// 		s := RandStringBytesMaskImprSrcSB(64)
// 		if !bytes.Equal(rawStrToBytes(s), StringToBytes2(s)) {
// 			t.Fatal("don't match")
// 		}
// 	}
// }

// go test -v -run=none -bench=^BenchmarkBytesConv -benchmem=true

func BenchmarkBytesConvBytesToStrRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawBytesToStr(testBytes)
	}
}

func BenchmarkBytesConvBytesToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString(testBytes)
	}
}

func BenchmarkBytesConvStrToBytesRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawStrToBytes(testString)
	}
}

func BenchmarkBytesConvStrToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(testString)
	}
}

func BenchmarkBytesConvStrToBytes0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes0(testString)
	}
}

// func BenchmarkBytesConvStrToBytes1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		StringToBytes1(testString)
// 	}
// }

// func BenchmarkBytesConvStrToBytes2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		StringToBytes2(testString)
// 	}
// }

func TestStringToBytesModify(t *testing.T) {
	a := "hello"
	ab := StringToBytes(a)
	t.Logf("a: %s, ab: %s\n", a, ab)

	// byte by byte
	for i := 0; i < len(ab); i++ {
		t.Logf("ab[i]: %v\n", ab[i])
	}

	// modify
	for i := 0; i < len(ab); i++ {
		// unexpected fault address 0x51e390
		// fatal error: fault
		// [signal SIGSEGV: segmentation violation code=0x2 addr=0x51e390 pc=0x4f08ec]
		// ab[i] = 'f' // 不能修改
	}
	t.Logf("a: %s, ab: %s\n", a, ab)
}

func TestBytesToStringModify(t *testing.T) {
	a := []byte("hello")
	ab := BytesToString(a)
	t.Logf("a: %s, ab: %s\n", a, ab)

	// loop
	for i := 0; i < len(ab); i++ {
		t.Logf("ab[i]: %v\n", ab[i])
	}

	// 字符串本身就是不能通过下标修改的
}
