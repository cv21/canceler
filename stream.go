package inverse

import "errors"

type Stream interface {
	Add(f Func) Index
	Remove(i Index) error
	Inverse() error
}

// stream holds a slice of Func and indexes for properly iterating.
type stream struct {
	funcs       map[Index]Func
	insertIndex Index
	cancelIndex Index
}

// Func it is so called cancel func which calls when you call stream's Cancel method.
type Func func() error

// It is uses for indexing Func inside stream.
type Index uint32

var ErrFuncNotFound = errors.New("function not found")

// Creates new inverseflow stream.
// Could receive inverseflow Func list for rapid initialization.
// Initialized Funcs in this way not provides its indexes, so you could not remove them from stream.
func NewStream(fs ...Func) Stream {
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

// Inverse inverseflow stream.
// If error occured, stream can inverse iteratively by handling errors and call Inverse method again.
// Stream holds a inverse progress.
func (p *stream) Inverse() error {
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
