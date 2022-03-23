package container

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Iterator[T any] interface {
	Valid() bool
	Value() T
	Next()
}
