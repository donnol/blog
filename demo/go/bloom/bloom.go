package bloom

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

func Bloom(m, n uint, data []byte) error {
	filter := bloom.New(m, n)

	filter.Add(data)

	if !filter.Test(data) {
		return fmt.Errorf("Test data %s failed", data)
	}

	return nil
}
