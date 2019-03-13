package sys

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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

func TestBase(t *testing.T) {
	assert.Equal(t, ".", filepath.Base(""))
	assert.Equal(t, "", Base(""))
	assert.Equal(t, "/", Base("/"))
	assert.Equal(t, "foo", Base("/foo"))
	assert.Equal(t, "foo", Base("/foo/"))
}

func TestDir(t *testing.T) {
	assert.Equal(t, ".", filepath.Dir(""))
	assert.Equal(t, "", Dir(""))
	assert.Equal(t, "/", Dir("/"))
	assert.Equal(t, "/", Dir("/foo"))
	assert.Equal(t, "/", Dir("/foo/"))
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

	// Handle slash at end
	assert.Equal(t, "/foo", SlicePath("/foo/bar/one/", 0, 0))
	assert.Equal(t, "foo/bar", SlicePath("foo/bar/one/", 0, 1))
	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one/", 0, 2))
	assert.Equal(t, "/foo/bar/one", SlicePath("/foo/bar/one/", 0, 3))
	assert.Equal(t, "foo/bar/one", SlicePath("foo/bar/one/", 0, 3))

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

	// Handle slash at end
	assert.Equal(t, "foo", SlicePath("foo/bar/", 0, -2))
	assert.Equal(t, "foo/bar", SlicePath("foo/bar/one/", 0, -2))

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

func TestAllFilesWithFileLink(t *testing.T) {
	cleanTmpDir()

	// temp/first/f0,f1
	firstDir, _ := MkdirP(path.Join(tmpDir, "first"))
	Touch(path.Join(firstDir, "f0"))
	Touch(path.Join(firstDir, "f1"))

	// temp/first/second/s0,s1
	secondDir, _ := MkdirP(path.Join(firstDir, "second"))
	Touch(path.Join(secondDir, "s0"))
	Touch(path.Join(secondDir, "s1"))

	// temp/first/third/t0,t1
	thirdDir, _ := MkdirP(path.Join(firstDir, "third"))
	Touch(path.Join(thirdDir, "t0"))
	Touch(path.Join(thirdDir, "t1"))

	// temp/first/second/t0 => temp/first/third/t0
	symlink := path.Join(secondDir, "t0")
	os.Symlink(path.Join(thirdDir, "t0"), symlink)

	// Test results using base names only
	results, _ := AllFiles(firstDir)
	for i := 0; i < len(results); i++ {
		results[i] = SlicePath(results[i], -2, -1)
	}
	//assert.Equal(t, []string{"first/f0", "first/f1", "second/s0", "second/s1", "second/t0", "third/t0", "third/t1"}, results)
}

func TestAllFiles(t *testing.T) {
	cleanTmpDir()

	// Single dir of files - No links
	// temp/first/f0,f1
	{
		targetDir, _ := MkdirP(path.Join(tmpDir, "first"))
		expected := []string{}
		for i := 0; i < 2; i++ {
			target := path.Join(targetDir, fmt.Sprintf("f%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		paths, err := AllFiles(targetDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}

	// Single dir of files - No links
	// temp/second/s0,s1
	{
		targetDir, _ := MkdirP(path.Join(tmpDir, "second"))
		expected := []string{}
		for i := 0; i < 2; i++ {
			target := path.Join(targetDir, fmt.Sprintf("s%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		paths, err := AllFiles(targetDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}

	// Now create a link to a another directory in second
	// temp/third/t0,t1
	// temp/second/s0,s1,third => temp/third, t0
	{
		secondDir, _ := Abs(path.Join(tmpDir, "second"))

		// Create third dir files
		thirdDir, _ := MkdirP(path.Join(tmpDir, "third"))
		expected := []string{}
		for i := 0; i < 2; i++ {
			target := path.Join(secondDir, fmt.Sprintf("s%d", i))
			expected = append(expected, target)
		}
		for i := 0; i < 2; i++ {
			target := path.Join(thirdDir, fmt.Sprintf("t%d", i))
			Touch(target)
			expected = append(expected, target)
		}

		// Create sysmlink in second dir to third dir
		symlink := path.Join(tmpDir, "second", "third")
		os.Symlink(thirdDir, symlink)

		// Create sysmlink in second dir to third dir file third-0
		symlink = path.Join(tmpDir, "second", "t0")
		os.Symlink(path.Join(thirdDir, "t0"), symlink)

		// Compute results using AllFiles
		paths, err := AllFiles(secondDir, newFollowOpt(false))
		assert.Nil(t, err)
		for i := range paths {
			paths[i] = SlicePath(paths[i], -2, -1)
		}
		// as both second/t0 and second/third are links they should be excluded
		assert.Equal(t, []string{"second/s0", "second/s1"}, paths)

		// Compute results using filepath.Walk
		paths2, err := filepathWalkAllFiles(secondDir)
		for i := range paths2 {
			paths2[i] = SlicePath(paths2[i], -2, -1)
		}
		assert.Nil(t, err)
		// filepath.Walk doesn't exclude symlinks
		assert.Equal(t, []string{"second/s0", "second/s1", "second/t0", "second/third"}, paths2)

		// Now ensure that links are followed
		paths, err = AllFiles(secondDir)
		assert.Nil(t, err)
		for i := range paths {
			paths[i] = SlicePath(paths[i], -2, -1)
		}
		// second/t0 is a link that will resolve to third/t0 so when third/t0 is walked it is dropped
		assert.Equal(t, []string{"second/s0", "second/s1", "third/t0", "third/t1"}, paths)
	}
}

func TestAllPaths(t *testing.T) {
	cleanTmpDir()

	// Single dir no links
	// temp/first/f0,f1
	{
		targetDir, _ := MkdirP(path.Join(tmpDir, "first"))
		expected := []string{}
		for i := 0; i < 2; i++ {
			target := path.Join(targetDir, fmt.Sprintf("f%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		expected = append([]string{targetDir}, expected...)
		paths, err := AllPaths(targetDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}

	// Second dir no links
	// temp/second/s0,s1
	{
		targetDir, _ := MkdirP(path.Join(tmpDir, "second"))
		expected := []string{}
		for i := 0; i < 2; i++ {
			target := path.Join(targetDir, fmt.Sprintf("s%d", i))
			Touch(target)
			expected = append(expected, target)
		}
		expected = append([]string{targetDir}, expected...)
		paths, err := AllPaths(targetDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}

	// Now create a link to a another directory in second
	// temp/second/s0,s1,third => third/t0,t1
	{
		secondDir, _ := Abs(path.Join(tmpDir, "second"))

		// Create third dir files
		thirdDir, _ := MkdirP(path.Join(tmpDir, "third"))
		for i := 0; i < 2; i++ {
			target := path.Join(thirdDir, fmt.Sprintf("t%d", i))
			Touch(target)
		}

		// Test before links
		paths, err := AllPaths(secondDir)
		assert.Nil(t, err)
		for i := range paths {
			paths[i] = SlicePath(paths[i], -2, -1)
		}
		assert.Equal(t, []string{"temp/second", "second/s0", "second/s1"}, paths)

		// Create sysmlink in second dir to third dir
		symlink := path.Join(tmpDir, "second", "third")
		os.Symlink(thirdDir, symlink)
		symlink = path.Join(tmpDir, "second", "t0")
		os.Symlink(path.Join(thirdDir, "t0"), symlink)

		// Test after links
		paths, err = AllPaths(secondDir)
		assert.Nil(t, err)
		for i := range paths {
			paths[i] = SlicePath(paths[i], -2, -1)
		}
		// distinctness in enforced so t0 will only show up once which changes order a bit when following links
		assert.Equal(t, []string{"temp/second", "second/s0", "second/s1", "second/t0", "third/t0", "second/third", "temp/third", "third/t1"}, paths)

		// Don't follow links now
		paths, err = AllPaths(secondDir, newFollowOpt(false))
		assert.Nil(t, err)
		for i := range paths {
			paths[i] = SlicePath(paths[i], -2, -1)
		}
		assert.Equal(t, []string{"temp/second", "second/s0", "second/s1", "second/t0", "second/third"}, paths)
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

func filepathWalkAllFiles(root string) (result []string, err error) {
	result = []string{}
	if root, err = Abs(root); err != nil {
		return
	}
	err = filepath.Walk(root, func(p string, i os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if p != root && p != "." && p != ".." && !i.IsDir() {
			absPath, e := Abs(p)
			if e != nil {
				return e
			}
			result = append(result, absPath)
		}
		return nil
	})
	return
}

func filepathWalkAllPaths(root string) (result []string, err error) {
	if root, err = Abs(root); err != nil {
		return
	}
	result = []string{root}
	err = filepath.Walk(root, func(p string, i os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if p != root && p != "." && p != ".." {
			absPath, e := Abs(p)
			if e != nil {
				return e
			}
			result = append(result, absPath)
		}
		return nil
	})
	return
}
