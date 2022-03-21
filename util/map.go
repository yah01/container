package util

import (
	. "github.com/yah01/container/iter"
)

type MapFunc[F any, T any] func(elem F) T

type MapResultIter[F any, T any] struct {
	mapF    MapFunc[F, T]
	srcIter Iterator[F]
}

func (iter *MapResultIter[F, T]) Valid() bool {
	return iter.srcIter.Valid()
}

func (iter *MapResultIter[F, T]) Value() T {
	return iter.mapF(iter.srcIter.Value())
}

func (iter *MapResultIter[F, T]) Next() {
	iter.srcIter.Next()
}

func Map[F any, T any](iter Iterator[F], mapF MapFunc[F, T]) Iterator[T] {
	return &MapResultIter[F, T]{
		mapF:    mapF,
		srcIter: iter,
	}
}
