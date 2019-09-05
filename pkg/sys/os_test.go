package sys

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

func TestCopyGlob(t *testing.T) {

	// single file to non-existing dst is a copy to not copy into
	{
		cleanTmpDir()
		// Create src dir and target file
		srcDir := path.Join(tmpDir, "src")
		_, err := MkdirP(srcDir)
		assert.Nil(t, err)
		_, err = Touch(path.Join(srcDir, "newfile1"))
		assert.Nil(t, err)

		// Now try to copy with bad glob pattern
		err = Copy(path.Join(tmpDir, "*/newfile*"), path.Join(tmpDir, "dst"))
		assert.Nil(t, err)

		// Validate resulting paths
		results, err := AllPaths(tmpDir)
		assert.Nil(t, err)
		tmpDirAbs, err := Abs(tmpDir)
		for i := range results {
			results[i] = strings.TrimPrefix(results[i], tmpDirAbs)
		}
		assert.Equal(t, []string{"", "/dst", "/src", "/src/newfile1"}, results)

		// Validate resulting file data
		data1, err := ReadString(path.Join(tmpDir, "src/newfile1"))
		assert.Nil(t, err)
		data2, err := ReadString(path.Join(tmpDir, "dst"))
		assert.Equal(t, data1, data2)
	}

	// multiple files to non-existing dst
	{
		cleanTmpDir()
		// Create src dir and target file
		srcDir := path.Join(tmpDir, "src")
		_, err := MkdirP(srcDir)
		assert.Nil(t, err)
		_, err = Touch(path.Join(srcDir, "newfile1"))
		assert.Nil(t, err)
		_, err = Touch(path.Join(srcDir, "newfile2"))
		assert.Nil(t, err)

		// Now try to copy with bad glob pattern
		err = Copy(path.Join(tmpDir, "*/newfile*"), path.Join(tmpDir, "dst"))
		assert.Nil(t, err)
		assert.FileExists(t, path.Join(tmpDir, "dst/newfile1"))
		assert.FileExists(t, path.Join(tmpDir, "dst/newfile2"))
	}

	// multiple files to pre-existing directory
	{
		cleanTmpDir()
		dst := path.Join(tmpDir)
		err := Copy("./*", dst)
		assert.Nil(t, err)

		expected, err := AllPaths(".")
		assert.Nil(t, err)
		results, err := AllPaths(tmpDir)
		assert.Nil(t, err)

		for i := range results {
			expected[i] = path.Base(expected[i])
			results[i] = path.Base(results[i])
		}
		assert.Equal(t, "sys", expected[0])
		assert.Equal(t, "temp", results[0])
		assert.Equal(t, expected[1:], results[1:])
	}
}

func TestCopyWithPermissionFailures(t *testing.T) {

	// try to create destination dirs in no write destination
	{
		cleanTmpDir()

		// Create src dir with no read permissions
		srcDir := path.Join(tmpDir, "src")
		_, err := MkdirP(srcDir)
		assert.Nil(t, err)
		_, err = Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Create dst dir with no write permissions
		dstDir := path.Join(tmpDir, "dst")
		_, err = MkdirP(dstDir)
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dstDir, 0444))

		// Now copy from src to sub dir under dst
		err = Copy(srcDir, path.Join(dstDir, "sub/file"))
		assert.True(t, strings.HasPrefix(err.Error(), "mkdir"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		// Fix the permission on the dstDir
		assert.Nil(t, os.Chmod(dstDir, 0755))
	}

	// read from no read permission source failure
	{
		cleanTmpDir()

		// Create src dir with no read permissions
		srcDir := path.Join(tmpDir, "src")
		_, err := MkdirP(srcDir)
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(srcDir, 0222))

		// Now try to copy from src
		err = Copy(srcDir, path.Join(tmpDir, "dst"))
		assert.True(t, strings.HasPrefix(err.Error(), "failed to read directory"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		// Fix the permission on the dstDir
		assert.Nil(t, os.Chmod(srcDir, 0755))
	}
}

func TestCopyDirLinksFailure(t *testing.T) {
	cleanTmpDir()

	// Create sub dir with link to it
	srcDir := path.Join(tmpDir, "src")
	_, err := MkdirP(srcDir)
	assert.Nil(t, err)
	linkDir := path.Join(tmpDir, "link")
	assert.Nil(t, os.Symlink(srcDir, linkDir))

	// Now create the destination with readonly permissions
	dstDir := path.Join(tmpDir, "dst")
	_, err = MkdirP(dstDir)
	assert.Nil(t, err)
	assert.Nil(t, os.Chmod(dstDir, 0444))

	// Now try to copy the linkDir to the dstDir
	err = Copy(linkDir, dstDir)
	assert.True(t, strings.HasPrefix(err.Error(), "symlink"))
	assert.True(t, strings.Contains(err.Error(), "permission denied"))

	// Fix the permission on the dstDir
	assert.Nil(t, os.Chmod(dstDir, 0755))
}

func TestCopyLinksRelativeNoFollow(t *testing.T) {
	cleanTmpDir()

	// temp/first/f0,f1
	firstDir, _ := MkdirP(path.Join(tmpDir, "first"))
	Touch(path.Join(firstDir, "f0"))
	Touch(path.Join(firstDir, "f1"))

	// temp/second/s0,s1
	secondDir, _ := MkdirP(path.Join(tmpDir, "second"))
	Touch(path.Join(secondDir, "s0"))
	Touch(path.Join(secondDir, "s1"))

	// Create sysmlink in first dir to second dir
	// temp/first/second => temp/second
	symlink := path.Join(tmpDir, "first", "second")
	assert.Nil(t, os.Symlink("../second", symlink))

	// Copy first dir to dst without following links
	{
		beforeInfo, err := Lstat(secondDir)
		assert.Nil(t, err)

		dstDir, _ := Abs(path.Join(tmpDir, "dst"))
		assert.Nil(t, Copy(firstDir, dstDir))

		// Compute results
		results, _ := AllPaths(dstDir)
		for i := 0; i < len(results); i++ {
			results[i] = SlicePath(results[i], -2, -1)
		}

		// Check that second is a link same as it was originally
		srcInfo, err := Lstat(path.Join(firstDir, "second"))
		assert.Nil(t, err)
		assert.True(t, srcInfo.IsSymlink())
		dstInfo, err := Lstat(path.Join(dstDir, "second"))
		assert.Nil(t, err)
		assert.True(t, dstInfo.IsSymlink())
		assert.Equal(t, srcInfo.Mode(), dstInfo.Mode())
		srcTarget, _ := srcInfo.SymlinkTarget()
		dstTarget, _ := dstInfo.SymlinkTarget()
		assert.Equal(t, srcTarget, dstTarget)
		assert.Equal(t, "../second", dstTarget)

		// Compare expected to results
		assert.Equal(t, []string{"temp/dst", "dst/f0", "dst/f1", "dst/second", "temp/second", "second/s0", "second/s1"}, results)

		afterInfo, err := Lstat(secondDir)
		assert.Nil(t, err)
		assert.Equal(t, beforeInfo.Mode(), afterInfo.Mode())
	}
}

func TestCopyLinksAbsNoFollow(t *testing.T) {
	cleanTmpDir()

	// temp/first/f0,f1
	firstDir, _ := MkdirP(path.Join(tmpDir, "first"))
	Touch(path.Join(firstDir, "f0"))
	Touch(path.Join(firstDir, "f1"))

	// temp/second/s0,s1
	secondDir, _ := MkdirP(path.Join(tmpDir, "second"))
	Touch(path.Join(secondDir, "s0"))
	Touch(path.Join(secondDir, "s1"))

	// Create sysmlink in first dir to second dir
	// temp/first/second => temp/second
	symlink := path.Join(tmpDir, "first", "second")
	assert.Nil(t, os.Symlink(secondDir, symlink))

	// Copy first dir to dst without following links
	{
		beforeInfo, err := Lstat(secondDir)
		assert.Nil(t, err)

		dstDir, _ := Abs(path.Join(tmpDir, "dst"))
		assert.Nil(t, Copy(firstDir, dstDir))

		// Compute results
		results, _ := AllPaths(dstDir)
		for i := 0; i < len(results); i++ {
			results[i] = SlicePath(results[i], -2, -1)
		}

		// Check that second is a link same as it was originally
		srcInfo, err := Lstat(path.Join(firstDir, "second"))
		assert.Nil(t, err)
		assert.True(t, srcInfo.IsSymlink())
		dstInfo, err := Lstat(path.Join(dstDir, "second"))
		assert.Nil(t, err)
		assert.True(t, dstInfo.IsSymlink())
		assert.Equal(t, srcInfo.Mode(), dstInfo.Mode())
		srcTarget, _ := srcInfo.SymlinkTarget()
		dstTarget, _ := dstInfo.SymlinkTarget()
		assert.Equal(t, srcTarget, dstTarget)
		assert.Equal(t, "test/temp/second", SlicePath(dstTarget, -3, -1))

		// Compare expected to results
		assert.Equal(t, []string{"temp/dst", "dst/f0", "dst/f1", "dst/second", "temp/second", "second/s0", "second/s1"}, results)

		afterInfo, err := Lstat(secondDir)
		assert.Nil(t, err)
		assert.Equal(t, beforeInfo.Mode(), afterInfo.Mode())
	}
}

func TestCopy(t *testing.T) {

	// invalid files
	{
		// invalid dst
		err := Copy("", "")
		assert.Equal(t, "empty string is an invalid path", err.Error())

		// invalid src
		err = Copy("", "foo")
		assert.Equal(t, "empty string is an invalid path", err.Error())

		// invalid file globbing i.e. doesn't exist
		err = Copy("foo", "bar")
		assert.True(t, strings.HasPrefix(err.Error(), "failed to get any sources for"))
	}

	{
		// test/temp/pkg does not exist
		// so Copy sys to pkg will be a clone
		cleanTmpDir()
		src := "."
		dst := path.Join(tmpDir, "pkg")

		assert.Nil(t, Copy(src, dst))
		srcPaths, err := AllPaths(src)
		assert.Nil(t, err)
		dstPaths, err := AllPaths(dst)
		assert.Nil(t, err)
		for i := range dstPaths {
			srcPaths[i] = path.Base(srcPaths[i])
			dstPaths[i] = path.Base(dstPaths[i])
		}
		assert.Equal(t, "sys", srcPaths[0])
		assert.Equal(t, "pkg", dstPaths[0])
		assert.Equal(t, srcPaths[1:], dstPaths[1:])
	}
	{
		// test/temp/pkg does exist
		// so Copy sys to pkg will be an into op
		cleanTmpDir()
		src := "."
		dst := path.Join(tmpDir, "pkg")
		MkdirP(dst)

		assert.Nil(t, Copy(src, dst))
		srcPaths, err := AllPaths(src)
		assert.Nil(t, err)
		dstPaths, err := AllPaths(path.Join(dst, "sys"))
		assert.Nil(t, err)
		for i := range dstPaths {
			srcPaths[i] = path.Base(srcPaths[i])
			dstPaths[i] = path.Base(dstPaths[i])
		}
		assert.Equal(t, srcPaths, dstPaths)
	}
}

func TestDarwin(t *testing.T) {
	if runtime.GOOS == "darwin" {
		assert.True(t, Darwin())
	} else {
		assert.False(t, Darwin())
	}
}

func TestLinux(t *testing.T) {
	if runtime.GOOS == "linux" {
		assert.True(t, Linux())
	} else {
		assert.False(t, Linux())
	}
}

func TestWindows(t *testing.T) {
	if runtime.GOOS == "windows" {
		assert.True(t, Windows())
	} else {
		assert.False(t, Windows())
	}
}

func TestCopyWithFileParentDoentExist(t *testing.T) {
	// test/temp/foo/bar/README.md does not exist and neither does its parent
	// so foo/bar will be created then Copy README.md to bar will be a clone
	cleanTmpDir()
	src := "./README.md"
	dst := path.Join(tmpDir, "foo/bar/readme")

	assert.False(t, Exists(dst))
	assert.Nil(t, Copy(src, dst))
	assert.True(t, Exists(dst))

	srcMD5, err := MD5(src)
	assert.Nil(t, err)
	dstMD5, err := MD5(dst)
	assert.Nil(t, err)
	assert.Equal(t, srcMD5, dstMD5)
}

func TestCopyFileParentDoentExist(t *testing.T) {
	// test/temp/foo/bar/README.md does not exist and neither does its parent
	// so foo/bar will be created then Copy README.md to bar will be a clone
	cleanTmpDir()
	src := "./README.md"
	dst := path.Join(tmpDir, "foo/bar/readme")

	assert.False(t, Exists(dst))
	_, err := CopyFile(src, dst)
	assert.Nil(t, err)
	assert.True(t, Exists(dst))

	srcMD5, err := MD5(src)
	assert.Nil(t, err)
	dstMD5, err := MD5(dst)
	assert.Nil(t, err)
	assert.Equal(t, srcMD5, dstMD5)
}

func TestCopyWithDirParentDoentExist(t *testing.T) {
	// test/temp/foo/bar/pkg does not exist and neither does its parent
	// so foo/bar will be created then Copy sys to pkg will be a clone
	cleanTmpDir()
	src := "."
	dst := path.Join(tmpDir, "foo/bar/pkg")

	assert.Nil(t, Copy(src, dst))
	srcPaths, err := AllPaths(src)
	assert.Nil(t, err)
	dstPaths, err := AllPaths(dst)
	assert.Nil(t, err)
	for i := range dstPaths {
		srcPaths[i] = path.Base(srcPaths[i])
		dstPaths[i] = path.Base(dstPaths[i])
	}
	assert.Equal(t, "sys", srcPaths[0])
	assert.Equal(t, "pkg", dstPaths[0])
	assert.Equal(t, srcPaths[1:], dstPaths[1:])
}

func TestCopyFile(t *testing.T) {
	cleanTmpDir()

	// copy symlink to readonly dest - failure
	{
		// Create link to a bogus file
		link := path.Join(tmpDir, "link")
		err := Symlink(path.Join(tmpDir, "bogus"), link)
		assert.Nil(t, err)

		// Create dst dir with readonly permissions
		dstDir, err := MkdirP(path.Join(tmpDir, "dst"))
		assert.Nil(t, err)
		err = os.Chmod(dstDir, 0444)
		assert.Nil(t, err)

		// Copy link to dst with readonly permssions and see failure
		result, err := CopyFile(link, dstDir)
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "symlink"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		// Reset permission so dst dir
		err = os.Chmod(dstDir, 0755)
		assert.Nil(t, err)
	}
	cleanTmpDir()

	// CopyFile symlink
	{
		// Create link to a bogus file
		link := path.Join(tmpDir, "link")
		err := Symlink(path.Join(tmpDir, "bogus"), link)
		assert.Nil(t, err)

		newlink := path.Join(tmpDir, "newlink")
		result, err := CopyFile(link, newlink)
		assert.Equal(t, SlicePath(newlink, -2, -1), SlicePath(result, -2, -1))
		assert.Nil(t, err)

		// Validate files and link locations
		linkInfo, err := Lstat(link)
		assert.Nil(t, err)
		assert.True(t, linkInfo.IsSymlink())
		assert.False(t, linkInfo.SymlinkTargetExists())
		linkTarget, err := linkInfo.SymlinkTarget()
		assert.Nil(t, err)
		assert.Equal(t, "../../test/temp/bogus", linkTarget)

		newlinkInfo, err := Lstat(newlink)
		assert.Nil(t, err)
		assert.True(t, newlinkInfo.IsSymlink())
		assert.False(t, newlinkInfo.SymlinkTargetExists())
		assert.False(t, SymlinkTargetExists(newlink))
		newlinkTarget, err := newlinkInfo.SymlinkTarget()
		assert.Nil(t, err)
		assert.Equal(t, "../../test/temp/bogus", newlinkTarget)

		// Create bogus file and test that symlink target exists
		_, err = Touch(path.Join(tmpDir, "bogus"))
		assert.Nil(t, err)
		assert.True(t, newlinkInfo.SymlinkTargetExists())
		assert.True(t, SymlinkTargetExists(newlink))
	}
	cleanTmpDir()

	// target file is not readable via permissions
	{
		// Write out a temp file
		err := WriteString(tmpfile, `This is a test of the emergency broadcast system.`)
		assert.Nil(t, err)

		// Revoke read permissions
		assert.Nil(t, os.Chmod(tmpfile, 0222))

		// Try to copy it and fail
		result, err := CopyFile(tmpfile, path.Join(tmpDir, "new"))
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open file"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))
		assert.Nil(t, os.Chmod(tmpfile, 0644))
		assert.Nil(t, Remove(tmpfile))
	}

	// source symlink and target doesn't exist
	{
		// Setup link to bogus file
		subDir, err := MkdirP(path.Join(tmpDir, "sub"))
		assert.Nil(t, err)
		linkDir := path.Join(tmpDir, "link")
		assert.Nil(t, Symlink(subDir, linkDir))

		result, err := CopyFile(linkDir, "new")
		assert.Equal(t, "", result)
		assert.Equal(t, "src target is not a regular file or a symlink to a file", err.Error())
	}

	cleanTmpDir()

	// empty destination
	{
		result, err := CopyFile(readme, "")
		assert.Equal(t, "", result)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// empty source
	{
		result, err := CopyFile("", "")
		assert.Equal(t, "", result)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// source doesn't exist
	{
		result, err := CopyFile(path.Join(tmpDir, "foo"), path.Join(tmpDir, "bar"))
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to execute Lstat against"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// empty info path
	{
		result, err := CopyFile(path.Join(tmpDir, "foo/foo"), "", newInfoOpt(&FileInfo{}))
		assert.Equal(t, "", result)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// pass in bad info
	{
		result, err := CopyFile(path.Join(tmpDir, "foo/foo"), "", newInfoOpt(&FileInfo{path: "foo/foo"}))
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to execute Lstat against"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// source is a directory
	{
		subdir, err := MkdirP(path.Join(tmpDir, "sub"))
		assert.Nil(t, err)
		result, err := CopyFile(subdir, path.Join(tmpDir, "bar"))
		assert.Equal(t, "", result)
		assert.Equal(t, "src target is not a regular file or a symlink to a file", err.Error())
	}

	// new destination name
	{
		result, err := CopyFile(readme, path.Join(tmpDir, "foo"))
		assert.Nil(t, err)
		assert.Equal(t, "temp/foo", SlicePath(result, -2, -1))
		assert.Nil(t, Remove(result))
	}

	// failed to create destination sub directory
	{
		subdir, err := MkdirP(path.Join(tmpDir, "sub"))
		assert.Nil(t, err)

		// Now make subdir readonly
		assert.Nil(t, os.Chmod(subdir, 0555))

		// Try to copy to a readonly directory
		result, err := CopyFile(readme, path.Join(subdir, "foo/bar"))
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "mkdir"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		// Fix permissions on subdir and remove it
		assert.Nil(t, os.Chmod(subdir, 0755))
		assert.Nil(t, RemoveAll(subdir))
	}

	// failed to stat destination
	{
		subdir, err := MkdirP(path.Join(tmpDir, "sub"))
		assert.Nil(t, err)

		// Now make subdir readonly
		assert.Nil(t, os.Chmod(subdir, 0444))

		// Try to copy to a readonly directory
		result, err := CopyFile(readme, path.Join(subdir, "foo/bar"))
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to Stat destination"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		// Fix permissions on subdir and remove it
		assert.Nil(t, os.Chmod(subdir, 0755))
		assert.Nil(t, RemoveAll(subdir))
	}

	// failed to create new file permission denied
	{
		subdir, err := MkdirP(path.Join(tmpDir, "sub"))
		assert.Nil(t, err)

		// Now make subdir readonly
		assert.Nil(t, os.Chmod(subdir, 0444))

		// Try to copy to a readonly directory
		result, err := CopyFile(readme, subdir)
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create file"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		// Fix permissions on subdir and remove it
		assert.Nil(t, os.Chmod(subdir, 0755))
		assert.Nil(t, RemoveAll(subdir))
	}

	// happy
	{
		// Copy regular file
		foo := path.Join(tmpDir, "foo")

		assert.False(t, Exists(foo))
		result, err := CopyFile(readme, foo)
		assert.Nil(t, err)
		assert.Equal(t, SlicePath(foo, -2, -1), SlicePath(result, -2, -1))
		assert.True(t, Exists(foo))

		srcMD5, err := MD5(readme)
		assert.Nil(t, err)
		dstMD5, err := MD5(foo)
		assert.Nil(t, err)
		assert.Equal(t, srcMD5, dstMD5)

		// Overwrite file
		result, err = CopyFile(testfile, foo)
		assert.Nil(t, err)
		assert.Equal(t, SlicePath(foo, -2, -1), SlicePath(result, -2, -1))
		srcMD5, err = MD5(testfile)
		assert.Nil(t, err)
		dstMD5, err = MD5(foo)
		assert.Nil(t, err)
		assert.Equal(t, srcMD5, dstMD5)
	}
}

func TestExists(t *testing.T) {
	assert.False(t, Exists("bob"))
	assert.True(t, Exists(readme))
}

func TestMD5(t *testing.T) {
	cleanTmpDir()

	// empty string
	{
		result, err := MD5("")
		assert.Equal(t, "", result)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// doesn't exist
	{
		result, err := MD5("foo")
		assert.Equal(t, "", result)
		assert.Equal(t, "file does not exist", err.Error())
	}

	// happy
	{
		f, _ := os.Create(tmpfile)
		defer f.Close()
		f.WriteString(`This is a test of the emergency broadcast system.`)

		expected := "067a8c38325b12159844261d16e5cb13"
		result, err := MD5(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	}

	// Remove read permissions from file
	{
		assert.Nil(t, os.Chmod(tmpfile, 0222))
		result, err := MD5(tmpfile)
		assert.Equal(t, "", result)
		assert.True(t, strings.HasPrefix(err.Error(), "failed opening target file"))
		assert.Nil(t, os.Chmod(tmpfile, 0644))
	}
}

func TestMkdirP(t *testing.T) {

	// happy
	{
		result, err := MkdirP(tmpDir)
		assert.Nil(t, err)
		assert.Equal(t, SlicePath(tmpDir, -2, -1), SlicePath(result, -2, -1))
		assert.True(t, Exists(tmpDir))
		assert.Nil(t, RemoveAll(result))
	}

	// permissions given
	{
		result, err := MkdirP(tmpDir, 0555)
		assert.Nil(t, err)
		assert.Equal(t, SlicePath(tmpDir, -2, -1), SlicePath(result, -2, -1))
		assert.True(t, Exists(tmpDir))
		mode := Mode(tmpDir)
		assert.Equal(t, os.ModeDir|os.FileMode(0555), mode)
	}

	// Remove read permissions from file
	{
		assert.Nil(t, os.Chmod(tmpDir, 0222))
		result, err := MkdirP(path.Join(tmpDir, "foo"))
		assert.Equal(t, "temp/foo", SlicePath(result, -2, -1))
		assert.True(t, strings.HasPrefix(err.Error(), "failed creating directories for"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))
		assert.Nil(t, os.Chmod(tmpDir, 0755))
	}

	// HOME not set
	{
		// unset HOME
		home := os.Getenv("HOME")
		os.Unsetenv("HOME")
		assert.Equal(t, "", os.Getenv("HOME"))
		defer os.Setenv("HOME", home)

		result, err := MkdirP("~/")
		assert.Equal(t, "failed to expand the given path ~/: failed to compute the user's home directory: $HOME is not defined", err.Error())
		assert.Equal(t, "", result)
	}
}

func TestMove(t *testing.T) {
	cleanTmpDir()

	// Copy file in to tmpDir then rename in same location
	assert.Nil(t, Copy(testfile, tmpDir))
	newTestFile := path.Join(tmpDir, "testfile")

	srcMd5, _ := MD5(newTestFile)
	assert.True(t, Exists(newTestFile))
	assert.False(t, Exists(tmpfile))
	result, err := Move(newTestFile, tmpfile)
	assert.Nil(t, err)
	assert.Equal(t, tmpfile, result)
	assert.True(t, Exists(tmpfile))
	dstMd5, err := MD5(tmpfile)
	assert.Nil(t, err)
	assert.False(t, Exists(newTestFile))
	assert.Equal(t, srcMd5, dstMd5)

	// Now create a sub directory and move it there
	subDir := path.Join(tmpDir, "sub")
	MkdirP(subDir)
	newfile, err := Move(tmpfile, subDir)
	assert.Nil(t, err)
	assert.Equal(t, path.Join(subDir, path.Base(tmpfile)), newfile)
	assert.False(t, Exists(tmpfile))
	assert.True(t, Exists(newfile))
	dstMd5, _ = MD5(newfile)
	assert.Equal(t, srcMd5, dstMd5)

	// permission denied
	assert.Nil(t, os.Chmod(subDir, 0222))
	result, err = Move(newfile, tmpfile)
	assert.Equal(t, "", result)
	assert.True(t, strings.HasPrefix(err.Error(), "failed renaming file"))
	assert.True(t, strings.Contains(err.Error(), "permission denied"))
	assert.Nil(t, os.Chmod(subDir, 0755))
}

func TestPwd(t *testing.T) {
	assert.Equal(t, "sys", path.Base(Pwd()))
}

func TestReadBytes(t *testing.T) {
	cleanTmpDir()

	// empty string
	{
		data, err := ReadBytes("")
		assert.Equal(t, ([]byte)(nil), data)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// invalid file
	{
		data, err := ReadBytes("foo")
		assert.Equal(t, ([]byte)(nil), data)
		assert.True(t, strings.HasPrefix(err.Error(), "failed reading the file"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// happy
	{
		// Write out test data
		assert.Nil(t, WriteString(tmpfile, "this is a test"))

		// Read the file back in and validate
		data, err := ReadBytes(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, "this is a test", string(data))
	}
}

func TestReadLines(t *testing.T) {
	cleanTmpDir()

	// empty string
	{
		data, err := ReadLines("")
		assert.Equal(t, ([]string)(nil), data)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// invalid file
	{
		data, err := ReadLines("foo")
		assert.Equal(t, ([]string)(nil), data)
		assert.True(t, strings.HasPrefix(err.Error(), "failed reading the file"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// happy
	{
		lines, err := ReadLines(testfile)
		assert.Nil(t, err)
		assert.Equal(t, 18, len(lines))
	}
}

func TestReadString(t *testing.T) {
	cleanTmpDir()

	// empty string
	{
		data, err := ReadString("")
		assert.Equal(t, "", data)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// invalid file
	{
		data, err := ReadString("foo")
		assert.Equal(t, "", data)
		assert.True(t, strings.HasPrefix(err.Error(), "failed reading the file"))
		assert.True(t, strings.Contains(err.Error(), "no such file or directory"))
	}

	// happy
	{
		// Write out test data
		assert.Nil(t, WriteString(tmpfile, "this is a test"))

		// Read the file back in and validate
		data, err := ReadString(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, "this is a test", data)
	}
}

func TestRemove(t *testing.T) {
	cleanTmpDir()

	// Write out test data
	assert.Nil(t, WriteString(tmpfile, "this is a test"))
	assert.True(t, Exists(tmpfile))

	// Now remove the file and validate
	assert.Nil(t, Remove(tmpfile))
	assert.False(t, Exists(tmpfile))
}

func TestSymlink(t *testing.T) {
	cleanTmpDir()

	_, err := Touch(tmpfile)
	assert.Nil(t, err)

	// Create file symlink
	newfilelink := path.Join(tmpDir, "filelink")
	assert.Nil(t, Symlink(tmpfile, newfilelink))
	assert.True(t, IsSymlink(newfilelink))
	assert.True(t, IsSymlinkFile(newfilelink))
	assert.False(t, IsSymlinkDir(newfilelink))

	// Create dir symlink
	subdir := path.Join(tmpDir, "sub")
	_, err = MkdirP(subdir)
	assert.Nil(t, err)
	newdirlink := path.Join(tmpDir, "sublink")
	assert.Nil(t, Symlink(subdir, newdirlink))
	assert.True(t, IsSymlink(newdirlink))
	assert.False(t, IsSymlinkFile(newdirlink))
	assert.True(t, IsSymlinkDir(newdirlink))
}

func TestTouch(t *testing.T) {
	cleanTmpDir()

	// Force failure of Close via monkey patch
	{
		OneShotForceIOCloseError()
		_, err := Touch(tmpfile)
		assert.Equal(t, fmt.Sprintf("failed closing file %s: invalid argument", tmpfile), err.Error())

		// Clean up
		err = Remove(tmpfile)
		assert.Nil(t, err)
	}

	// empty string
	{
		result, err := Touch("")
		assert.Equal(t, "", result)
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// permission denied
	{
		// Create the tmpfile
		result, err := Touch(tmpfile)
		assert.Equal(t, SlicePath(tmpfile, -2, -1), SlicePath(result, -2, -1))

		// Now try to truncate it after setting to readonly
		assert.Nil(t, os.Chmod(tmpfile, 0444))
		result, err = Touch(tmpfile)
		assert.Equal(t, SlicePath(tmpfile, -2, -1), SlicePath(result, -2, -1))
		assert.True(t, strings.HasPrefix(err.Error(), "failed creating/truncating file"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))
		assert.Nil(t, os.Chmod(tmpfile, 0755))
		assert.Nil(t, Remove(tmpfile))
	}

	// happy
	{
		// Doesn't exist so create
		assert.False(t, Exists(tmpfile))
		_, err := Touch(tmpfile)
		assert.Nil(t, err)
		assert.True(t, Exists(tmpfile))

		// Truncate and re-create it
		_, err = Touch(tmpfile)
		assert.Nil(t, err)
	}
}

func TestWriteBytes(t *testing.T) {
	cleanTmpDir()

	// attemp to write to a readonly dst
	{
		dstDir, err := MkdirP(path.Join(tmpDir, "dst"))
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dstDir, 0444))

		err = WriteBytes(path.Join(dstDir, "file"), []byte("test"))
		assert.True(t, strings.HasPrefix(err.Error(), "failed writing bytes to file"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		assert.Nil(t, os.Chmod(dstDir, 0444))
		assert.Nil(t, Remove(dstDir))
	}

	// empty target
	{
		err := WriteBytes("", []byte("test"))
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// happy
	{
		// Read and write file
		data, err := ioutil.ReadFile(testfile)
		assert.Nil(t, err)
		err = WriteBytes(tmpfile, data, 0644)
		assert.Nil(t, err)

		// Test the resulting file
		data2, err := ioutil.ReadFile(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, data, data2)
	}
}

func TestWriteLines(t *testing.T) {
	cleanTmpDir()

	// attemp to write to a readonly dst
	{
		dstDir, err := MkdirP(path.Join(tmpDir, "dst"))
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dstDir, 0444))

		err = WriteLines(path.Join(dstDir, "file"), []string{"test"})
		assert.True(t, strings.HasPrefix(err.Error(), "failed writing lines to file"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		assert.Nil(t, os.Chmod(dstDir, 0444))
		assert.Nil(t, Remove(dstDir))
	}

	// empty target
	{
		err := WriteLines("", []string{"test"})
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// happy
	{
		lines, err := ReadLines(testfile)
		assert.Nil(t, err)
		assert.Equal(t, 18, len(lines))
		err = WriteLines(tmpfile, lines, 0644)
		assert.Nil(t, err)
		{
			lines2, err := ReadLines(tmpfile)
			assert.Nil(t, err)
			assert.Equal(t, lines, lines2)
		}
	}
}

func TestWriteStream(t *testing.T) {

	// force sync and close errors
	{
		OneShotForceIOSyncError()
		OneShotForceIOCloseError()
		reader, err := os.Open(testfile)
		assert.Nil(t, err)
		err = WriteStream(reader, tmpfile)
		assert.Nil(t, reader.Close())
		assert.True(t, strings.HasPrefix(err.Error(), "failed to close file"))
		assert.True(t, strings.Contains(err.Error(), ": failed syncing stream to file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
		assert.Nil(t, os.Remove(tmpfile))
	}

	// force sync and close errors
	{
		OneShotForceIOSyncError()
		OneShotForceIOCloseError()
		reader, err := os.Open(testfile)
		assert.Nil(t, err)
		err = WriteStream(reader, tmpfile)
		assert.Nil(t, reader.Close())
		assert.True(t, strings.HasPrefix(err.Error(), "failed to close file"))
		assert.True(t, strings.Contains(err.Error(), ": failed syncing stream to file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
		assert.Nil(t, os.Remove(tmpfile))
	}

	// attemp to read from iotest TimeoutReader and force failure close
	{
		OneShotForceIOCloseError()
		reader, err := os.Open(testfile)
		assert.Nil(t, err)
		testReader := iotest.TimeoutReader(reader)
		err = WriteStream(testReader, tmpfile)
		assert.Nil(t, reader.Close())
		assert.True(t, strings.HasPrefix(err.Error(), "failed to close file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": failed copying stream data: timeout"))
		assert.Nil(t, os.Remove(tmpfile))
	}

	// attemp to write to a readonly dst
	{
		dstDir, err := MkdirP(path.Join(tmpDir, "dst"))
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dstDir, 0444))

		err = WriteStream(&os.File{}, path.Join(dstDir, "file"))
		assert.True(t, strings.HasPrefix(err.Error(), "failed opening file"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		assert.Nil(t, os.Chmod(dstDir, 0444))
		assert.Nil(t, Remove(dstDir))
	}

	// empty destination file
	{
		err := WriteStream(&os.File{}, "")
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	var expectedData []byte
	expectedData, err := ioutil.ReadFile(testfile)
	assert.Nil(t, err)

	// No file exists
	{
		cleanTmpDir()

		// Read and write file
		reader, err := os.Open(testfile)
		assert.Nil(t, err)
		err = WriteStream(reader, tmpfile, 0644)
		assert.Nil(t, reader.Close())
		assert.Nil(t, err)

		// Test the resulting file
		var data []byte
		data, err = ioutil.ReadFile(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, expectedData, data)
		assert.Nil(t, os.Remove(tmpfile))
	}

	// Overwrite and truncate file
	{
		// Read and write file
		reader, err := os.Open(testfile)
		assert.Nil(t, err)
		err = WriteStream(reader, tmpfile)
		assert.Nil(t, reader.Close())
		assert.Nil(t, err)

		// Test the resulting file
		var data []byte
		data, err = ioutil.ReadFile(testfile)
		assert.Nil(t, err)
		assert.Equal(t, expectedData, data)
		assert.Nil(t, os.Remove(tmpfile))
	}
}

func TestWriteString(t *testing.T) {
	cleanTmpDir()

	// attemp to write to a readonly dst
	{
		dstDir, err := MkdirP(path.Join(tmpDir, "dst"))
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dstDir, 0444))

		err = WriteString(path.Join(dstDir, "file"), "test")
		assert.True(t, strings.HasPrefix(err.Error(), "failed writing string to file"))
		assert.True(t, strings.Contains(err.Error(), "permission denied"))

		assert.Nil(t, os.Chmod(dstDir, 0444))
		assert.Nil(t, Remove(dstDir))
	}

	// empty target
	{
		err := WriteString("", "test")
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// happy
	{
		// Read and write file
		data, err := ioutil.ReadFile(testfile)
		assert.Nil(t, err)
		err = WriteString(tmpfile, string(data), 0644)
		assert.Nil(t, err)

		// Test the resulting file
		data2, err := ioutil.ReadFile(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, data, data2)
	}
}

func cleanTmpDir() {
	if Exists(tmpDir) {
		RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
}
