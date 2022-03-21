package container

import . "github.com/yah01/container/iter"

type Container[T any] interface {
	Push(elem T)
	Iter() Iterator[T]
}

type PairContainer[K any, V any] interface {
}

func Collect[C Container[T], T any](iter Iterator[T], container C) Container[T] {
	for ; iter.Valid(); iter.Next() {
		container.Push(iter.Value())
	}

	return container
}
