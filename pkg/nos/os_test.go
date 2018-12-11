package nos

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

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

func TestCopy(t *testing.T) {
	cleanTmpDir()
	src := "../../pkg"
	dst := path.Join(tmpDir, "pkg")

	assert.False(t, Exists(dst))
	Copy(src, tmpDir)
	assert.True(t, Exists(dst))

	srcPaths := Paths(src)
	dstPaths := Paths(dst)
	assert.Equal(t, len(srcPaths), len(dstPaths))
	for i := range srcPaths {
		assert.Equal(t, srcPaths[i], strings.Replace(dstPaths[i], "/test/temp", "", -1))
	}
}

func TestCopyFile(t *testing.T) {
	{
		cleanTmpDir()
		foo := path.Join(tmpDir, "foo")

		assert.False(t, Exists(foo))
		CopyFile(readme, foo)
		assert.True(t, Exists(foo))

		srcMD5, _ := MD5(readme)
		dstMD5, _ := MD5(foo)
		assert.Equal(t, srcMD5, dstMD5)
	}
}

func TestExists(t *testing.T) {
	assert.False(t, Exists("bob"))
	assert.True(t, Exists(readme))
}

func TestIsDir(t *testing.T) {
	assert.False(t, IsDir(readme))
	assert.True(t, IsDir("../.."))
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile(readme))
	assert.False(t, IsFile("../.."))
}

func TestMD5(t *testing.T) {
	if Exists(tmpfile) {
		os.Remove(tmpfile)
	}
	f, _ := os.Create(tmpfile)
	defer f.Close()
	f.WriteString(`This is a test of the emergency broadcast system.`)

	expected := "067a8c38325b12159844261d16e5cb13"
	result, _ := MD5(tmpfile)
	assert.Equal(t, expected, result)
}

func TestMkdirP(t *testing.T) {
	if Exists(tmpDir) {
		os.RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
	assert.True(t, Exists(tmpDir))
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

func TestReadLines(t *testing.T) {
	lines, err := ReadLines(testfile)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(lines))
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

func TestTouch(t *testing.T) {
	cleanTmpDir()
	assert.False(t, Exists(tmpfile))
	assert.Nil(t, Touch(tmpfile))
	assert.True(t, Exists(tmpfile))
	assert.Nil(t, Touch(tmpfile))
}

func TestWriteLines(t *testing.T) {
	cleanTmpDir()
	lines, err := ReadLines(testfile)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(lines))
	err = WriteLines(tmpfile, lines)
	assert.Nil(t, err)
	{
		lines2, err := ReadLines(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, lines, lines2)
	}
}

func cleanTmpDir() {
	if Exists(tmpDir) {
		os.RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
}