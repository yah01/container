package container

var (
	_ Iterator[int] = &SliceIterator[int]{}
)

type SliceIterator[T any] struct {
	idx   int
	slice []T
}

func SliceIter[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		idx:   0,
		slice: slice,
	}
}

func (iter *SliceIterator[T]) Valid() bool {
	return iter.idx < len(iter.slice)
}

func (iter *SliceIterator[T]) Value() T {
	return iter.slice[iter.idx]
}

func (iter *SliceIterator[T]) Next() {
	iter.idx++
}
