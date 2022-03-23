package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yah01/container"
)

func TestMapIter(t *testing.T) {
	slice := container.NewSlice(5, 4, 3, 2, 1)
	iter := slice.Iter()

	mapIter := Map(slice.Iter(), func(elem int) int {
		return 6 - elem
	})

	for ; mapIter.Valid() && iter.Valid(); mapIter.Next() {
		assert.Equal(t, mapIter.Value(), 6-iter.Value())
		iter.Next()
	}
}
