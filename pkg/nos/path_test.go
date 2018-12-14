package nos

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathAbs(t *testing.T) {
	{
		result, _ := Path("~/").Abs()
		assert.True(t, strings.Contains(result, "home"))
	}
	{
		result, _ := Path("test").Abs()
		assert.True(t, strings.Contains(result, "home"))
		assert.True(t, strings.HasSuffix(result, "nos/test"))
	}
}

func TestPathSlice(t *testing.T) {
	assert.Equal(t, "", Path("").Slice(0, -1))
	assert.Equal(t, "/", Path("/").Slice(0, -1))
	assert.Equal(t, "/foo", Path("/foo").Slice(0, -1))

	// Slice first count
	assert.Equal(t, "", Path("").Slice(0, 1))
	assert.Equal(t, "/", Path("/").Slice(0, 1))
	assert.Equal(t, "foo", Path("foo").Slice(0, 1))
	assert.Equal(t, "/foo", Path("/foo").Slice(0, 1))
	assert.Equal(t, "/foo/bar", Path("/foo/bar/one").Slice(0, 1))

	assert.Equal(t, "/foo", Path("/foo/bar/one").Slice(0, 0))
	assert.Equal(t, "/foo/bar", Path("/foo/bar/one").Slice(0, 1))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Slice(0, 2))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Slice(0, 3))
	assert.Equal(t, "foo/bar/one", Path("foo/bar/one").Slice(0, 3))

	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Slice(0, -1))
	assert.Equal(t, "/foo/bar", Path("/foo/bar/one").Slice(0, -2))
	assert.Equal(t, "/foo", Path("/foo/bar/one").Slice(0, -3))
	assert.Equal(t, "", Path("/foo/bar/one").Slice(0, -4))

	// Slice last cnt
	assert.Equal(t, "", Path("").Slice(-2, -1))
	assert.Equal(t, "/", Path("/").Slice(-2, -1))
	assert.Equal(t, "foo", Path("foo").Slice(-2, -1))
	assert.Equal(t, "/foo", Path("/foo").Slice(-2, -1))
	assert.Equal(t, "one", Path("/foo/bar/one").Slice(-1, -1))
	assert.Equal(t, "one", Path("foo/bar/one").Slice(-1, -1))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Slice(-3, -1))
	assert.Equal(t, "bar/one", Path("/foo/bar/one").Slice(-2, -1))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Slice(-3, -1))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Slice(-5, 2))
}
