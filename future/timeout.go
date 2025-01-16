package future

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("[Future] err:run timeout")

// WaitTimeout wait for timeout duration to get result otherwise an ErrTimeout will be returned.
func (f *Future[T]) WaitTimeout(timeout time.Duration) (T, error) {
	return Timeout(timeout, f.Wait)
}

// WaitTimeout wait for timeout duration to get result otherwise an ErrTimeout will be returned.
func (f *Future2[T, V]) WaitTimeout(timeout time.Duration) (T, V, error) {
	ret, err := f.f.WaitTimeout(timeout)
	if ret != nil {
		r := ret.(T2[T, V])
		return r.V1, r.V2, err
	}
	var r T2[T, V]
	return r.V1, r.V2, err
}

// WaitTimeout wait for timeout duration to get result otherwise an ErrTimeout will be returned.
func (f *Future3[T, V, M]) WaitTimeout(timeout time.Duration) (T, V, M, error) {
	ret, err := f.f.WaitTimeout(timeout)
	if ret != nil {
		r := ret.(T3[T, V, M])
		return r.V1, r.V2, r.V3, err
	}
	var r T3[T, V, M]
	return r.V1, r.V2, r.V3, err
}

// WaitTimeout wait for timeout duration to get result otherwise an ErrTimeout will be returned.
func (f *Future4[T, V, M, N]) WaitTimeout(timeout time.Duration) (T, V, M, N, error) {
	ret, err := f.f.WaitTimeout(timeout)
	if ret != nil {
		r := ret.(T4[T, V, M, N])
		return r.V1, r.V2, r.V3, r.V4, err
	}
	var r T4[T, V, M, N]
	return r.V1, r.V2, r.V3, r.V4, err
}

// WaitTimeout wait for timeout duration to get result otherwise an ErrTimeout will be returned.
func (f *Future5[T, V, M, N, O]) WaitTimeout(timeout time.Duration) (T, V, M, N, O, error) {
	ret, err := f.f.WaitTimeout(timeout)
	if ret != nil {
		r := ret.(T5[T, V, M, N, O])
		return r.V1, r.V2, r.V3, r.V4, r.V5, err
	}
	var r T5[T, V, M, N, O]
	return r.V1, r.V2, r.V3, r.V4, r.V5, err
}

// JoinTimeout wait for the duration of this Future t1 and the other t2,
// combine both returned values and return a new combined value Future2.
// if the Joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future[T]) JoinTimeout(other futureI, t1, t2 time.Duration) *Future2[T, any] {
	return Go2(func() (T, any, error) {
		val1, err1 := f.WaitTimeout(t1)
		val2, err2 := other.waitTimeout(t2)
		return val1, unwrap(val2), errors.Join(err1, err2)
	})
}

// JoinTimeout wait for the duration of this Future t1 and the other t2,
// combine both returned values and return a new combined value Future2.
// if the Joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future2[T, V]) JoinTimeout(other futureI, t1, t2 time.Duration) *Future3[T, V, any] {
	return Go3(func() (T, V, any, error) {
		val1, val2, err1 := f.WaitTimeout(t1)
		val3, err2 := other.waitTimeout(t2)
		return val1, val2, unwrap(val3), errors.Join(err1, err2)
	})
}

// JoinTimeout wait for the duration of this Future t1 and the other t2,
// combine both returned values and return a new combined value Future2.
// if the Joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future3[T, V, M]) JoinTimeout(other futureI, t1, t2 time.Duration) *Future4[T, V, M, any] {
	return Go4(func() (T, V, M, any, error) {
		val1, val2, val3, err1 := f.WaitTimeout(t1)
		val4, err2 := other.waitTimeout(t2)
		return val1, val2, val3, unwrap(val4), errors.Join(err1, err2)
	})
}

// JoinTimeout wait for the duration of this Future t1 and the other t2,
// combine both returned values and return a new combined value Future2.
// if the Joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future4[T, V, M, N]) JoinTimeout(other futureI, t1, t2 time.Duration) *Future5[T, V, M, N, any] {
	return Go5(func() (T, V, M, N, any, error) {
		val1, val2, val3, val4, err1 := f.WaitTimeout(t1)
		val5, err2 := other.waitTimeout(t2)
		return val1, val2, val3, val4, unwrap(val5), errors.Join(err1, err2)
	})
}

// JoinTimeout wait for the duration of this Future t1 and the other t2,
// combine both returned values and return a new combined value Future2.
// if the Joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future5[T, V, M, N, O]) JoinTimeout(other futureI, t1, t2 time.Duration) *Future[[6]any] {
	return Go(func() ([6]any, error) {
		val1, val2, val3, val4, val5, err1 := f.WaitTimeout(t1)
		val6, err2 := other.waitTimeout(t2)
		return [...]any{val1, val2, val3, val4, val5, unwrap(val6)}, errors.Join(err1, err2)
	})
}

// Timeout wait for 'timeout' duration to get the function f return,if not an ErrTimeout returned
func Timeout[T any](timeout time.Duration, f func() (T, error)) (T, error) {
	if timeout <= 0 {
		return f()
	}
	var getRet = make(chan T2[T, error], 1)
	go func() {
		ret, err := f()
		getRet <- T2[T, error]{ret, err}
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	
	select {
	case <-timer.C:
		var t T
		return t, ErrTimeout
	case ret := <-getRet:
		return ret.V1, ret.V2
	}
}
