package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not a test more of an experiment that I'll frequently change as I validate sanity
func TestTest(t *testing.T) {
	assert.Equal(t, "foo", A("::foo").Split("::").Last().A())
}
