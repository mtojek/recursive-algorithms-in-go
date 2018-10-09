package recursive_algorithms_in_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, 1, gcd(4, 5))
	assert.Equal(t, 2, gcd(4,2))
	assert.Equal(t, 6, gcd(18,12))
	assert.Equal(t, 4, gcd(56,36))
}
