package cc

import (
	"sync"
)

// Future wrap value which can be Wait to get.
type Future[T any] struct {
	retChan chan interface{}
	val     T
	flag    bool
	lock    sync.Mutex
	e       error
}

func (f *Future[T]) waitUnify() ([]any, error) {
	val, err := f.Wait()
	return []any{val}, err
}

// Wait and return the value when it is ready,or else blocked
func (f *Future[T]) Wait() (T, error) {
	f.lock.Lock()
	if f.flag {
		f.lock.Unlock()
		return f.val, f.e
	}
	f.val, _ = (<-f.retChan).(T)
	f.e, _ = (<-f.retChan).(error)
	f.flag = true
	f.lock.Unlock()
	return f.val, f.e
}

type worker[T any] func() (T, error)

// Go run function in a new goroutine. result value wrapped in Future
func Go[T any](w worker[T]) *Future[T] {
	var f = &Future[T]{
		retChan: make(chan interface{}, 2),
	}
	go runFuture(f, w)
	return f
}

func runFuture[T any](f *Future[T], w worker[T]) {
	val, err := w()
	f.retChan <- val
	f.retChan <- err
	close(f.retChan)
}

// T2 wrap a pair of values.
type T2[T, V any] struct {
	V1 T
	V2 V
}

// Future2 wrap 2 values which can be Wait to get.
type Future2[T, V any] struct {
	f *Future[any]
}

func (f *Future2[T, V]) waitUnify() ([]any, error) {
	val1, val2, err := f.Wait()
	return []any{val1, val2}, err
}

// Wait and return the values when they are ready,or else blocked
func (f *Future2[T, V]) Wait() (T, V, error) {
	t, err := f.f.Wait()
	t2 := t.(T2[T, V])
	return t2.V1, t2.V2, err
}

// Go2 run function in a new goroutine.2 result values wrapped in Future2
func Go2[T, V any](f func() (T, V, error)) *Future2[T, V] {
	ret := Go(func() (any, error) {
		v1, v2, err := f()
		var ret T2[T, V]
		ret.V1 = v1
		ret.V2 = v2
		return ret, err
	})
	var ret2 Future2[T, V]
	ret2.f = ret
	return &ret2
}

// T3 wrap 3 values.
type T3[T, V, M any] struct {
	V1 T
	V2 V
	V3 M
}

// Future3 wrap 3 values which can be Wait to get.
type Future3[T, V, M any] struct {
	f *Future[any]
}

func (f *Future3[T, V, M]) waitUnify() ([]any, error) {
	val1, val2, val3, err := f.Wait()
	return []any{val1, val2, val3}, err
}

// Wait and return the values when they are ready,or else blocked
func (f *Future3[T, V, M]) Wait() (T, V, M, error) {
	t, err := f.f.Wait()
	t3 := t.(T3[T, V, M])
	return t3.V1, t3.V2, t3.V3, err
}

// Go3 run function in a new goroutine. 3 result values wrapped in Future3
func Go3[T, V, M any](f func() (T, V, M, error)) *Future3[T, V, M] {
	ret := Go(func() (any, error) {
		v1, v2, v3, err := f()
		var ret T3[T, V, M]
		ret.V1 = v1
		ret.V2 = v2
		ret.V3 = v3
		return ret, err
	})
	var ret3 Future3[T, V, M]
	ret3.f = ret
	return &ret3
}

// T4 wrap 4 values.
type T4[T, V, M, N any] struct {
	V1 T
	V2 V
	V3 M
	V4 N
}

// Future4 wrap 4 values which can be Wait to get.
type Future4[T, V, M, N any] struct {
	f *Future[any]
}

func (f *Future4[T, V, M, N]) waitUnify() ([]any, error) {
	val1, val2, val3, val4, err := f.Wait()
	return []any{val1, val2, val3, val4}, err
}

// Wait and return the values when they are ready,or else blocked
func (f *Future4[T, V, M, N]) Wait() (T, V, M, N, error) {
	t, err := f.f.Wait()
	t4 := t.(T4[T, V, M, N])
	return t4.V1, t4.V2, t4.V3, t4.V4, err
}

// Go4 run function in a new goroutine. 4 result values wrapped in Future4
func Go4[T, V, M, N any](f func() (T, V, M, N, error)) *Future4[T, V, M, N] {
	ret := Go(func() (any, error) {
		v1, v2, v3, v4, err := f()
		var ret T4[T, V, M, N]
		ret.V1 = v1
		ret.V2 = v2
		ret.V3 = v3
		ret.V4 = v4
		return ret, err
	})
	var ret4 Future4[T, V, M, N]
	ret4.f = ret
	return &ret4
}

// T5 wrap 5 values.
type T5[T, V, M, N, O any] struct {
	V1 T
	V2 V
	V3 M
	V4 N
	V5 O
}

// Future5 wrap 5 values which can be Wait to get.
type Future5[T, V, M, N, O any] struct {
	f *Future[any]
}

func (f *Future5[T, V, M, N, O]) waitUnify() ([]any, error) {
	val1, val2, val3, val4, val5, err := f.Wait()
	return []any{val1, val2, val3, val4, val5}, err
}

// Wait and return the values when they are ready,or else blocked
func (f *Future5[T, V, M, N, O]) Wait() (T, V, M, N, O, error) {
	t, err := f.f.Wait()
	t5 := t.(T5[T, V, M, N, O])
	return t5.V1, t5.V2, t5.V3, t5.V4, t5.V5, err
}

// Go5 run function in a new goroutine. 5 result values wrapped in Future5
func Go5[T, V, M, N, O any](f func() (T, V, M, N, O, error)) *Future5[T, V, M, N, O] {
	ret := Go(func() (any, error) {
		v1, v2, v3, v4, v5, err := f()
		var ret T5[T, V, M, N, O]
		ret.V1 = v1
		ret.V2 = v2
		ret.V3 = v3
		ret.V4 = v4
		ret.V5 = v5
		return ret, err
	})
	var ret5 Future5[T, V, M, N, O]
	ret5.f = ret
	return &ret5
}
