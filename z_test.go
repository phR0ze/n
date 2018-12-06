package n

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not a test more of an experiment that I'll frequently change as I validate sanity
func TestTest(t *testing.T) {
	assert.True(t, strings.HasPrefix("spec:", "spec:"))
}
