package future

import "errors"

var ErrNotRun = errors.New("[Future] err:not run")

// GoIf run worker in a new goroutine when can() reports true.
func GoIf[T any](cond bool, w worker[T]) *Future[T] {
	if !cond {
		return &Future[T]{
			e: ErrNotRun,
		}
	}
	return Go(w)
}

// Go2If run worker in a new goroutine when can() reports true.
func Go2If[T, V any](cond bool, w func() (T, V, error)) *Future2[T, V] {
	if !cond {
		return &Future2[T, V]{
			f: &Future[any]{
				e:   ErrNotRun,
				val: T2[T, V]{},
			},
		}
	}
	return Go2(w)
}

// Go3If run worker in a new goroutine when can() reports true.
func Go3If[T, V, M any](cond bool, w func() (T, V, M, error)) *Future3[T, V, M] {
	if !cond {
		return &Future3[T, V, M]{
			f: &Future[any]{
				e:   ErrNotRun,
				val: T3[T, V, M]{},
			},
		}
	}
	return Go3(w)
}

// Go4If run worker in a new goroutine when can() reports true.
func Go4If[T, V, M, N any](cond bool, w func() (T, V, M, N, error)) *Future4[T, V, M, N] {
	if !cond {
		return &Future4[T, V, M, N]{
			f: &Future[any]{
				e:   ErrNotRun,
				val: T4[T, V, M, N]{},
			},
		}
	}
	return Go4(w)
}

// Go5If run worker in a new goroutine when can() reports true.
func Go5If[T, V, M, N, O any](cond bool, w func() (T, V, M, N, O, error)) *Future5[T, V, M, N, O] {
	if !cond {
		return &Future5[T, V, M, N, O]{
			f: &Future[any]{
				e:   ErrNotRun,
				val: T5[T, V, M, N, O]{},
			},
		}
	}
	return Go5(w)
}
