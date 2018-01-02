//go:generate moq -out canceler_mock.go . Canceler
package canceler

// Canceler provides Cancel method which can cancel appropriate action and return error.
type Canceler interface {
	Cancel() error
}

// Pool holds a slice of Canceler.
// Provides Add() and Cancel() methods for add cancelers and initiate cancel for each of holding cancelers.
type Pool struct {
	cancelers []Canceler
}

// Creates new pool of cancelers.
func NewPool(cs ...Canceler) *Pool {
	return &Pool{cancelers: cs}
}

// Add new cancelers to canceler pool.
func (p *Pool) Add(c ...Canceler) {
	p.cancelers = append(p.cancelers, c...)
}

// Cancels canceler pool.
// Call cancel method for each of pool cancelers.
func (p *Pool) Cancel() error {
	for _, c := range p.cancelers {
		if err := c.Cancel(); err != nil {
			return err
		}
	}

	return nil
}
