package container

import . "github.com/yah01/container/iter"

type Container[T any] interface {
	Push(elem T)
	Iter() Iterator[T]
}

func Collect[C Container[T], T any](iter Iterator[T], container C) Container[T] {
	for ; iter.Valid(); iter.Next() {
		container.Push(iter.Value())
	}

	return container
}
