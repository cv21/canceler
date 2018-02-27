package canceler

import "testing"

type testCanceler struct {
	v int
}

func (t testCanceler) Cancel() error {
	t.v++
	return nil
}

func BenchmarkCancelCanceler(b *testing.B) {
	canceler := NewPool(testCanceler{})
	for n := 0; n < b.N; n++ {
		canceler.Cancel()
	}
}

func BenchmarkCancelNative(b *testing.B) {
	v := uint(0)

	f := addFunc(&v)

	for n := 0; n < b.N; n++ {
		f()
	}
}