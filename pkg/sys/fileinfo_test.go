package sys

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testTime = time.Date(2018, time.May, 13, 1, 2, 3, 4, time.UTC)

func TestPath(t *testing.T) {
	info, err := Lstat(testfile)
	assert.Nil(t, err)
	assert.Equal(t, "test/testfile", SlicePath(info.Path, -2, -1))
}

func TestFileInfoInterface(t *testing.T) {
	{
		// file doesn't exist ensure error is returned
		info, err := Lstat("")
		assert.NotNil(t, err)
		assert.Nil(t, info)
	}
	{
		// regular file
		info, err := Lstat(testfile)
		assert.Nil(t, err)
		assert.NotNil(t, info)
		assert.Equal(t, "testfile", info.Name())
		assert.Equal(t, int64(604), info.Size())
		assert.Equal(t, os.FileMode(0x1a4), info.Mode())
	}
}

func TestSize(t *testing.T) {

	// class
	{
		info, err := Lstat(testfile)
		assert.Nil(t, err)
		assert.Equal(t, int64(604), info.Size())
	}

	// relative path file
	{
		assert.Equal(t, int64(604), Size(testfile))
	}

	// Ensure expansion is happening
	{
		home, err := UserHome()
		assert.Nil(t, err)
		target, err := Abs(testfile)
		assert.Nil(t, err)
		target = "~" + strings.TrimPrefix(target, home)

		assert.Equal(t, int64(604), Size(target))
	}
}

func TestMode(t *testing.T) {

	// class
	{
		assert.Nil(t, Copy(testfile, tmpfile))
		assert.Nil(t, os.Chmod(tmpfile, 0644))

		info, err := Lstat(testfile)
		assert.Nil(t, err)
		assert.Equal(t, os.FileMode(0644), info.Mode())
	}

	// relative path file
	{
		assert.Equal(t, os.FileMode(0644), Mode(testfile))
	}

	// Ensure expansion is happening
	{
		home, err := UserHome()
		assert.Nil(t, err)
		target, err := Abs(tmpfile)
		assert.Nil(t, err)
		target = "~" + strings.TrimPrefix(target, home)

		assert.Equal(t, os.FileMode(0644), Mode(target))
	}
}

func TestIsDir(t *testing.T) {

	// FileInfo
	{
		info, err := Lstat(readme)
		assert.Nil(t, err)
		assert.False(t, info.IsDir())

		info, err = Lstat("../../")
		assert.Nil(t, err)
		assert.True(t, info.IsDir())
	}

	// Standalone
	{
		// Ensure expansion is happening
		home, err := UserHome()
		assert.Nil(t, err)
		target, err := Abs(readme)
		assert.Nil(t, err)
		target = "~" + strings.TrimPrefix(target, home)

		assert.False(t, IsDir(target))
		assert.True(t, IsDir(path.Dir(target)))
	}
}

func TestAnyDir(t *testing.T) {
	// reset state
	dir := filepath.Join(tmpDir, "AnyDir")
	RemoveAll(dir)
	dir, err := MkdirP(dir)
	assert.NoError(t, err)

	// create a dir to test
	dir1, err := MkdirP(filepath.Join(dir, "dir1"))
	assert.NoError(t, err)

	// create a link to test
	link1 := filepath.Join(dir, "link1")
	err = Symlink(dir1, link1)
	assert.NoError(t, err)

	// error case
	{
		assert.False(t, AnyDir(path.Join(tmpDir, "bogus")))
	}

	// FileInfo
	{
		// regular dir
		info, err := Lstat(dir1)
		assert.NoError(t, err)
		assert.True(t, info.AnyDir())

		// link not a dir
		info, err = Lstat(link1)
		assert.NoError(t, err)
		assert.False(t, info.IsDir())

		// link points to dir
		assert.True(t, info.AnyDir())
	}

	// Standalone
	{
		// regular dir
		assert.True(t, AnyDir(dir1))

		// link not a dir
		assert.False(t, IsDir(link1))

		// link points to dir
		assert.True(t, AnyDir(link1))
	}
}

func TestIsFile(t *testing.T) {
	resetTest()

	// sad
	{
		// Standalone
		assert.False(t, IsFile(path.Join(tmpDir, "bogus")))
	}

	// happy
	{
		// FileInfo
		info, err := Lstat(readme)
		assert.Nil(t, err)
		assert.True(t, info.IsFile())

		info, err = Lstat("../../")
		assert.Nil(t, err)
		assert.False(t, info.IsFile())
	}

	// Standalone
	{
		// Ensure expansion is happening
		home, err := UserHome()
		assert.Nil(t, err)
		target, err := Abs(readme)
		assert.Nil(t, err)
		target = "~" + strings.TrimPrefix(target, home)

		assert.True(t, IsFile(target))
		assert.False(t, IsFile(path.Dir(target)))
	}
}

func TestAnyFile(t *testing.T) {
	// reset state
	dir := filepath.Join(tmpDir, "AnyFile")
	RemoveAll(dir)
	dir, err := MkdirP(dir)
	assert.NoError(t, err)

	// create a file to test
	file1 := filepath.Join(dir, "file1")
	file1, err = Touch(file1)
	assert.NoError(t, err)

	// create a link to test
	link1 := filepath.Join(dir, "link1")
	err = Symlink(file1, link1)
	assert.NoError(t, err)

	// error case
	{
		assert.False(t, AnyFile(path.Join(tmpDir, "bogus")))
	}

	// FileInfo
	{
		// regular file
		info, err := Lstat(file1)
		assert.NoError(t, err)
		assert.True(t, info.AnyFile())

		// link not a file
		info, err = Lstat(link1)
		assert.NoError(t, err)
		assert.False(t, info.IsFile())

		// link points to file
		assert.True(t, info.AnyFile())
	}

	// Standalone
	{
		// regular file
		assert.True(t, AnyFile(file1))

		// link not a file
		assert.False(t, IsFile(link1))

		// link points to file
		assert.True(t, AnyFile(link1))
	}
}

func TestIsSymlink(t *testing.T) {
	resetTest()

	// sad
	{
		// Standalone
		assert.False(t, IsSymlink(path.Join(tmpDir, "bogus")))
	}

	// Standalone
	{
		// Check a regular file
		assert.False(t, IsSymlink(testfile))
		assert.True(t, IsFile(testfile))

		// Create sysmlink
		symlink := path.Join(tmpDir, "symlink")
		os.Symlink(testfile, symlink)
		assert.True(t, IsSymlink(symlink))

		// Check that IsFile works
		assert.False(t, IsFile(symlink))
	}

	// FileInfo
	{
		// Check a regular file
		info, err := Lstat(testfile)
		assert.Nil(t, err)
		assert.False(t, info.IsSymlink())
		assert.True(t, info.IsFile())

		// Create sysmlink
		symlink := path.Join(tmpDir, "symlink")
		os.Symlink(testfile, symlink)
		info, err = Lstat(symlink)
		assert.Nil(t, err)

		// Check that IsFile works
		assert.True(t, info.IsSymlink())
		assert.False(t, info.IsFile())
	}
}

func TestIsSymlinkDir(t *testing.T) {
	resetTest()

	// sad
	{
		// Standalone
		assert.False(t, IsSymlinkDir(path.Join(tmpDir, "bogus")))
	}

	// Create a directory
	dir := path.Join(tmpDir, "dir")
	MkdirP(dir)

	// Create a symlink to the first directory
	symlink := path.Join(tmpDir, "symlink")
	os.Symlink(dir, symlink)

	// Test that the symlink points to a dir
	info, err := Lstat(symlink)
	assert.Nil(t, err)
	assert.True(t, info.IsSymlink())
	assert.True(t, info.IsSymlinkDir())
	assert.False(t, info.IsSymlinkFile())
	assert.True(t, IsSymlinkDir(symlink))
	assert.False(t, IsSymlinkFile(symlink))
}

func TestIsSymlinkFile(t *testing.T) {
	resetTest()

	// sad
	{
		// Standalone
		assert.False(t, IsSymlinkFile(path.Join(tmpDir, "bogus")))
	}

	// Create a symlink to a file
	symlink := path.Join(tmpDir, "symlink")
	os.Symlink(testfile, symlink)

	// Test that the symlink points to a file
	info, err := Lstat(symlink)
	assert.Nil(t, err)
	assert.True(t, info.IsSymlink())
	assert.False(t, info.IsSymlinkDir())
	assert.True(t, info.IsSymlinkFile())
	assert.False(t, IsSymlinkDir(symlink))
	assert.True(t, IsSymlinkFile(symlink))
}

func TestSymlinkTarget(t *testing.T) {
	resetTest()

	// link doesn't exist
	{
		result, err := SymlinkTarget(path.Join(tmpDir, "bogus"))
		assert.Empty(t, result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to execute Lstat against"))
		assert.True(t, strings.HasSuffix(err.Error(), ": no such file or directory"))
	}

	// // Force readlink error
	// {
	// 	symlink := path.Join(tmpDir, "symlink")
	// 	assert.Nil(t, os.Symlink(testfile, symlink))

	// 	info, err := Lstat(symlink)
	// 	assert.Nil(t, err)
	// 	test.OneShotForceOSReadlinkError()
	// 	target, err := info.SymlinkTarget()
	// 	assert.Empty(t, target)
	// 	assert.Equal(t, "failed to read the link target", err.Error())

	// 	assert.Nil(t, Remove(symlink))
	// }

	// Symlink to a file
	{
		info, err := Lstat(testfile)
		assert.Nil(t, err)
		target, err := info.SymlinkTarget()
		assert.Empty(t, target)
		assert.Equal(t, "not a symlink", err.Error())
	}

	// Symlink to a file
	{
		symlink := path.Join(tmpDir, "symlink")
		os.Symlink(testfile, symlink)

		info, _ := Lstat(symlink)
		target, err := info.SymlinkTarget()
		assert.Nil(t, err)
		assert.Equal(t, "../../test/testfile", target)
		target, err = SymlinkTarget(symlink)
		assert.Nil(t, err)
		assert.Equal(t, "../../test/testfile", target)

		assert.Nil(t, Remove(symlink))
	}

	// Symlink to a dir
	{
		dir := path.Join(tmpDir, "dir")
		MkdirP(dir)
		symlink := path.Join(tmpDir, "symlink")
		os.Symlink(dir, symlink)

		info, _ := Lstat(symlink)
		target, err := info.SymlinkTarget()
		assert.Nil(t, err)
		assert.Equal(t, "../../test/temp/dir", target)
		target, err = SymlinkTarget(symlink)
		assert.Nil(t, err)
		assert.Equal(t, "../../test/temp/dir", target)
	}
}

func TestSymlinkTargetExists(t *testing.T) {

	// link doesn't exist
	{
		assert.False(t, SymlinkTargetExists(path.Join(tmpDir, "bogus")))
	}
}
