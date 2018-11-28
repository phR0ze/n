package nub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIQ(t *testing.T) {
	q := Q(5)
	assert.True(t, q.Any())
	assert.Equal(t, 1, q.Len())
	q2 := q.Append(2)
	assert.True(t, q2.Any())
	assert.Equal(t, q, q2)
	assert.Equal(t, 2, q.Len())
	assert.Equal(t, 2, q2.Len())
	assert.Equal(t, 5, q.At(0).Int())
	assert.Equal(t, 2, q.At(1).Int())
}

func TestInt(t *testing.T) {
	assert.Equal(t, 1, Q(1).Int())
}

func TestInts(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Q([]int{1, 2, 3}).Ints())
}
