package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not a test
// Using this to experiment with
func TestTest(t *testing.T) {
	result, _ := go-homedir.Expand("~/")
	assert.Equal(t, "", result)
}
