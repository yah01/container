package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/yah01/container/cslice"
	"github.com/yah01/container/util"
)

func TestCollect(t *testing.T) {
	slice := NewSlice(1, 2, 3, 4, 5)
	iter := util.Map(slice.Iter(),
		func(elem int) int {
			return 6 - elem
		})

	container := Collect(iter, &Slice[int]{})
	iter = container.Iter()

	origin := slice.Iter()
	for ; iter.Valid() && origin.Valid(); iter.Next() {
		assert.Equal(t, iter.Value(), 6-origin.Value())
		origin.Next()
	}
}
