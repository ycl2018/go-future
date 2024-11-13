package future

import (
	"errors"
	"time"
)

var ErrAllFailed = errors.New("[Future] err: all futures failed")

// -------------------- Collect 2 Values --------------------

// Collect 2 futures, wait and return joined values and errors.
func Collect[T, V any](f1 *Future[T], f2 *Future[V]) (T, V, error) {
	v1, err1 := f1.Wait()
	v2, err2 := f2.Wait()
	err := errors.Join(err1, err2)
	return v1, v2, err
}

// -------------------- Collect 3 Values --------------------

// Collect3 futures, wait and return joined values and errors.
func Collect3[T, V, M any](f1 *Future[T], f2 *Future[V], f3 *Future[M]) (T, V, M, error) {
	v1, v2, cErr := Collect(f1, f2)
	v3, err := f3.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, err
}

// Collect1u2 combine Future with Future2, wait and return joined values and errors.
func Collect1u2[T, V, M any](f1 *Future[T], f2 *Future2[V, M]) (T, V, M, error) {
	v1, err := f1.Wait()
	v2, v3, err := f2.Wait()
	return v1, v2, v3, err
}

// -------------------- Collect 4 Values --------------------

// Collect4 futures, wait and return joined values and errors.
func Collect4[T, V, M, N any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N]) (T, V, M, N, error) {
	v1, v2, v3, cErr := Collect3(f1, f2, f3)
	v4, err := f4.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, err
}

// Collect1u3 combine Future with Future3, wait and return joined values and errors.
func Collect1u3[T, V, M, N any](f1 *Future[T], f2 *Future3[V, M, N]) (T, V, M, N, error) {
	v1, err := f1.Wait()
	v2, v3, v4, err := f2.Wait()
	return v1, v2, v3, v4, err
}

// Collect2x1u2 combine 2 Futures with Future2, wait and return joined values and errors.
func Collect2x1u2[T, V, M, N any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N]) (T, V, M, N, error) {
	v1, v2, err1 := Collect(f1, f2)
	v3, v4, err2 := f3.Wait()
	return v1, v2, v3, v4, errors.Join(err1, err2)
}

// Collect2x2 combine Future2 with Future2, wait and return joined values and errors.
func Collect2x2[T, V, M, N any](f1 *Future2[T, V], f2 *Future2[M, N]) (T, V, M, N, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, err2 := f2.Wait()
	return v1, v2, v3, v4, errors.Join(err1, err2)
}

// -------------------- Collect 5 Values --------------------

// Collect5 futures, wait and return joined values and errors.
func Collect5[T, V, M, N, O any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O]) (
	T, V, M, N, O, error) {
	v1, v2, v3, v4, cErr := Collect4(f1, f2, f3, f4)
	v5, err := f5.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, err
}

// Collect1u4 combine Future with Future4, wait and return joined values and errors.
func Collect1u4[T, V, M, N, O any](f1 *Future[T], f2 *Future4[V, M, N, O]) (T, V, M, N, O, error) {
	v1, err1 := f1.Wait()
	v2, v3, v4, v5, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Collect1u2x2 combine Future with 2 Future2(s), wait and return joined values and errors.
func Collect1u2x2[T, V, M, N, O any](f1 *Future[T], f2 *Future2[V, M], f3 *Future2[N, O]) (T, V, M, N, O, error) {
	v1, err1 := f1.Wait()
	v2, v3, v4, v5, err2 := Collect2x2(f2, f3)
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Collect2x1u3 combine 2 Future(s) with Future3, wait and return joined values and errors.
func Collect2x1u3[T, V, M, N, O any](f1 *Future[T], f2 *Future[V], f3 *Future3[M, N, O]) (T, V, M, N, O, error) {
	v1, v2, err1 := Collect(f1, f2)
	v3, v4, v5, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Collect3x1u2 combine 2 Future(s) with Future3, wait and return joined values and errors.
func Collect3x1u2[T, V, M, N, O any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future2[N, O]) (T, V, M, N, O, error) {
	v1, v2, v3, err1 := Collect3(f1, f2, f3)
	v4, v5, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Collect2u3 combine Future2 with Future3, wait and return joined values and errors.
func Collect2u3[T, V, M, N, O any](f1 *Future2[T, V], f2 *Future3[M, N, O]) (T, V, M, N, O, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, v5, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// -------------------- Collect 6 Values --------------------

// Collect6 futures, wait and return joined values and errors.
func Collect6[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, v5, cErr := Collect5(f1, f2, f3, f4, f5)
	v6, err := f6.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, err
}

// Collect1u5 combine Future with Future5, wait and return joined values and errors.
func Collect1u5[T, V, M, N, O, P any](f1 *Future[T], f2 *Future5[V, M, N, O, P]) (T, V, M, N, O, P, error) {
	v1, err1 := f1.Wait()
	v2, v3, v4, v5, v6, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect1u2u3 combine Future, Future2 with Future3, wait and return joined values and errors.
func Collect1u2u3[T, V, M, N, O, P any](f1 *Future[T], f2 *Future2[V, M], f3 *Future3[N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, err1 := Collect1u2(f1, f2)
	v4, v5, v6, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect2x1u4 combine 2 Future(s) with Future4, wait and return joined values and errors.
func Collect2x1u4[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future4[M, N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, err1 := Collect(f1, f2)
	v3, v4, v5, v6, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect2x1u2x2 combine 2 Future(s) with 2 Future2(s), wait and return joined values and errors.
func Collect2x1u2x2[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 *Future2[O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, err1 := Collect2x1u2(f1, f2, f3)
	v5, v6, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect3x1u3 combine 3 Future(s) with Future3, wait and return joined values and errors.
func Collect3x1u3[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future3[N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, err1 := Collect3(f1, f2, f3)
	v4, v5, v6, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect4x1u2 combine 4 Future(s) with Future2, wait and return joined values and errors.
func Collect4x1u2[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future2[O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, err1 := Collect4(f1, f2, f3, f4)
	v5, v6, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect2u4 combine Future2 with Future4, wait and return joined values and errors.
func Collect2u4[T, V, M, N, O, P any](f1 *Future2[T, V], f2 *Future4[M, N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, v5, v6, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect3x2 combine 3 Future2(s), wait and return joined values and errors.
func Collect3x2[T, V, M, N, O, P any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future2[O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, err1 := Collect2x2(f1, f2)
	v5, v6, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Collect2x3 combine 2 Future3(s), wait and return joined values and errors.
func Collect2x3[T, V, M, N, O, P any](f1 *Future3[T, V, M], f2 *Future3[N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, err1 := f1.Wait()
	v4, v5, v6, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// -------------------- Collect 7 Values --------------------

// Collect7 Future(s), wait and return joined values and errors.
func Collect7[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, v6, cErr := Collect6(f1, f2, f3, f4, f5, f6)
	v7, err := f7.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, v7, err
}

// Collect1u2u4 combine Future, Future2 with Future4, wait and return joined values and errors.
func Collect1u2u4[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future2[V, M], f3 *Future4[N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, err1 := Collect1u2(f1, f2)
	v4, v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect1u2x3 combine Future with 2 Future3(s), wait and return joined values and errors.
func Collect1u2x3[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future3[V, M, N], f3 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Collect1u3(f1, f2)
	v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect1u3x2 combine Future with 3 Future2(s), wait and return joined values and errors.
func Collect1u3x2[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future2[V, M], f3 *Future2[N, O], f4 Future2[P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, err1 := Collect1u2x2(f1, f2, f3)
	v6, v7, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect2x1u5 combine Future with Future5, wait and return joined values and errors.
func Collect2x1u5[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future5[M, N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, err1 := Collect(f1, f2)
	v3, v4, v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect2x1u2u3 combine 2 Future(s), Future2 with Future3, wait and return joined values and errors.
func Collect2x1u2u3[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Collect2x1u2(f1, f2, f3)
	v5, v6, v7, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect3x1u4 combine 3 Future(s) with Future4, wait and return joined values and errors.
func Collect3x1u4[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future4[N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, err1 := Collect3(f1, f2, f3)
	v4, v5, v6, v7, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect3x1u2x2 combine 3 Future(s) with Future4, wait and return joined values and errors.
func Collect3x1u2x2[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future2[N, O], f5 *Future2[P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, err1 := Collect3x1u2(f1, f2, f3, f4)
	v6, v7, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect4x1u3 combine 4 Future(s) with Future3, wait and return joined values and errors.
func Collect4x1u3[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Collect4(f1, f2, f3, f4)
	v5, v6, v7, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect5x1u2 combine 5 Future(s) with Future2, wait and return joined values and errors.
func Collect5x1u2[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O], f6 *Future2[P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, err1 := Collect5(f1, f2, f3, f4, f5)
	v6, v7, err2 := f6.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect2u5 combine Future2 with Future5, wait and return joined values and errors.
func Collect2u5[T, V, M, N, O, P, Q any](f1 *Future2[T, V], f2 *Future5[M, N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, v5, v6, v7, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect2x2u3 combine 2 Future2 with Future3, wait and return joined values and errors.
func Collect2x2u3[T, V, M, N, O, P, Q any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Collect2x2(f1, f2)
	v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Collect3u4 combine Future3 with Future4, wait and return joined values and errors.
func Collect3u4[T, V, M, N, O, P, Q any](f1 *Future3[T, V, M], f2 *Future4[N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, err1 := f1.Wait()
	v4, v5, v6, v7, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// -------------------- Collect 8 Values --------------------

// Collect8 futures, wait and return joined values and errors.
func Collect8[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q], f8 *Future[R]) (T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, v7, err1 := Collect7(f1, f2, f3, f4, f5, f6, f7)
	v8, err2 := f8.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect2x1u2u4 combine 2 Future(s) with Future2 and Future4, wait and return joined values and errors.
func Collect2x1u2u4[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 Future4[O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, err1 := Collect2x1u2(f1, f2, f3)
	v5, v6, v7, v8, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect2x1u3x2 combine 2 Future(s) with 3 Future2, wait and return joined values and errors.
func Collect2x1u3x2[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 *Future2[O, P], f5 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Collect2x1u2x2(f1, f2, f3, f4)
	v7, v8, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect2x1u2x3 combine 2 Future(s) with 2 Future3, wait and return joined values and errors.
func Collect2x1u2x3[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future3[M, N, O], f4 *Future3[P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, err1 := Collect2x1u3(f1, f2, f3)
	v6, v7, v8, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect3x1u2u3 combine 3 Future(s) with Future2 and Future3, wait and return joined values and errors.
func Collect3x1u2u3[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future2[N, O], f5 *Future3[P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, err1 := Collect3x1u2(f1, f2, f3, f4)
	v6, v7, v8, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect4x1u2x2 combine 4 Future(s) with 2 Future2(s), wait and return joined values and errors.
func Collect4x1u2x2[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future2[O, P], f6 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Collect4x1u2(f1, f2, f3, f4, f5)
	v7, v8, err2 := f6.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect5x1u3 combine 5 Future(s) with Future3, wait and return joined values and errors.
func Collect5x1u3[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O], f6 *Future3[P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, err1 := Collect5(f1, f2, f3, f4, f5)
	v6, v7, v8, err2 := f6.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect6x1u2 combine 6 Future(s) with Future2, wait and return joined values and errors.
func Collect6x1u2[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O], f6 *Future[P], f7 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Collect6(f1, f2, f3, f4, f5, f6)
	v7, v8, err2 := f7.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect2x2u4 combine 2 Future2(s) with Future4, wait and return joined values and errors.
func Collect2x2u4[T, V, M, N, O, P, Q, R any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future4[O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, err1 := Collect2x2(f1, f2)
	v5, v6, v7, v8, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect4x2 combine 4 Future2(s), wait and return joined values and errors.
func Collect4x2[T, V, M, N, O, P, Q, R any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future2[O, P], f4 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Collect3x2(f1, f2, f3)
	v7, v8, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect3u5 combine Future3 with Future5, wait and return joined values and errors.
func Collect3u5[T, V, M, N, O, P, Q, R any](f1 *Future3[T, V, M], f2 *Future5[N, O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, err1 := f1.Wait()
	v4, v5, v6, v7, v8, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect2x3u2 combine 2 Future3(s) with Future2, wait and return joined values and errors.
func Collect2x3u2[T, V, M, N, O, P, Q, R any](f1 *Future3[T, V, M], f2 *Future3[N, O, P], f3 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Collect2x3(f1, f2)
	v7, v8, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Collect2x4 combine 2 Future4(s), wait and return joined values and errors.
func Collect2x4[T, V, M, N, O, P, Q, R any](f1 *Future4[T, V, M, N], f2 *Future4[O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, err1 := f1.Wait()
	v5, v6, v7, v8, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// -------------------- Collect more than 8 Values --------------------

// CollectSlice futures which return same type, wait and return all joined values and errors.
func CollectSlice[T any](fs ...*Future[T]) ([]T, error) {
	var err error
	var ret []T
	for _, f := range fs {
		v, curErr := f.Wait()
		err = errors.Join(curErr)
		ret = append(ret, v)
	}
	return ret, err
}

// CollectSliceTimeout combine futures which return same type, wait for timeout duration and return all values with joined errors,
// otherwise an ErrTimeout returned
func CollectSliceTimeout[T any](timeout time.Duration, fs ...*Future[T]) ([]T, error) {
	return Timeout(timeout, func() ([]T, error) {
		return CollectSlice(fs...)
	})
}

type futureI interface {
	waitUnify() ([]any, error)
	waitTimeout(time.Duration) ([]any, error)
}

// CollectAll futures, wait and return all values and joined errors.
func CollectAll(fs ...futureI) ([]any, error) {
	var err error
	var ret []any
	for _, f := range fs {
		v, curErr := f.waitUnify()
		err = errors.Join(curErr)
		ret = append(ret, v...)
	}
	return ret, err
}

// CollectAllTimeout combine futures, wait for timeout duration and return all values and joined errors,
// otherwise an ErrTimeout returned
func CollectAllTimeout(timeout time.Duration, fs ...futureI) ([]any, error) {
	return Timeout(timeout, func() ([]any, error) {
		return CollectAll(fs...)
	})
}

// CollectAny futures and return as soon as any one succeeds.
func CollectAny[T any](fs ...*Future[T]) (T, error) {
	for _, f := range fs {
		ret, err := f.Wait()
		if err == nil {
			return ret, nil
		}
	}
	var t T
	return t, ErrAllFailed
}

// CollectAnyTimeout wait for timeout duration and return as soon as any one of Futures succeeds
// otherwise an ErrTimeout returned.
func CollectAnyTimeout[T any](timeout time.Duration, fs ...*Future[T]) (T, error) {
	return Timeout(timeout, func() (T, error) {
		return CollectAny(fs...)
	})
}
