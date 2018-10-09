package recursive_algorithms_in_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	assert.Equal(t, 0, fibo(0))
	assert.Equal(t, 1, fibo(1))
	assert.Equal(t, 1, fibo(2))
	assert.Equal(t, 2, fibo(3))
	assert.Equal(t, 8, fibo(6))
	assert.Equal(t, 13, fibo(7))
}
