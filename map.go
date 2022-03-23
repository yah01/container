package container

type KVS[K comparable, V any] struct {
	key   K
	value V
}

var (
	_ Container[KVS[string, int]] = &Map[string, int]{}
)

type Map[K comparable, V any] struct {
	inner map[K]V
}

func (cmap *Map[K, V]) Push(kv KVS[K, V]) {
	cmap.inner[kv.key] = kv.value
}

func (cmap *Map[K, V]) Iter() Iterator[KVS[K, V]] {
	return &SliceIterator[KVS[K, V]]{}
}
