package util

import . "github.com/yah01/container"

type WithIndex[T any] struct {
	Idx   int
	Value T
}

func Enumerate[T any](iter Iterator[T]) Connector[T, WithIndex[T]] {
	idx := 0
	return Connector[T, WithIndex[T]]{
		mapFunc: func(elem T) WithIndex[T] {
			withIndex := WithIndex[T]{
				Idx:   idx,
				Value: elem,
			}
			idx++
			return withIndex
		},
		src: iter,
	}
}
