package container

type Container[T any] interface {
	Push(elem T)
	Iter() Iterator[T]
}
