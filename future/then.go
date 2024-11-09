package future

// Then compose other function which run after this Future complete without error, and return a new Future.
func (f *Future[T]) Then(w func(T) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		val, err := f.Wait()
		if err != nil {
			return nil, err
		}
		return w(val)
	})
}

// Then compose other function which run after this Future complete without error, and return a new Future.
func (f *Future2[T, V]) Then(w func(T, V) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		val1, val2, err := f.Wait()
		if err != nil {
			return nil, err
		}
		return w(val1, val2)
	})
}

// Then compose other function which run after this Future complete without error, and return a new Future.
func (f *Future3[T, V, M]) Then(w func(T, V, M) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		val1, val2, val3, err := f.Wait()
		if err != nil {
			return nil, err
		}
		return w(val1, val2, val3)
	})
}

// Then compose other function which run after this Future complete without error, and return a new Future.
func (f *Future4[T, V, M, N]) Then(w func(T, V, M, N) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		val1, val2, val3, val4, err := f.Wait()
		if err != nil {
			return nil, err
		}
		return w(val1, val2, val3, val4)
	})
}

// Then compose other function which run after this Future complete without error, and return a new Future.
func (f *Future5[T, V, M, N, O]) Then(w func(T, V, M, N, O) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		val1, val2, val3, val4, val5, err := f.Wait()
		if err != nil {
			return nil, err
		}
		return w(val1, val2, val3, val4, val5)
	})
}
