package util

import . "github.com/yah01/container/iter"

type WithIndex[T any] struct {
	Idx   int
	Value T
}

func Enumerate[T any](iter Iterator[T]) MapResultIter[T, WithIndex[T]] {
	idx := 0
	return MapResultIter[T, WithIndex[T]]{
		mapF: func(elem T) WithIndex[T] {
			withIndex := WithIndex[T]{
				Idx:   idx,
				Value: elem,
			}
			idx++
			return withIndex
		},
		srcIter: iter,
	}
}
