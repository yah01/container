package util

import . "github.com/yah01/container"

func Collect[C Container[T], T any](iter Iterator[T], container C) Container[T] {
	for ; iter.Valid(); iter.Next() {
		container.Push(iter.Value())
	}

	return container
}