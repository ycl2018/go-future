package future

import "errors"

func unwrap(x []any) any {
	if len(x) == 1 {
		return x[0]
	}
	return x
}

// Join other Future task and return a new combined Future2.
// if the joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future[T]) Join(other futureI) *Future2[T, any] {
	return Go2(func() (T, any, error) {
		val1, err1 := f.Wait()
		val2, err2 := other.waitUnify()
		return val1, unwrap(val2), errors.Join(err1, err2)
	})
}

// Join other Future task and return a new combined Future3.
// if the joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future2[T, V]) Join(other futureI) *Future3[T, V, any] {
	return Go3(func() (T, V, any, error) {
		val1, val2, err1 := f.Wait()
		val3, err2 := other.waitUnify()
		return val1, val2, unwrap(val3), errors.Join(err1, err2)
	})
}

// Join other Future task and return a new combined Future4.
// if the joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future3[T, V, M]) Join(other futureI) *Future4[T, V, M, any] {
	return Go4(func() (T, V, M, any, error) {
		val1, val2, val3, err1 := f.Wait()
		val4, err2 := other.waitUnify()
		return val1, val2, val3, unwrap(val4), errors.Join(err1, err2)
	})
}

// Join other Future task and return a new combined Future5.
// if the joined Future return one value, combined Future's type 'any' will be the value exactly,
// or else it's real type is []any.
func (f *Future4[T, V, M, N]) Join(other futureI) *Future5[T, V, M, N, any] {
	return Go5(func() (T, V, M, N, any, error) {
		val1, val2, val3, val4, err1 := f.Wait()
		val5, err2 := other.waitUnify()
		return val1, val2, val3, val4, unwrap(val5), errors.Join(err1, err2)
	})
}

// Join other Future task and return a new combined Future.
// in the combined Future's '[6]any',5 values before are current Future's returned values,
// the last is the combined Future's returned values.
func (f *Future5[T, V, M, N, O]) Join(other futureI) *Future[[6]any] {
	return Go(func() ([6]any, error) {
		val1, val2, val3, val4, val5, err1 := f.Wait()
		val6, err2 := other.waitUnify()
		return [...]any{val1, val2, val3, val4, val5, unwrap(val6)}, errors.Join(err1, err2)
	})
}
