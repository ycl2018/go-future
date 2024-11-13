package future

// Handle compose other function which run after this Future complete, and return a new Future.
func (f *Future[T]) Handle(w func(T, error) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		return w(f.Wait())
	})
}

// Handle compose other function which run after this Future complete, and return a new Future.
func (f *Future2[T, V]) Handle(w func(T, V, error) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		return w(f.Wait())
	})
}

// Handle compose other function which run after this Future complete, and return a new Future.
func (f *Future3[T, V, M]) Handle(w func(T, V, M, error) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		return w(f.Wait())
	})
}

// Handle compose other function which run after this Future complete, and return a new Future.
func (f *Future4[T, V, M, N]) Handle(w func(T, V, M, N, error) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		return w(f.Wait())
	})
}

// Handle compose other function which run after this Future complete, and return a new Future.
func (f *Future5[T, V, M, N, O]) Handle(w func(T, V, M, N, O, error) (any, error)) *Future[any] {
	return Go(func() (any, error) {
		return w(f.Wait())
	})
}
