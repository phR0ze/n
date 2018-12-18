package nos

import (
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	{
		result, err := Abs("")
		assert.NotNil(t, err)
		assert.Empty(t, result)
	}
	{
		result, err := Abs("~/")
		assert.Nil(t, err)
		assert.True(t, strings.Contains(result, "home"))
	}
	{
		result, err := Abs("test")
		assert.Nil(t, err)
		assert.True(t, strings.Contains(result, "home"))
		assert.True(t, strings.HasSuffix(result, "nos/test"))
	}
	{
		result, err := Abs("file://../utils")
		assert.Nil(t, err)
		fmt.Println(result)
		assert.True(t, strings.HasSuffix(result, "n/pkg/utils"))
	}
	{
		result, err := Abs("http://../utils")
		assert.Nil(t, err)
		fmt.Println(result)
		assert.True(t, strings.HasSuffix(result, "n/pkg/utils"))
	}
	{
		result, err := Abs("file:///utils")
		assert.Nil(t, err)
		fmt.Println(result)
		assert.Equal(t, "/utils", result)
	}
}

func TestDirs(t *testing.T) {
	{
		assert.Len(t, Dirs(""), 0)
	}
	{
		dirs := Dirs("../")
		assert.NotEmpty(t, dirs)
		for _, dir := range dirs {
			assert.True(t, strings.Contains(dir, "n/pkg"))
		}
	}
}

func TestSlicePath(t *testing.T) {
	assert.Equal(t, "", SlicePath("", 0, -1))
	assert.Equal(t, "/", SlicePath("/", 0, -1))
	assert.Equal(t, "/foo", SlicePath("/foo", 0, -1))

	// Slice first count
	assert.Equal(t, "", SlicePath("", 0, 1))
	assert.Equal(t, "/", SlicePath("/", 0, 1))
	assert.Equal(t, "foo", SlicePath("foo", 0, 1))
	assert.Equal(t, "/foo", SlicePath("/foo", 0, 1))
	assert.Equal(t, "/foo/bar", SlicePath("/foo/bar/one", 0, 1))

	assert.Equal(t, "/foo", SlicePath("/foo/bar/one", 0, 0))
	assert.Equal(t, "/foo/bar", SlicePath("/foo/bar/one", 0, 1))
	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one", 0, 2))
	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one", 0, 3))
	assert.Equal(t, "foo/bar/one", SlicePath("foo/bar/one", 0, 3))

	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one", 0, -1))
	assert.Equal(t, "/foo/bar", SlicePath("/foo/bar/one", 0, -2))
	assert.Equal(t, "/foo", SlicePath("/foo/bar/one", 0, -3))
	assert.Equal(t, "", SlicePath("/foo/bar/one", 0, -4))

	// Slice last cnt
	assert.Equal(t, "", SlicePath("", -2, -1))
	assert.Equal(t, "/", SlicePath("/", -2, -1))
	assert.Equal(t, "foo", SlicePath("foo", -2, -1))
	assert.Equal(t, "/foo", SlicePath("/foo", -2, -1))
	assert.Equal(t, "one", SlicePath("/foo/bar/one", -1, -1))
	assert.Equal(t, "one", SlicePath("foo/bar/one", -1, -1))
	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one", -3, -1))
	assert.Equal(t, "bar/one", SlicePath("/foo/bar/one", -2, -1))
	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one", -3, -1))
	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one", -5, 2))
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
		expected := []string{}
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
		expected := []string{}
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
