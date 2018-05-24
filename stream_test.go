package inverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInverse(t *testing.T) {
	var cnt uint
	p := NewStream(addPtrFunc(&cnt))
	p.Add(addPtrFunc(&cnt))
	p.Add(addPtrFunc(&cnt))
	p.Inverse()
	assert.Equal(t, cnt, uint(3), "inverseflow called all functions")
}

func addPtrFunc(v *uint) func() error {
	return func() error {
		*v++
		return nil
	}
}

func addFunc(v int) Func {
	return func() error {
		v++
		return nil
	}
}
