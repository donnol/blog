package generic

// https://github.com/golang/go/issues/50281

type byteseq interface {
	string | []byte
}

// This should allow to eliminate the two functions above.
func AppendByteString[source byteseq](buf []byte, s source) []byte {
	// cannot use s[1:6] (value of type source constrained by byteseq) as type []byte in argument to append:
	// cannot assign string (in source) to []byte
	// return append(buf, s[1:6]...)

	// work
	// 但是要手动转型
	return append(buf, []byte(s[1:6])...)
}
