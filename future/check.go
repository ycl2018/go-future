package future

// Check invoke function when Future complete and return a new Future of the same type.
// this method support a chance to check error and rewrite result.
func (f *Future[T]) Check(w func(T, error) (T, error)) *Future[T] {
	return Go(func() (T, error) {
		val, err := f.Wait()
		return w(val, err)
	})
}

// Check invoke function when Future complete and return a new Future of the same type.
// this method support a chance to check error and rewrite result.
func (f *Future2[T, V]) Check(w func(T, V, error) (T, V, error)) *Future2[T, V] {
	return Go2(func() (T, V, error) {
		val1, val2, err := f.Wait()
		return w(val1, val2, err)
	})
}

// Check invoke function when Future complete and return a new Future of the same type.
// this method support a chance to check error and rewrite result.
func (f *Future3[T, V, M]) Check(w func(T, V, M, error) (T, V, M, error)) *Future3[T, V, M] {
	return Go3(func() (T, V, M, error) {
		val1, val2, val3, err := f.Wait()
		return w(val1, val2, val3, err)
	})
}

// Check invoke function when Future complete and return a new Future of the same type.
// this method support a chance to check error and rewrite result.
func (f *Future4[T, V, M, N]) Check(w func(T, V, M, N, error) (T, V, M, N, error)) *Future4[T, V, M, N] {
	return Go4(func() (T, V, M, N, error) {
		val1, val2, val3, val4, err := f.Wait()
		return w(val1, val2, val3, val4, err)
	})
}

// Check invoke function when Future complete and return a new Future of the same type.
// this method support a chance to check error and rewrite result.
func (f *Future5[T, V, M, N, O]) Check(w func(T, V, M, N, O, error) (T, V, M, N, O, error)) *Future5[T, V, M, N, O] {
	return Go5(func() (T, V, M, N, O, error) {
		val1, val2, val3, val4, val5, err := f.Wait()
		return w(val1, val2, val3, val4, val5, err)
	})
}
