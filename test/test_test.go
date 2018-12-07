package n

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Not a test
// Using this to experiment with
func TestTest(t *testing.T) {
	assert.Equal(t, "/foo/bar", path.Join("/foo", "/bar"))
}
