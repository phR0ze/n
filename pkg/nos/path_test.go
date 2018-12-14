package nos

import (
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathAbs(t *testing.T) {
	{
		result, err := Path("~/").Abs()
		assert.Nil(t, err)
		assert.True(t, strings.Contains(result, "home"))
	}
	{
		result, err := Path("test").Abs()
		assert.Nil(t, err)
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

func TestHome(t *testing.T) {
	result, err := Home()
	assert.Nil(t, err)
	assert.True(t, strings.Contains(result, "home"))
}

func TestPaths(t *testing.T) {
	cleanTmpDir()
	{
		targetDir := path.Join(tmpDir, "first")
		expected := []string{targetDir}
		MkdirP(targetDir)
		for i := 0; i < 10; i++ {
			target := path.Join(targetDir, fmt.Sprintf("%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		assert.Equal(t, expected, Paths(targetDir))
	}
	{
		targetDir := path.Join(tmpDir, "second")
		expected := []string{targetDir}
		MkdirP(targetDir)
		for i := 0; i < 5; i++ {
			target := path.Join(targetDir, fmt.Sprintf("%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		assert.Equal(t, expected, Paths(targetDir))
	}
}

func TestSharedDir(t *testing.T) {
	{
		first := ""
		second := ""
		assert.Equal(t, "", SharedDir(first, second))
	}
	{
		first := "/bob"
		second := "/foo"
		assert.Equal(t, "", SharedDir(first, second))
	}
	{
		first := "/foo/bar1"
		second := "/foo/bar2"
		assert.Equal(t, "/foo", SharedDir(first, second))
	}
	{
		first := "foo/bar1"
		second := "foo/bar2"
		assert.Equal(t, "foo", SharedDir(first, second))
	}
	{
		first := "/foo/bar/1"
		second := "/foo/bar/2"
		assert.Equal(t, "/foo/bar", SharedDir(first, second))
	}
}
