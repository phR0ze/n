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
var tmpFile = "../../test/temp/.tmp"
var readme = "../../README.md"

func TestFirstLast(t *testing.T) {
	assert.Equal(t, "", Path("").First())
	assert.Equal(t, "/", Path("/").First())
	assert.Equal(t, "foo", Path("foo").First())
	assert.Equal(t, "/foo", Path("/foo").First())
	assert.Equal(t, "", Path("/foo/bar/one").First(0))
	assert.Equal(t, "/foo", Path("/foo/bar/one").First())
	assert.Equal(t, "/foo", Path("/foo/bar/one").First(1))
	assert.Equal(t, "/foo/bar", Path("/foo/bar/one").First(2))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").First(3))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").First(5))
}

func TestPathLast(t *testing.T) {
	assert.Equal(t, "", Path("").Last())
	assert.Equal(t, "/", Path("/").Last())
	assert.Equal(t, "foo", Path("foo").Last())
	assert.Equal(t, "/foo", Path("/foo").Last())
	assert.Equal(t, "", Path("/foo/bar/one").Last(0))
	assert.Equal(t, "one", Path("/foo/bar/one").Last())
	assert.Equal(t, "one", Path("/foo/bar/one").Last(1))
	assert.Equal(t, "bar/one", Path("/foo/bar/one").Last(2))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Last(3))
	assert.Equal(t, "/foo/bar/one", Path("/foo/bar/one").Last(5))
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
	if Exists(tmpFile) {
		os.Remove(tmpFile)
	}
	f, _ := os.Create(tmpFile)
	defer f.Close()
	f.WriteString(`This is a test of the emergency broadcast system.`)

	expected := "067a8c38325b12159844261d16e5cb13"
	result, _ := MD5(tmpFile)
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
	assert.False(t, Exists(tmpFile))
	assert.Nil(t, Touch(tmpFile))
	assert.True(t, Exists(tmpFile))
	assert.Nil(t, Touch(tmpFile))
}

func cleanTmpDir() {
	if Exists(tmpDir) {
		os.RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
}
