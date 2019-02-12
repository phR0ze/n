package sys

import (
	"fmt"
	"os"
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
		assert.True(t, pathContainsHome(result))
	}
	{
		result, err := Abs("test")
		assert.Nil(t, err)
		assert.True(t, pathContainsHome(result))
		assert.True(t, strings.HasSuffix(result, "sys/test"))
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

func TestFiles(t *testing.T) {
	{
		assert.Len(t, Files(""), 0)
	}
	{
		files := Files(".")
		assert.NotEmpty(t, files)
		for _, file := range files {
			assert.True(t, strings.Contains(file, "n/pkg/sys/"))
		}
	}
}

func TestPaths(t *testing.T) {
	{
		assert.Len(t, Paths(""), 0)
	}
	{
		paths := Paths("../../")
		assert.NotEmpty(t, paths)

		// Find at least one dir
		dirFound := false
		for _, path := range paths {
			if strings.HasSuffix(path, "n/pkg") {
				dirFound = true
				break
			}
		}
		assert.True(t, dirFound)

		// Find at least one file
		fileFound := false
		for _, path := range paths {
			if strings.HasSuffix(path, "README.md") {
				fileFound = true
				break
			}
		}
		assert.True(t, fileFound)
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
	assert.True(t, pathContainsHome(result))
}

func TestAllPaths(t *testing.T) {
	cleanTmpDir()
	{
		targetDir := path.Join(tmpDir, "first")
		targetDir, err := Abs(targetDir)
		assert.Nil(t, err)
		expected := []string{}
		MkdirP(targetDir)
		for i := 0; i < 10; i++ {
			target := path.Join(targetDir, fmt.Sprintf("%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		expected = append([]string{targetDir}, expected...)
		paths, err := AllPaths(targetDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}
	{
		targetDir := path.Join(tmpDir, "second")
		targetDir, err := Abs(targetDir)
		assert.Nil(t, err)
		expected := []string{}
		MkdirP(targetDir)
		for i := 0; i < 5; i++ {
			target := path.Join(targetDir, fmt.Sprintf("%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		expected = append([]string{targetDir}, expected...)
		paths, err := AllPaths(targetDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}
	{
		// Now create a link to a another directory in second
		secondDir, _ := Abs(path.Join(tmpDir, "second"))

		// Create third dir files
		thirdDir := path.Join(tmpDir, "third")
		thirdDir, err := Abs(thirdDir)
		assert.Nil(t, err)
		MkdirP(thirdDir)
		expected := []string{secondDir}
		for i := 0; i < 5; i++ {
			target := path.Join(secondDir, fmt.Sprintf("%d", i))
			expected = append(expected, target)
		}
		expected = append(expected, thirdDir)
		for i := 0; i < 5; i++ {
			target := path.Join(thirdDir, fmt.Sprintf("third-%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		expected = append(expected, path.Join(secondDir, "third"))

		// Create sysmlink in second dir to third dir
		symlink := path.Join(tmpDir, "second", "third")
		os.Symlink(thirdDir, symlink)

		// Check the result
		paths, err := AllPaths(secondDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
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

func TestTrimExt(t *testing.T) {
	assert.Equal(t, "", TrimExt(""))
	assert.Equal(t, "/foo/bar", TrimExt("/foo/bar"))
	assert.Equal(t, "/foo/bar", TrimExt("/foo/bar.mkv"))
	assert.Equal(t, "/foo/bar.mkv", TrimExt("/foo/bar.mkv.mp4"))
}

func TestTrimProtocol(t *testing.T) {
	assert.Equal(t, "/foo/bar", TrimProtocol("file:///foo/bar"))
	assert.Equal(t, "foo/bar", TrimProtocol("ftp://foo/bar"))
	assert.Equal(t, "foo/bar", TrimProtocol("http://foo/bar"))
	assert.Equal(t, "foo/bar", TrimProtocol("https://foo/bar"))
	assert.Equal(t, "foo://foo/bar", TrimProtocol("foo://foo/bar"))
}

func pathContainsHome(path string) (result bool) {
	for _, x := range []string{"home", "Users"} {
		if x == strings.Split(path, "/")[1] {
			result = true
			return
		}
	}
	return
}
