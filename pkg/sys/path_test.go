package sys

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/phR0ze/n/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {

	// force filepathAbs error
	{
		test.OneShotForceFilePathAbsError()
		result, err := Abs("foo")
		assert.Empty(t, result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to compute the absolute path for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

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

	// HOME not set
	{
		// unset HOME
		home := os.Getenv("HOME")
		os.Unsetenv("HOME")
		assert.Equal(t, "", os.Getenv("HOME"))
		defer os.Setenv("HOME", home)

		result, err := Abs("~/")
		assert.Equal(t, "failed to expand the given path ~/: failed to compute the user's home directory: $HOME is not defined", err.Error())
		assert.Equal(t, "", result)
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

func TestGlob(t *testing.T) {
	resetTest()

	// Create test files in dir for globbing and valide modes
	dir, err := MkdirP(path.Join(tmpDir, "dir"))
	assert.Nil(t, err)
	file1, err := CopyFile(testfile, path.Join(dir, "file1"))
	assert.Nil(t, err)
	assert.Equal(t, os.FileMode(0644), Mode(file1))
	file2, err := CopyFile(testfile, path.Join(dir, "file2"))
	assert.Nil(t, err)
	assert.Equal(t, os.FileMode(0644), Mode(file2))
	bob1, err := CopyFile(testfile, path.Join(dir, "bob1"))
	assert.Nil(t, err)
	assert.Equal(t, os.FileMode(0644), Mode(bob1))

	// glob single file pattern
	{
		sources, err := Glob(path.Join(dir, "bob*"))
		assert.Nil(t, err)
		assert.Len(t, sources, 1)
		assert.Equal(t, bob1, sources[0])
	}

	// glob single file by full path
	{
		sources, err := Glob(testfile)
		assert.Nil(t, err)
		assert.Len(t, sources, 1)
		assert.Equal(t, path.Base(testfile), path.Base(sources[0]))
	}

	// glob pattern
	{
		sources, err := Glob(path.Join(dir, "*1"))
		assert.Nil(t, err)
		assert.Len(t, sources, 2)
		assert.Equal(t, bob1, sources[0])
		assert.Equal(t, file1, sources[1])
	}

	// recurse
	{
		sources, err := Glob(dir, RecurseOpt(true))
		assert.Nil(t, err)
		assert.Len(t, sources, 4)
		assert.Equal(t, dir, sources[0])
		assert.Equal(t, bob1, sources[1])
		assert.Equal(t, file1, sources[2])
		assert.Equal(t, file2, sources[3])
	}

	// force Glob error
	{
		test.OneShotForceFilePathGlobError()
		sources, err := Glob(testfile)
		assert.Len(t, sources, 0)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to get glob for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// empty path
	{
		sources, err := Glob("")
		assert.Len(t, sources, 0)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}
}

func TestPaths(t *testing.T) {

	// an invalid path should return nothing
	{
		assert.Len(t, Paths(""), 0)
	}

	// get files recursively
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

func TestAllFilesWithFileLink(t *testing.T) {
	resetTest()

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

func TestAllDirs(t *testing.T) {
	resetTest()

	// doesn't exist
	{
		paths, err := AllDirs(path.Join(tmpDir, "foo"))
		assert.Equal(t, ([]string)(nil), paths)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to execute Lstat against"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// empty dir name
	{
		paths, err := AllDirs("")
		assert.Equal(t, ([]string)(nil), paths)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// Single dir
	// temp/first
	{
		targetDir, _ := MkdirP(path.Join(tmpDir, "first"))
		expected := []string{targetDir}
		paths, err := AllDirs(tmpDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}

	// Second dir
	// temp/second
	{
		targetDir, _ := MkdirP(path.Join(tmpDir, "second"))
		expected := []string{path.Join(path.Dir(targetDir), "first"), targetDir}
		paths, err := AllDirs(tmpDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}

	// Nested dirs
	// temp/first/d1
	// temp/second/d2
	{
		firstDir, _ := Abs(path.Join(tmpDir, "first"))
		secondDir, _ := Abs(path.Join(tmpDir, "second"))
		d1, _ := MkdirP(path.Join(firstDir, "d1"))
		d2, _ := MkdirP(path.Join(secondDir, "d2"))
		d3, _ := MkdirP(path.Join(d2, "d3"))

		expected := []string{firstDir, d1, secondDir, d2, d3}
		paths, err := AllDirs(tmpDir)
		assert.Nil(t, err)
		assert.Equal(t, expected, paths)
	}
}

func TestAllFiles(t *testing.T) {
	resetTest()

	// doesn't exist
	{
		paths, err := AllFiles(path.Join(tmpDir, "foo"))
		assert.Equal(t, ([]string)(nil), paths)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to execute Lstat against"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// invalid dir
	{
		paths, err := AllFiles("")
		assert.Equal(t, ([]string)(nil), paths)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

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
		paths, err := AllFiles(secondDir, FollowOpt(false))
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
	resetTest()

	// doesn't exist
	{
		expected, err := Abs(path.Join(tmpDir, "foo"))
		assert.Nil(t, err)
		paths, err := AllPaths(path.Join(tmpDir, "foo"))
		assert.Equal(t, []string{expected}, paths)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to execute Lstat against"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// invalid dir
	{
		paths, err := AllPaths("")
		assert.Equal(t, ([]string)(nil), paths)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

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
		paths, err = AllPaths(secondDir, FollowOpt(false))
		assert.Nil(t, err)
		for i := range paths {
			paths[i] = SlicePath(paths[i], -2, -1)
		}
		assert.Equal(t, []string{"temp/second", "second/s0", "second/s1", "second/t0", "second/third"}, paths)
	}
}

func TestExpand(t *testing.T) {

	// More than one ~
	{
		result, err := Expand("~/foo~")
		assert.Equal(t, "invalid expansion requested", err.Error())
		assert.Equal(t, "", result)
	}

	// invalid path
	{
		result, err := Expand("~foo")
		assert.Equal(t, "failed to expand invalid path", err.Error())
		assert.Equal(t, "", result)
	}

	// happy
	{
		result, err := Expand("~/")
		assert.Nil(t, err)
		assert.True(t, pathContainsHome(result))
	}

	// HOME not set
	{
		// unset HOME
		home := os.Getenv("HOME")
		os.Unsetenv("HOME")
		assert.Equal(t, "", os.Getenv("HOME"))
		defer os.Setenv("HOME", home)

		result, err := Expand("~/")
		assert.Equal(t, "failed to compute the user's home directory: $HOME is not defined", err.Error())
		assert.Equal(t, "", result)
	}
}

func TestHome(t *testing.T) {

	// force filepathAbs error
	{
		test.OneShotForceFilePathAbsError()
		result, err := UserHome()
		assert.Empty(t, result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to compute the absolute path for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// happy
	{
		result, err := UserHome()
		assert.Nil(t, err)
		assert.True(t, pathContainsHome(result))
	}

	// HOME not set
	{
		// unset HOME
		home := os.Getenv("HOME")
		os.Unsetenv("HOME")
		assert.Equal(t, "", os.Getenv("HOME"))
		defer os.Setenv("HOME", home)

		home, err := UserHome()
		assert.Equal(t, "", home)
		assert.Equal(t, "failed to compute the user's home directory: $HOME is not defined", err.Error())
	}
}

func TestReadDir(t *testing.T) {
	resetTest()

	// read dirnames from valid directory
	{
		// crate a directory and test files to check
		dirname, err := MkdirP(path.Join(tmpDir, "dir"))
		assert.Nil(t, err)
		expected := []string{}
		for i := 0; i < 10; i++ {
			file, err := Touch(path.Join(dirname, fmt.Sprintf("file%d", i)))
			assert.Nil(t, err)
			expected = append(expected, SlicePath(file, -1, -1))
		}

		// Pull in the dir's filenames
		dirs, err := ReadDir(dirname)
		assert.Nil(t, err)
		results := []string{}
		for _, info := range dirs {
			results = append(results, info.Name())
		}
		assert.Equal(t, expected, results)

		// clean up
		assert.Nil(t, RemoveAll(dirname))
	}

	// try a file
	{
		dirs, err := ReadDir(testfile)
		assert.Equal(t, ([]*FileInfo)(nil), dirs)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to read directory"))
		assert.True(t, strings.Contains(err.Error(), ": readdirent"))
	}

	// writeonly directory
	{
		// Create writeonly directory
		dirname, err := MkdirP(path.Join(tmpDir, "dir"))
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dirname, 0222))

		dirs, err := ReadDir(dirname)
		assert.Equal(t, ([]*FileInfo)(nil), dirs)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open directory"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Reset the permissions and remove the dir
		assert.Nil(t, os.Chmod(dirname, 0755))
		assert.Nil(t, RemoveAll(dirname))
	}

	// empty dirname
	{
		dirs, err := ReadDir("")
		assert.Equal(t, ([]*FileInfo)(nil), dirs)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}
}

func TestReadDirnames(t *testing.T) {
	resetTest()

	// read dirnames from valid directory
	{
		// crate a directory and test files to check
		dirname, err := MkdirP(path.Join(tmpDir, "dir"))
		assert.Nil(t, err)
		expected := []string{}
		for i := 0; i < 10; i++ {
			file, err := Touch(path.Join(dirname, fmt.Sprintf("file%d", i)))
			assert.Nil(t, err)
			expected = append(expected, SlicePath(file, -1, -1))
		}

		// Pull in the dir's filenames
		dirs, err := ReadDirnames(dirname)
		assert.Nil(t, err)
		assert.Equal(t, expected, dirs)

		// clean up
		assert.Nil(t, RemoveAll(dirname))
	}

	// try a file
	{
		dirs, err := ReadDirnames(testfile)
		assert.Equal(t, []string{}, dirs)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to read directory names for"))
		assert.True(t, strings.Contains(err.Error(), ": readdirent"))
	}

	// writeonly directory
	{
		// Create writeonly directory
		dirname, err := MkdirP(path.Join(tmpDir, "dir"))
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dirname, 0222))

		dirs, err := ReadDirnames(dirname)
		assert.Equal(t, ([]string)(nil), dirs)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open directory"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Reset the permissions and remove the dir
		assert.Nil(t, os.Chmod(dirname, 0755))
		assert.Nil(t, RemoveAll(dirname))
	}

	// empty dirname
	{
		dirs, err := ReadDirnames("")
		assert.Equal(t, ([]string)(nil), dirs)
		assert.Equal(t, "empty string is an invalid path", err.Error())
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

	// Not a path
	assert.Equal(t, "foobar", SlicePath("foobar", -5, -1))
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

func TestWalkSkip(t *testing.T) {
	resetTest()

	// Create dirs to walk
	_, err := MkdirP(path.Join(tmpDir, "skipme"))
	assert.Nil(t, err)
	_, err = MkdirP(path.Join(tmpDir, "noskipme"))
	assert.Nil(t, err)

	// Call walk with skip
	results := []string{}
	err = Walk(tmpDir, func(p string, i *FileInfo, e error) error {
		if path.Base(p) == "skipme" {
			return filepath.SkipDir
		}
		results = append(results, path.Base(p))
		return nil
	})
	assert.Equal(t, []string{"temp", "noskipme"}, results)
}

func TestWalkDirPermissions(t *testing.T) {
	resetTest()

	// Create dirs to walk
	skipMe, err := MkdirP(path.Join(tmpDir, "skipme"))
	assert.Nil(t, err)
	_, err = MkdirP(path.Join(tmpDir, "noskipme"))
	assert.Nil(t, err)
	err = os.Chmod(skipMe, 0222)
	defer os.Chmod(skipMe, 0755)

	// Call walk with skip
	cnt := 0
	err = Walk(tmpDir, func(p string, i *FileInfo, e error) error {
		if path.Base(p) == "skipme" {
			cnt++
			if cnt == 2 {
				assert.True(t, strings.HasPrefix(e.Error(), "failed to open directory"))
				assert.True(t, strings.Contains(e.Error(), "permission denied"))
				// By passing it back to the walk function we abort
				return e
			}
			// There won't be a failure the first time we get called with skipme as
			// that is the user opportunity to skip the second call will be the read permissions
			// error
			assert.Nil(t, e)
		} else {
			assert.Nil(t, e)
		}
		return nil
	})
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
