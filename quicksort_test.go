package recursive_algorithms_in_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuicksort(t *testing.T) {
	assert.Equal(t, []int{3, 3, 3}, quicksort([]int{3, 3, 3}))
	assert.Equal(t, []int{1, 2, 3}, quicksort([]int{1, 2, 3}))
	assert.Equal(t, []int{2, 3, 5}, quicksort([]int{3, 2, 5}))
	assert.Equal(t, []int{2, 2, 3, 3, 5, 6, 45, 99, 532}, quicksort([]int{3, 2, 5, 99, 45, 2, 532, 3, 6}))
}
