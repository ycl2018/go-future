package future

import (
	"errors"
	"time"
)

var ErrAllFailed = errors.New("[Future] err: all futures failed")

// -------------------- Combine 2 Values --------------------

// Combine 2 futures, wait and return joined values and errors.
func Combine[T, V any](f1 *Future[T], f2 *Future[V]) (T, V, error) {
	v1, err1 := f1.Wait()
	v2, err2 := f2.Wait()
	err := errors.Join(err1, err2)
	return v1, v2, err
}

// -------------------- Combine 3 Values --------------------

// Combine3 futures, wait and return joined values and errors.
func Combine3[T, V, M any](f1 *Future[T], f2 *Future[V], f3 *Future[M]) (T, V, M, error) {
	v1, v2, cErr := Combine(f1, f2)
	v3, err := f3.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, err
}

// Combine1u2 combine Future with Future2, wait and return joined values and errors.
func Combine1u2[T, V, M any](f1 *Future[T], f2 *Future2[V, M]) (T, V, M, error) {
	v1, err := f1.Wait()
	v2, v3, err := f2.Wait()
	return v1, v2, v3, err
}

// -------------------- Combine 4 Values --------------------

// Combine4 futures, wait and return joined values and errors.
func Combine4[T, V, M, N any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N]) (T, V, M, N, error) {
	v1, v2, v3, cErr := Combine3(f1, f2, f3)
	v4, err := f4.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, err
}

// Combine1u3 combine Future with Future3, wait and return joined values and errors.
func Combine1u3[T, V, M, N any](f1 *Future[T], f2 *Future3[V, M, N]) (T, V, M, N, error) {
	v1, err := f1.Wait()
	v2, v3, v4, err := f2.Wait()
	return v1, v2, v3, v4, err
}

// Combine2x1u2 combine 2 Futures with Future2, wait and return joined values and errors.
func Combine2x1u2[T, V, M, N any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N]) (T, V, M, N, error) {
	v1, v2, err1 := Combine(f1, f2)
	v3, v4, err2 := f3.Wait()
	return v1, v2, v3, v4, errors.Join(err1, err2)
}

// Combine2x2 combine Future2 with Future2, wait and return joined values and errors.
func Combine2x2[T, V, M, N any](f1 *Future2[T, V], f2 *Future2[M, N]) (T, V, M, N, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, err2 := f2.Wait()
	return v1, v2, v3, v4, errors.Join(err1, err2)
}

// -------------------- Combine 5 Values --------------------

// Combine5 futures, wait and return joined values and errors.
func Combine5[T, V, M, N, O any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O]) (
	T, V, M, N, O, error) {
	v1, v2, v3, v4, cErr := Combine4(f1, f2, f3, f4)
	v5, err := f5.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, err
}

// Combine1u4 combine Future with Future4, wait and return joined values and errors.
func Combine1u4[T, V, M, N, O any](f1 *Future[T], f2 *Future4[V, M, N, O]) (T, V, M, N, O, error) {
	v1, err1 := f1.Wait()
	v2, v3, v4, v5, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Combine1u2x2 combine Future with 2 Future2(s), wait and return joined values and errors.
func Combine1u2x2[T, V, M, N, O any](f1 *Future[T], f2 *Future2[V, M], f3 *Future2[N, O]) (T, V, M, N, O, error) {
	v1, err1 := f1.Wait()
	v2, v3, v4, v5, err2 := Combine2x2(f2, f3)
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Combine2x1u3 combine 2 Future(s) with Future3, wait and return joined values and errors.
func Combine2x1u3[T, V, M, N, O any](f1 *Future[T], f2 *Future[V], f3 *Future3[M, N, O]) (T, V, M, N, O, error) {
	v1, v2, err1 := Combine(f1, f2)
	v3, v4, v5, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Combine3x1u2 combine 2 Future(s) with Future3, wait and return joined values and errors.
func Combine3x1u2[T, V, M, N, O any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future2[N, O]) (T, V, M, N, O, error) {
	v1, v2, v3, err1 := Combine3(f1, f2, f3)
	v4, v5, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// Combine2u3 combine Future2 with Future3, wait and return joined values and errors.
func Combine2u3[T, V, M, N, O any](f1 *Future2[T, V], f2 *Future3[M, N, O]) (T, V, M, N, O, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, v5, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, errors.Join(err1, err2)
}

// -------------------- Combine 6 Values --------------------

// Combine6 futures, wait and return joined values and errors.
func Combine6[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, v5, cErr := Combine5(f1, f2, f3, f4, f5)
	v6, err := f6.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, err
}

// Combine1u5 combine Future with Future5, wait and return joined values and errors.
func Combine1u5[T, V, M, N, O, P any](f1 *Future[T], f2 *Future5[V, M, N, O, P]) (T, V, M, N, O, P, error) {
	v1, err1 := f1.Wait()
	v2, v3, v4, v5, v6, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine1u2u3 combine Future, Future2 with Future3, wait and return joined values and errors.
func Combine1u2u3[T, V, M, N, O, P any](f1 *Future[T], f2 *Future2[V, M], f3 *Future3[N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, err1 := Combine1u2(f1, f2)
	v4, v5, v6, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine2x1u4 combine 2 Future(s) with Future4, wait and return joined values and errors.
func Combine2x1u4[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future4[M, N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, err1 := Combine(f1, f2)
	v3, v4, v5, v6, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine2x1u2x2 combine 2 Future(s) with 2 Future2(s), wait and return joined values and errors.
func Combine2x1u2x2[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 *Future2[O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, err1 := Combine2x1u2(f1, f2, f3)
	v5, v6, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine3x1u3 combine 3 Future(s) with Future3, wait and return joined values and errors.
func Combine3x1u3[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future3[N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, err1 := Combine3(f1, f2, f3)
	v4, v5, v6, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine4x1u2 combine 4 Future(s) with Future2, wait and return joined values and errors.
func Combine4x1u2[T, V, M, N, O, P any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future2[O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, err1 := Combine4(f1, f2, f3, f4)
	v5, v6, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine2u4 combine Future2 with Future4, wait and return joined values and errors.
func Combine2u4[T, V, M, N, O, P any](f1 *Future2[T, V], f2 *Future4[M, N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, v5, v6, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine3x2 combine 3 Future2(s), wait and return joined values and errors.
func Combine3x2[T, V, M, N, O, P any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future2[O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, v4, err1 := Combine2x2(f1, f2)
	v5, v6, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// Combine2x3 combine 2 Future3(s), wait and return joined values and errors.
func Combine2x3[T, V, M, N, O, P any](f1 *Future3[T, V, M], f2 *Future3[N, O, P]) (T, V, M, N, O, P, error) {
	v1, v2, v3, err1 := f1.Wait()
	v4, v5, v6, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, errors.Join(err1, err2)
}

// -------------------- Combine 7 Values --------------------

// Combine7 Future(s), wait and return joined values and errors.
func Combine7[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, v6, cErr := Combine6(f1, f2, f3, f4, f5, f6)
	v7, err := f7.Wait()
	err = errors.Join(cErr, err)
	return v1, v2, v3, v4, v5, v6, v7, err
}

// Combine1u2u4 combine Future, Future2 with Future4, wait and return joined values and errors.
func Combine1u2u4[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future2[V, M], f3 *Future4[N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, err1 := Combine1u2(f1, f2)
	v4, v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine1u2x3 combine Future with 2 Future3(s), wait and return joined values and errors.
func Combine1u2x3[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future3[V, M, N], f3 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Combine1u3(f1, f2)
	v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine1u3x2 combine Future with 3 Future2(s), wait and return joined values and errors.
func Combine1u3x2[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future2[V, M], f3 *Future2[N, O], f4 Future2[P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, err1 := Combine1u2x2(f1, f2, f3)
	v6, v7, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine2x1u5 combine Future with Future5, wait and return joined values and errors.
func Combine2x1u5[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future5[M, N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, err1 := Combine(f1, f2)
	v3, v4, v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine2x1u2u3 combine 2 Future(s), Future2 with Future3, wait and return joined values and errors.
func Combine2x1u2u3[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Combine2x1u2(f1, f2, f3)
	v5, v6, v7, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine3x1u4 combine 3 Future(s) with Future4, wait and return joined values and errors.
func Combine3x1u4[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future4[N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, err1 := Combine3(f1, f2, f3)
	v4, v5, v6, v7, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine3x1u2x2 combine 3 Future(s) with Future4, wait and return joined values and errors.
func Combine3x1u2x2[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future2[N, O], f5 *Future2[P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, err1 := Combine3x1u2(f1, f2, f3, f4)
	v6, v7, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine4x1u3 combine 4 Future(s) with Future3, wait and return joined values and errors.
func Combine4x1u3[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Combine4(f1, f2, f3, f4)
	v5, v6, v7, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine5x1u2 combine 5 Future(s) with Future2, wait and return joined values and errors.
func Combine5x1u2[T, V, M, N, O, P, Q any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O], f6 *Future2[P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, v5, err1 := Combine5(f1, f2, f3, f4, f5)
	v6, v7, err2 := f6.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine2u5 combine Future2 with Future5, wait and return joined values and errors.
func Combine2u5[T, V, M, N, O, P, Q any](f1 *Future2[T, V], f2 *Future5[M, N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, err1 := f1.Wait()
	v3, v4, v5, v6, v7, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine2x2u3 combine 2 Future2 with Future3, wait and return joined values and errors.
func Combine2x2u3[T, V, M, N, O, P, Q any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future3[O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, v4, err1 := Combine2x2(f1, f2)
	v5, v6, v7, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// Combine3u4 combine Future3 with Future4, wait and return joined values and errors.
func Combine3u4[T, V, M, N, O, P, Q any](f1 *Future3[T, V, M], f2 *Future4[N, O, P, Q]) (T, V, M, N, O, P, Q, error) {
	v1, v2, v3, err1 := f1.Wait()
	v4, v5, v6, v7, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, errors.Join(err1, err2)
}

// -------------------- Combine 8 Values --------------------

// Combine8 futures, wait and return joined values and errors.
func Combine8[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O],
	f6 *Future[P], f7 *Future[Q], f8 *Future[R]) (T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, v7, err1 := Combine7(f1, f2, f3, f4, f5, f6, f7)
	v8, err2 := f8.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine2x1u2u4 combine 2 Future(s) with Future2 and Future4, wait and return joined values and errors.
func Combine2x1u2u4[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 Future4[O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, err1 := Combine2x1u2(f1, f2, f3)
	v5, v6, v7, v8, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine2x1u3x2 combine 2 Future(s) with 3 Future2, wait and return joined values and errors.
func Combine2x1u3x2[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future2[M, N], f4 *Future2[O, P], f5 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Combine2x1u2x2(f1, f2, f3, f4)
	v7, v8, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine2x1u2x3 combine 2 Future(s) with 2 Future3, wait and return joined values and errors.
func Combine2x1u2x3[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future3[M, N, O], f4 *Future3[P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, err1 := Combine2x1u3(f1, f2, f3)
	v6, v7, v8, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine3x1u2u3 combine 3 Future(s) with Future2 and Future3, wait and return joined values and errors.
func Combine3x1u2u3[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future2[N, O], f5 *Future3[P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, err1 := Combine3x1u2(f1, f2, f3, f4)
	v6, v7, v8, err2 := f5.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine4x1u2x2 combine 4 Future(s) with 2 Future2(s), wait and return joined values and errors.
func Combine4x1u2x2[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future2[O, P], f6 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Combine4x1u2(f1, f2, f3, f4, f5)
	v7, v8, err2 := f6.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine5x1u3 combine 5 Future(s) with Future3, wait and return joined values and errors.
func Combine5x1u3[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O], f6 *Future3[P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, err1 := Combine5(f1, f2, f3, f4, f5)
	v6, v7, v8, err2 := f6.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine6x1u2 combine 6 Future(s) with Future2, wait and return joined values and errors.
func Combine6x1u2[T, V, M, N, O, P, Q, R any](f1 *Future[T], f2 *Future[V], f3 *Future[M], f4 *Future[N], f5 *Future[O], f6 *Future[P], f7 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Combine6(f1, f2, f3, f4, f5, f6)
	v7, v8, err2 := f7.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine2x2u4 combine 2 Future2(s) with Future4, wait and return joined values and errors.
func Combine2x2u4[T, V, M, N, O, P, Q, R any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future4[O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, err1 := Combine2x2(f1, f2)
	v5, v6, v7, v8, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine4x2 combine 4 Future2(s), wait and return joined values and errors.
func Combine4x2[T, V, M, N, O, P, Q, R any](f1 *Future2[T, V], f2 *Future2[M, N], f3 *Future2[O, P], f4 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Combine3x2(f1, f2, f3)
	v7, v8, err2 := f4.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine3u5 combine Future3 with Future5, wait and return joined values and errors.
func Combine3u5[T, V, M, N, O, P, Q, R any](f1 *Future3[T, V, M], f2 *Future5[N, O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, err1 := f1.Wait()
	v4, v5, v6, v7, v8, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine2x3u2 combine 2 Future3(s) with Future2, wait and return joined values and errors.
func Combine2x3u2[T, V, M, N, O, P, Q, R any](f1 *Future3[T, V, M], f2 *Future3[N, O, P], f3 *Future2[Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, v5, v6, err1 := Combine2x3(f1, f2)
	v7, v8, err2 := f3.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// Combine2x4 combine 2 Future4(s), wait and return joined values and errors.
func Combine2x4[T, V, M, N, O, P, Q, R any](f1 *Future4[T, V, M, N], f2 *Future4[O, P, Q, R]) (
	T, V, M, N, O, P, Q, R, error) {
	v1, v2, v3, v4, err1 := f1.Wait()
	v5, v6, v7, v8, err2 := f2.Wait()
	return v1, v2, v3, v4, v5, v6, v7, v8, errors.Join(err1, err2)
}

// -------------------- Combine more than 8 Values --------------------

// CombineSame futures which return same type, wait and return all joined values and errors.
func CombineSame[T any](fs ...*Future[T]) ([]T, error) {
	var err error
	var ret []T
	for _, f := range fs {
		v, curErr := f.Wait()
		err = errors.Join(curErr)
		ret = append(ret, v)
	}
	return ret, err
}

// CombineSameTimeout combine futures which return same type, wait for timeout duration and return all values with joined errors,
// otherwise an ErrTimeout returned
func CombineSameTimeout[T any](timeout time.Duration, fs ...*Future[T]) ([]T, error) {
	return Timeout(timeout, func() ([]T, error) {
		return CombineSame(fs...)
	})
}

type futureI interface {
	waitUnify() ([]any, error)
	waitTimeout(time.Duration) ([]any, error)
}

// CombineAll futures, wait and return all values and joined errors.
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

// CombineAllTimeout combine futures, wait for timeout duration and return all values and joined errors,
// otherwise an ErrTimeout returned
func CombineAllTimeout(timeout time.Duration, fs ...futureI) ([]any, error) {
	return Timeout(timeout, func() ([]any, error) {
		return CombineAll(fs...)
	})
}

// CombineAny futures and return as soon as any one succeeds.
func CombineAny[T any](fs ...*Future[T]) (T, error) {
	for _, f := range fs {
		ret, err := f.Wait()
		if err == nil {
			return ret, nil
		}
	}
	var t T
	return t, ErrAllFailed
}

// CombineAnyTimeout wait for timeout duration and return as soon as any one of Futures succeeds
// otherwise an ErrTimeout returned.
func CombineAnyTimeout[T any](timeout time.Duration, fs ...*Future[T]) (T, error) {
	return Timeout(timeout, func() (T, error) {
		return CombineAny(fs...)
	})
}
