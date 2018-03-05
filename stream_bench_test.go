package inverseflow

import (
	"testing"
)

func BenchmarkCancelCanceler(b *testing.B) {
	canceler := NewStream(addFunc(0))
	for n := 0; n < b.N; n++ {
		canceler.Cancel()
	}
}

func BenchmarkCancelNative(b *testing.B) {
	f := addFunc(0)
	for n := 0; n < b.N; n++ {
		f()
	}
}
