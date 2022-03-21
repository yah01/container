package util

import (
	. "github.com/yah01/container/iter"
)

type MapFunc[F any, T any] func(elem F) T

type Connector[F any, T any] struct {
	mapFunc MapFunc[F, T]
	src Iterator[F]
}

func (iter *Connector[F, T]) Valid() bool {
	return iter.src.Valid()
}

func (iter *Connector[F, T]) Value() T {
	return iter.mapFunc(iter.src.Value())
}

func (iter *Connector[F, T]) Next() {
	iter.src.Next()
}

func Map[F any, T any](iter Iterator[F], mapFunc MapFunc[F, T]) Iterator[T] {
	return &Connector[F, T]{
		mapFunc: mapFunc,
		src: iter,
	}
}
