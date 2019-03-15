package sys

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		//mst, _ := time.LoadLocation("MST")
		//expected := time.Date(2018, time.December, 10, 9, 27, 55, 0, mst)
		//assert.Equal(t, time.Duration(0), expected.Sub(info.ModeTime()))
	}
}

func TestIsDir(t *testing.T) {
	{
		// FileInfo
		info, err := Lstat(readme)
		assert.Nil(t, err)
		assert.False(t, info.IsDir())

		info, err = Lstat("../../")
		assert.Nil(t, err)
		assert.True(t, info.IsDir())
	}
	{
		// Standalone
		assert.False(t, IsDir(readme))
		assert.True(t, IsDir("../.."))
	}
}

func TestIsFile(t *testing.T) {
	{
		// FileInfo
		info, err := Lstat(readme)
		assert.Nil(t, err)
		assert.True(t, info.IsFile())

		info, err = Lstat("../../")
		assert.Nil(t, err)
		assert.False(t, info.IsFile())
	}
	{
		// Standalone
		assert.True(t, IsFile(readme))
		assert.False(t, IsFile("../.."))
	}
}

func TestIsSymlink(t *testing.T) {
	cleanTmpDir()

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
	cleanTmpDir()

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
	cleanTmpDir()

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

	// Symlink to a file
	cleanTmpDir()
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
	}

	// Symlink to a dir
	cleanTmpDir()
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
