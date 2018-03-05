package inverseflow

import "errors"

// stream holds a slice of Func and indexes for properly iterating.
type stream struct {
	funcs       map[Index]Func
	insertIndex Index
	cancelIndex Index
}

// Func is a main function type of inverseflow.
// Elements with such type calls when you call stream's Cancel method.
type Func func() error

// Index is a internal type for comfortable usage of indexes outside of the inverseflow.
type Index uint32

var ErrFuncNotFound = errors.New("function not found")

// Creates new inverseflow stream.
// Could recieve inverseflow Func list for rapid initialization.
func NewStream(fs ...Func) *stream {
	p := &stream{funcs: make(map[Index]Func)}
	for _, f := range fs {
		p.Add(f)
	}
	return p
}

// Add new Func to inverseflow stream.
// Returns an index of stream Func which can be used for removing Func from stream.
func (p *stream) Add(f Func) Index {
	i := p.insertIndex
	p.funcs[i] = f
	p.insertIndex++
	return i
}

// Remove func from inverseflow stream by index.
// It is not common use case, but might be helpful for any purposes.
func (p *stream) Remove(i Index) error {
	_, ok := p.funcs[i]
	if !ok {
		return ErrFuncNotFound
	}

	delete(p.funcs, i)
	return nil
}

// Cancels inverseflow stream.
// Stream could be canceled iteratively by handling errors and call Cancel method again.
// Stream can hold a cancel progress.
func (p *stream) Cancel() error {
	var err error
	for ; p.cancelIndex < p.insertIndex; p.cancelIndex++ {
		if v, ok := p.funcs[p.cancelIndex]; ok {
			if err = v(); err != nil {
				return err
			}
		}
	}

	return nil
}
