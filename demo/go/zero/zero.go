package zero

import "unsafe"

func Zero[T any](v T) bool {
	bp := (*byte)(unsafe.Pointer(&v))
	sz := unsafe.Sizeof(v)
	for _, v := range unsafe.Slice(bp, sz) {
		if v != 0 {
			return false
		}
	}
	return true
}
