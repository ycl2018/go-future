//go:build 1.23

package future

import "iter"

func (g *Group[T]) All() (iter.Seq[T], error) {
	ret, err := g.Wait()
	if err != nil {
		return iter.Seq[T]{}, err
	}
	return func(yield func(T) bool) {
		for _, t := range ret {
			if !yield(t) {
				return
			}
		}
	}, nil
}
