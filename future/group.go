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
}

// Run a worker
func (g *Group[T]) Run(w worker[T]) {
	if g.sealed.Load() {
		panic("group is sealed")
	}
	g.ff = append(g.ff, Go(w))
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
