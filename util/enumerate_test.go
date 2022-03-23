package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/yah01/container"
)

func TestEnumerate(t *testing.T) {
	slice := NewSlice(5, 4, 3, 2, 1)
	indexIter := Enumerate(slice.Iter())

	idx := 0
	for ; indexIter.Valid(); indexIter.Next() {
		v := indexIter.Value()
		assert.Equal(t, idx, v.Idx)
		assert.Equal(t, 5-idx, v.Value)
		idx++
	}
}
