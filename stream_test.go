package inverseflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPool(t *testing.T) {
	p := NewStream(addFunc(0), addFunc(0))
	assert.Equal(t, len(p.funcs), 2, "inverseflow slice length is right")
}

func TestAdd(t *testing.T) {
	p := NewStream(addFunc(0))
	p.Add(addFunc(0))
	p.Add(addFunc(0))
	assert.Equal(t, len(p.funcs), 3, "inverseflow slice length is right")
}

func TestCancel(t *testing.T) {
	var cnt uint
	p := NewStream(addPtrFunc(&cnt))
	p.Add(addPtrFunc(&cnt))
	p.Add(addPtrFunc(&cnt))
	p.Cancel()
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
