package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_absIndex(t *testing.T) {
	//             -4,-3,-2,-1
	//              0, 1, 2, 3
	assert.Equal(t, 3, absIndex(4, -1))
	assert.Equal(t, 2, absIndex(4, -2))
	assert.Equal(t, 1, absIndex(4, -3))
	assert.Equal(t, 0, absIndex(4, -4))

	assert.Equal(t, 0, absIndex(4, 0))
	assert.Equal(t, 1, absIndex(4, 1))
	assert.Equal(t, 2, absIndex(4, 2))
	assert.Equal(t, 3, absIndex(4, 3))

	// out of bounds
	assert.Equal(t, -1, absIndex(4, 4))
	assert.Equal(t, -1, absIndex(4, -5))
}

func TestSlice_absIndices(t *testing.T) {
	len := 4
	// -4,-3,-2,-1
	//  0, 1, 2, 3

	// no indicies given
	{
		i, j, err := absIndices(len)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// one index given
	{
		i, j, err := absIndices(len, 1)
		assert.Equal(t, 0, i)
		assert.Equal(t, -1, j)
		assert.Equal(t, "only one index given", err.Error())
	}

	// end
	{
		i, j, err := absIndices(len, -3, -1)
		assert.Equal(t, 1, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 1, 3)
		assert.Equal(t, 1, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// middle
	{
		i, j, err := absIndices(len, 1, 2)
		assert.Equal(t, 1, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -3, -2)
		assert.Equal(t, 1, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)
	}

	// begining
	{
		i, j, err := absIndices(len, 0, 2)
		assert.Equal(t, 0, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -4, -2)
		assert.Equal(t, 0, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)
	}

	// move within bounds
	{
		i, j, err := absIndices(len, -5, 5)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 0, 5)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -5, -1)
		assert.Equal(t, 0, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)
	}

	// mutually exclusive
	{
		i, j, err := absIndices(len, -1, -3)
		assert.Equal(t, 3, i)
		assert.Equal(t, 1, j)
		assert.NotNil(t, err)

		i, j, err = absIndices(len, 3, 1)
		assert.Equal(t, 3, i)
		assert.Equal(t, 1, j)
		assert.NotNil(t, err)
	}

	// single
	{
		i, j, err := absIndices(len, 0, 0)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 1, 1)
		assert.Equal(t, 1, i)
		assert.Equal(t, 2, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, 3, 3)
		assert.Equal(t, 3, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -1, -1)
		assert.Equal(t, 3, i)
		assert.Equal(t, 4, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -2, -2)
		assert.Equal(t, 2, i)
		assert.Equal(t, 3, j)
		assert.Nil(t, err)

		i, j, err = absIndices(len, -4, -4)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
		assert.Nil(t, err)
	}
}
