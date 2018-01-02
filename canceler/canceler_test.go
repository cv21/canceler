package canceler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNewPool(t *testing.T) {
	p := NewPool(&CancelerMock{}, &CancelerMock{})
	assert.Equal(t, len(p.cancelers), 2, "canceler slice length is right")
}

func TestAdd(t *testing.T) {
	p := NewPool(&CancelerMock{})
	p.Add(&CancelerMock{}, &CancelerMock{})
	assert.Equal(t, len(p.cancelers), 3, "canceler slice length is right")
}

// TODO finish testing cancel method here
func TestCancel(t *testing.T) {
	p := NewPool(&CancelerMock{CancelFunc: func() error {

	}})
	p.Add(&CancelerMock{}, &CancelerMock{})
	p.Cancel()
}