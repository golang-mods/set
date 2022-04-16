package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedOrderedIntersection(t *testing.T) {
	set := SortedOrderedIntersection([]int{0, 1, 2, 3, 4, 6}, []int{1, 4, 5})
	assert.Equal(t, []int{1, 4}, set)
}

func TestSortedOrderedDifference(t *testing.T) {
	set := SortedOrderedDifference([]int{0, 1, 2, 3, 4, 6}, []int{1, 4, 5})
	assert.Equal(t, []int{0, 2, 3, 6}, set)
}
