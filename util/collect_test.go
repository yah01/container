package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yah01/container"
)

func TestCollect(t *testing.T) {
	slice := container.NewSlice(1, 2, 3, 4, 5)
	iter := Map(slice.Iter(),
		func(elem int) int {
			return 6 - elem
		})

	container := Collect(iter, &container.Slice[int]{})
	iter = container.Iter()

	origin := slice.Iter()
	for ; iter.Valid() && origin.Valid(); iter.Next() {
		assert.Equal(t, iter.Value(), 6-origin.Value())
		origin.Next()
	}
}
