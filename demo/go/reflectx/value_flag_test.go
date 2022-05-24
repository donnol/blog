package reflectx

import "testing"

func TestValueFlag(t *testing.T) {
	t.Logf("%b\n", flagKindMask) // 11111
	t.Logf("%b\n", flagStickyRO) // 100000
	t.Logf("%b\n", flagEmbedRO)  // 1000000
	t.Logf("%b\n", flagIndir)    // 10000000
	t.Logf("%b\n", flagAddr)     // 100000000
	t.Logf("%b\n", flagMethod)   // 1000000000
	t.Logf("%b\n", flagRO)       // 1100000
}
