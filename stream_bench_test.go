package inverseflow

import (
	"testing"
)

func BenchmarkInverseStream(b *testing.B) {
	s := NewStream(addFunc(0))
	for n := 0; n < b.N; n++ {
		s.Inverse()
	}
}

func BenchmarkInverseNative(b *testing.B) {
	f := addFunc(0)
	for n := 0; n < b.N; n++ {
		f()
	}
}
