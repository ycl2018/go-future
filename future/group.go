package future

import (
	"sync/atomic"
	"time"
)

// AnyGroup represents a group of Futures which returns any with an error
type AnyGroup = Group[any]

// Group represents a group of Futures which returns same type T with an error
type Group[T any] struct {
	ff     []*Future[T]
	sealed atomic.Bool
	limit  chan struct{}
}

// NewLimitGroup create a Group with limit parallelism
func NewLimitGroup[T any](parallel int) *Group[T] {
	if parallel <= 0 {
		panic("limit must be greater than 0")
	}
	return &Group[T]{limit: make(chan struct{}, parallel)}
}

// Run a worker
func (g *Group[T]) Run(w worker[T]) {
	if g.sealed.Load() {
		panic("group is sealed")
	}
	var f *Future[T]
	if cap(g.limit) > 0 {
		g.limit <- struct{}{}
		f = Go(w)
		f.Check(func(t T, err error) (T, error) {
			<-g.limit
			return t, err
		})
	} else {
		f = Go(w)
	}
	g.ff = append(g.ff, f)
}

// Add a Future to group
func (g *Group[T]) Add(f *Future[T]) {
	if g.sealed.Load() {
		panic("group is sealed")
	}
	g.ff = append(g.ff, f)
}

// Wait all the Futures return
func (g *Group[T]) Wait() ([]T, error) {
	g.sealed.Store(true)
	return CollectSlice(g.ff...)
}

// WaitTimeout wait for timeout duration to get all the worker return
func (g *Group[T]) WaitTimeout(timeout time.Duration) ([]T, error) {
	g.sealed.Store(true)
	return CollectSliceTimeout(timeout, g.ff...)
}

// ErrGroup represents a group of Futures which returns only an error
type ErrGroup struct {
	g Group[struct{}]
}

// Run a worker
func (e *ErrGroup) Run(w func() error) {
	e.g.Run(func() (struct{}, error) {
		err := w()
		return struct{}{}, err
	})
}

// Add a Future to group
func (e *ErrGroup) Add(f futureI) {
	e.g.Run(func() (struct{}, error) {
		_, err := f.waitUnify()
		return struct{}{}, err
	})
}

// Wait all the Futures return
func (e *ErrGroup) Wait() error {
	_, err := e.g.Wait()
	return err
}

// WaitTimeout wait for timeout duration to get all the worker return
func (e *ErrGroup) WaitTimeout(timeout time.Duration) error {
	_, err := e.g.WaitTimeout(timeout)
	return err
}
