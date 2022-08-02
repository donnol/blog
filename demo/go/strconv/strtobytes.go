package strconv

import (
	"reflect"
	"unsafe"
)

// 字符串转为[]byte后，
// 得到的结果只读，如果用下标修改会panic

// StringToBytes converts string to byte slice without a memory allocation.
// 此版本不需要引入reflect
// BenchmarkBytesConvStrToBytes-4 1000000000 0.2451 ns/op 0 B/op 0 allocs/op
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string // Data & Len ; 使用string就不需要引入reflect
			Cap    int
		}{s, len(s)},
	))
}

// 此版本需要引入reflect
// 并且性能没有上面的方法高
// BenchmarkBytesConvStrToBytes0-4 549816554 2.189 ns/op 0 B/op 0 allocs/op
func StringToBytes0(s string) []byte {
	if s == "" {
		return nil
	}
	return unsafe.Slice(
		(*byte)(unsafe.Pointer(
			(*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
		)),
		len(s),
	)
}

// func StringToBytes1(s string) []byte {
// 	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
// 	bh := reflect.SliceHeader{
// 		Data: sh.Data,
// 		Len:  sh.Len,
// 		Cap:  sh.Len,
// 	}
// 	return *(*[]byte)(unsafe.Pointer(&bh))
// }

// func StringToBytes2(s string) (b []byte) {
// 	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
// 	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
// 	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
// 	return b
// }

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
