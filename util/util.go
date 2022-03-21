package util

import (
	. "github.com/yah01/container/iter"
	"golang.org/x/exp/constraints"
)

type ForFunc[T any] func(elem T)
type ReduceFunc[T constraints.Ordered] func(result T, elem T) T

func Sum[T constraints.Ordered](iter Iterator[T]) T {
	var sum T
	for ; iter.Valid(); iter.Next() {
		sum += iter.Value()
	}

	return sum
}

func Reduce[T constraints.Ordered](iter Iterator[T], reduceF ReduceFunc[T]) T {
	var ret T
	for ; iter.Valid(); iter.Next() {
		ret = reduceF(ret, iter.Value())
	}

	return ret
}

func For[T any](iter Iterator[T], forFunc ForFunc[T]) {
	for ; iter.Valid(); iter.Next() {
		forFunc(iter.Value())
	}
}