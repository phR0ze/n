package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not a test
// Using this to experiment with
func TestTest(t *testing.T) {
	assert.Equal(t, []int(nil), foo())
}

func foo() (result []int) {
	//result = append(result, 2)
	return
}
