package fmtx

import "fmt"

var (
	Enable = false
)

func Printf(format string, args ...interface{}) {
	if !Enable {
		return
	}
	fmt.Printf(format, args...)
}
