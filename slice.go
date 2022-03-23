package container


type Slice[T any] struct {
	inner []T
}

func NewSlice[T any](elems ...T) Slice[T] {
	return Slice[T]{
		inner: elems,
	}
}

func (slice *Slice[T]) Push(elem T) {
	slice.inner = append(slice.inner, elem)
}

func (slice *Slice[T]) Pop() T {
	ret := slice.Last()
	slice.inner = slice.inner[:slice.Len()-1]

	return ret
}

func (slice *Slice[T]) First() T {
	return slice.inner[0]
}

func (slice *Slice[T]) Last() T {
	return slice.inner[slice.Len()-1]
}

func (slice *Slice[T]) Len() int {
	return len(slice.inner)
}

func (slice *Slice[T]) Cap() int {
	return cap(slice.inner)
}

func (slice *Slice[T]) Iter() Iterator[T] {
	return SliceIter(slice.inner)
}

func (slice *Slice[T]) Inner() []T {
	return slice.inner
}
