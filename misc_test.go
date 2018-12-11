package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	assert.Equal(t, []int{0}, Range(0, 0))
	assert.Equal(t, []int{0, 1}, Range(0, 1))
	assert.Equal(t, []int{3, 4, 5, 6, 7, 8}, Range(3, 8))
}
