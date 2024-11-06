package cc

import "errors"

// Combine 2 futures, wait and return values with joined errors.
func Combine[T, V any](f1 *Future[T], f2 *Future[V]) (T, V, error) {
	v1, err1 := f1.Wait()
	v2, err2 := f2.Wait()
	err := errors.Join(err1, err2)
	return v1, v2, err
}

// Combine3 futures, wait and return values with joined errors.
func Combine3[T, V, M any](f1 *Future[T], f2 *Future[V], f3 *Future[M]) (T, V, M, error) {
	v1, v2, cErr := Combine(f1, f2)
	v3, err := f3.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, err
}

// Combine4 futures, wait and return values with joined errors.
func Combine4[T, V, M, N any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N]) (T, V, M, N, error) {
	v1, v2, v3, cErr := Combine3(f1, f2, f3)
	v4, err := f4.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, err
}

// Combine5 futures, wait and return values with joined errors.
func Combine5[T, V, M, N, O any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O]) (
	T, V, M, N, O, error) {
	v1, v2, v3, v4, cErr := Combine4(f1, f2, f3, f4)
	v5, err := f5.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, err
}

// Combine6 futures, wait and return values with joined errors.
func Combine6[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, v5, cErr := Combine5(f1, f2, f3, f4, f5)
	v6, err := f6.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, err
}

// Combine7 futures, wait and return values with joined errors.
func Combine7[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, v6, cErr := Combine6(f1, f2, f3, f4, f5, f6)
	v7, err := f7.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, v7, err
}

// Combine8 futures, wait and return values with joined errors.
func Combine8[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q], f8 *Future[R]) (T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, v7, cErr := Combine7(f1, f2, f3, f4, f5, f6, f7)
	v8, err := f8.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, v7, v8, err
}

// Combine9 futures, wait and return values with joined errors.
func Combine9[T, V, M, N, O, P, Q, R, S any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q], f8 *Future[R], f9 *Future[S]) (T, V, M, N, O, P, Q, R, S, error) {
	v1, v2, v3, v4, v5, v6, v7, v8, cErr := Combine8(f1, f2, f3, f4, f5, f6, f7, f8)
	v9, err := f9.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9, err
}

// Combine10 futures, wait and return values with joined errors.
func Combine10[T, V, M, N, O, P, Q, R, S, U any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q], f8 *Future[R], f9 *Future[S], f10 *Future[U]) (T, V, M, N, O, P, Q, R, S, U, error) {
	v1, v2, v3, v4, v5, v6, v7, v8, v9, cErr := Combine9(f1, f2, f3, f4, f5, f6, f7, f8, f9)
	v10, err := f10.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, err
}

// CombineN futures which return same type, wait and return values with joined errors.
func CombineN[T any](fs ...*Future[T]) ([]T, error) {
	var err error
	var ret []T
	for _, f := range fs {
		v, curErr := f.Wait()
		err = errors.Join(curErr)
		ret = append(ret, v)
	}
	return ret, err
}

type futureI interface {
	waitUnify() ([]any, error)
}

// CombineAll futures, wait and return values by any slice with joined errors.
func CombineAll(fs ...futureI) ([]any, error) {
	var err error
	var ret []any
	for _, f := range fs {
		v, curErr := f.waitUnify()
		err = errors.Join(curErr)
		ret = append(ret, v...)
	}
	return ret, err
}
