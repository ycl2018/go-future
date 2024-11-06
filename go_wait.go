package cc

import "sync"

// Future wrap value which can be Wait to get.
type Future[T any] struct {
	retChan chan interface{}
	val     T
	flag    bool
	lock    sync.Mutex
	e       error
}

// Wait and return the value when it is ready,or else blocked
func (r *Future[T]) Wait() (T, error) {
	r.lock.Lock()
	if r.flag {
		r.lock.Unlock()
		return r.val, r.e
	}
	r.val, _ = (<-r.retChan).(T)
	r.e, _ = (<-r.retChan).(error)
	r.flag = true
	r.lock.Unlock()
	return r.val, r.e
}

// Go run function in a new goroutine. result value wrapped in Future
func Go[T any](f func() (value T, err error)) *Future[T] {
	var r = &Future[T]{
		retChan: make(chan interface{}, 2),
	}
	go func() {
		val, err := f()
		r.retChan <- val
		r.retChan <- err
		close(r.retChan)
	}()
	return r
}

// T2 wrap a pair of values.
type T2[T, V any] struct {
	V1 T
	V2 V
}

// Future2 wrap 2 values which can be Wait to get.
type Future2[T, V any] struct {
	r *Future[T2[T, V]]
}

// Wait and return the values when they are ready,or else blocked
func (t *Future2[T, V]) Wait() (T, V, error) {
	t2, err := t.r.Wait()
	return t2.V1, t2.V2, err
}

// Go2 run function in a new goroutine.2 result values wrapped in Future2
func Go2[T, V any](f func() (T, V, error)) *Future2[T, V] {
	ret := Go(func() (T2[T, V], error) {
		v1, v2, err := f()
		var ret T2[T, V]
		ret.V1 = v1
		ret.V2 = v2
		return ret, err
	})
	var ret2 Future2[T, V]
	ret2.r = ret
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
	r *Future[T3[T, V, M]]
}

// Wait and return the values when they are ready,or else blocked
func (t *Future3[T, V, M]) Wait() (T, V, M, error) {
	t3, err := t.r.Wait()
	return t3.V1, t3.V2, t3.V3, err
}

// Go3 run function in a new goroutine. 3 result values wrapped in Future3
func Go3[T, V, M any](f func() (T, V, M, error)) *Future3[T, V, M] {
	ret := Go(func() (T3[T, V, M], error) {
		v1, v2, v3, err := f()
		var ret T3[T, V, M]
		ret.V1 = v1
		ret.V2 = v2
		ret.V3 = v3
		return ret, err
	})
	var ret3 Future3[T, V, M]
	ret3.r = ret
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
	r *Future[T4[T, V, M, N]]
}

// Wait and return the values when they are ready,or else blocked
func (t *Future4[T, V, M, N]) Wait() (T, V, M, N, error) {
	t4, err := t.r.Wait()
	return t4.V1, t4.V2, t4.V3, t4.V4, err
}

// Go4 run function in a new goroutine. 4 result values wrapped in Future4
func Go4[T, V, M, N any](f func() (T, V, M, N, error)) *Future4[T, V, M, N] {
	ret := Go(func() (T4[T, V, M, N], error) {
		v1, v2, v3, v4, err := f()
		var ret T4[T, V, M, N]
		ret.V1 = v1
		ret.V2 = v2
		ret.V3 = v3
		ret.V4 = v4
		return ret, err
	})
	var ret4 Future4[T, V, M, N]
	ret4.r = ret
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
	r *Future[T5[T, V, M, N, O]]
}

// Wait and return the values when they are ready,or else blocked
func (t *Future5[T, V, M, N, O]) Wait() (T, V, M, N, O, error) {
	t5, err := t.r.Wait()
	return t5.V1, t5.V2, t5.V3, t5.V4, t5.V5, err
}

// Go5 run function in a new goroutine. 5 result values wrapped in Future5
func Go5[T, V, M, N, O any](f func() (T, V, M, N, O, error)) *Future5[T, V, M, N, O] {
	ret := Go(func() (T5[T, V, M, N, O], error) {
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
	ret5.r = ret
	return &ret5
}
