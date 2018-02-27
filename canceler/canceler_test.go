package canceler

import (
	"testing"
	"github.com/stretchr/testify/assert"
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

func TestCancel(t *testing.T) {
	var cnt uint
	cm1 := &CancelerMock{CancelFunc: addFunc(&cnt)}
	cm2 := &CancelerMock{CancelFunc: addFunc(&cnt)}
	cm3 := &CancelerMock{CancelFunc: addFunc(&cnt)}

	p := NewPool(cm1)

	p.Add(cm2, cm3)
	p.Cancel()
	assert.Equal(t, cnt, uint(3), "canceler called all functions")
	assert.Equal(t, len(cm1.calls.Cancel), 1)
	assert.Equal(t, len(cm2.calls.Cancel), 1)
	assert.Equal(t, len(cm3.calls.Cancel), 1)
}

func addFunc(v *uint) func() error {
	return func() error {
		*v++
		return nil
	}
}