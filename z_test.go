package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not a test more of an experiment that I'll frequently change as I validate sanity
func TestTest(t *testing.T) {
	result := A("2").Split(":").S()
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "2", A("2").Split(":").Last().A())
}
