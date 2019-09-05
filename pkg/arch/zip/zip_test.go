package zip

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/bouk/monkey"
	"github.com/phR0ze/n"
	"github.com/phR0ze/n/pkg/sys"
	"github.com/phR0ze/n/pkg/test"
	"github.com/stretchr/testify/assert"
)

var testZipfile1 = "../../../test/file1.zip"
var testZipfile2 = "../../../test/file2.zip"

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"
var tempZipfile1 = "../../../test/temp/file1.zip"
var tempZipfile2 = "../../../test/temp/file2.zip"

func TestCreateSad(t *testing.T) {
	prepTmpDir()

	// force zip.Create error
	{
		// Create new source directory/file
		srcDir := path.Join(tmpDir, "src")
		_, err := sys.MkdirP(srcDir)
		assert.Nil(t, err)
		_, err = sys.Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Now attempt to create th zip but force the io.Copy error
		OneShotForceZipCreateError()
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to add target file"))
		assert.True(t, strings.HasSuffix(err.Error(), "to zip: invalid argument"))

		// Clean up
		assert.Nil(t, sys.RemoveAll(srcDir))
		assert.Nil(t, sys.Remove(tmpfile))
	}

	// force io.Copy error
	{
		// Create new source directory/file
		srcDir := path.Join(tmpDir, "src")
		_, err := sys.MkdirP(srcDir)
		assert.Nil(t, err)
		_, err = sys.Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Now attempt to create th zip but force the io.Copy error
		test.OneShotForceIOCopyError()
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to copy data from reader to writer for zip target"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))

		// Clean up
		assert.Nil(t, sys.RemoveAll(srcDir))
		assert.Nil(t, sys.Remove(tmpfile))
	}

	// attempt to read writeonly source file
	{
		// Create new source directory
		srcDir := path.Join(tmpDir, "src")
		_, err := sys.MkdirP(srcDir)
		assert.Nil(t, err)

		// Add a file to the source directory
		srcfile, err := sys.Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Now set the source file to be write only
		assert.Nil(t, os.Chmod(srcfile, 0222))

		// Now attempt to read from the write only source file
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open target file"))
		assert.True(t, strings.Contains(err.Error(), "for zip: open"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(srcDir, 0755))
		assert.Nil(t, sys.RemoveAll(srcDir))
		assert.Nil(t, sys.Remove(tmpfile))
	}

	// attempt to read writeonly source directory
	{
		// Create new source directory
		srcDir := path.Join(tmpDir, "src")
		_, err := sys.MkdirP(srcDir)
		assert.Nil(t, err)

		// Add a file to the source directory
		_, err = sys.Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Now set the source directory to be write only
		assert.Nil(t, os.Chmod(srcDir, 0222))

		// Now attempt to read from the write only source
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to read directory"))
		assert.True(t, strings.Contains(err.Error(), "to add files from: open"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(srcDir, 0755))
		assert.Nil(t, sys.RemoveAll(srcDir))
		assert.Nil(t, sys.Remove(tmpfile))
	}

	// attempt to create zip in read only directory
	{
		// Create new destination directory
		dstDir := path.Join(tmpDir, "dst")
		_, err := sys.MkdirP(dstDir)
		assert.Nil(t, err)

		// Now set the destination to be readonly
		assert.Nil(t, os.Chmod(dstDir, 0444))

		// Now attempt to create the zip in the read only directory
		err = Create(path.Join(dstDir, "zip"), tmpfile)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create zipfile"))
		assert.True(t, strings.Contains(err.Error(), ": open /"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(dstDir, 0755))
		assert.Nil(t, sys.RemoveAll(dstDir))
	}

	// source path is empty
	{
		assert.Equal(t, "empty string is an invalid path", Create(tmpfile, "").Error())
	}

	// target zipfile is empty
	{
		assert.Equal(t, "empty string is an invalid path", Create("", "").Error())
	}
}

func TestCreateHappy(t *testing.T) {
	prepTmpDir()

	sys.Copy("../../net", tmpDir)

	// Create the new zipfile
	src := path.Join(tmpDir, "net")
	err := Create(tmpfile, src)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tmpfile))

	// Remove zip target files
	sys.RemoveAll(src)
	assert.False(t, sys.Exists(src))

	// Call out and extract the files
	_, err = sys.ExecOut("unzip %s -d %s", tmpfile, tmpDir)
	assert.Nil(t, err)
	paths, err := sys.AllPaths(tmpDir)
	assert.Nil(t, err)
	result := n.S(paths).Map(func(x n.O) n.O { return sys.SlicePath(x.(string), -3, -1) }).ToStrs()
	expected := []string{
		"n/test/temp",
		"test/temp/.tmp",
		"test/temp/agent",
		"temp/agent/agent.go",
		"test/temp/mech",
		"temp/mech/example",
		"mech/example/mech.go",
		"temp/mech/mech.go",
		"temp/mech/mech_test.go",
		"temp/mech/page.go",
		"test/temp/net.go",
		"test/temp/net_test.go",
	}
	assert.Equal(t, expected, result)
}

func TestExtractAllSad(t *testing.T) {
	prepTmpDir()

	// force os.Chtimes error
	{
		// Copy over zipfile
		assert.Nil(t, sys.Copy(testZipfile1, tempZipfile1))

		test.OneShotForceOSChtimesError()
		err := ExtractAll(tempZipfile1, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to set file access times for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// force os.Chmod error
	{
		test.OneShotForceOSChmodError()
		err := ExtractAll(tempZipfile1, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to set file mode for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// force os.Copy error
	{
		test.OneShotForceIOCopyError()
		err := ExtractAll(tempZipfile1, tmpDir)
		assert.Equal(t, "failed to copy data from zip to disk: invalid argument", err.Error())
	}

	// force zip.FileOpen error
	{
		OneShotForceZipFileOpenError()
		err := ExtractAll(tempZipfile1, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open zipfile target"))
		assert.True(t, strings.HasSuffix(err.Error(), "for reading: invalid argument"))
	}

	// force os.Create error
	{
		test.OneShotForceOSCreateError()
		err := ExtractAll(tempZipfile1, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create file"))
		assert.True(t, strings.HasSuffix(err.Error(), "from zipfile: invalid argument"))
	}

	// force zip.OpenReader error
	{
		OneShotForceZipOpenReaderError()
		err := ExtractAll(tempZipfile1, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open zipfile"))
		assert.True(t, strings.HasSuffix(err.Error(), "for reading: invalid argument"))
	}

	// attempt to read writeonly source file
	{
		// Create new source directory
		srcDir := path.Join(tmpDir, "src")
		_, err := sys.MkdirP(srcDir)
		assert.Nil(t, err)

		// Add a file to the source directory
		srcfile, err := sys.Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Now set the source file to be write only
		assert.Nil(t, os.Chmod(srcfile, 0222))

		// Now attempt to read from the write only source file
		test.OneShotForceOSCloseError()
		OneShotForceZipCloseError()
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to close file writer: failed to close zipfile writer: failed to open target file"))
		assert.True(t, strings.Contains(err.Error(), "for zip: open"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(srcDir, 0755))
		assert.Nil(t, sys.RemoveAll(srcDir))
		assert.Nil(t, sys.Remove(tmpfile))
	}

	// attempt to read writeonly zipfile
	{
		// Copy over zipfile and set as write only
		assert.Nil(t, sys.Copy(testZipfile1, tempZipfile1))
		assert.Nil(t, os.Chmod(tempZipfile1, 0222))

		// Now attempt to read the write only zip
		err := ExtractAll(tempZipfile1, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open zipfile 'file1.zip' to detect zip data:"))
		assert.True(t, strings.HasSuffix(err.Error(), "file1.zip: permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(tempZipfile1, 0644))
		assert.Nil(t, sys.Remove(tempZipfile1))
	}

	// Destination path is empty
	{
		assert.Equal(t, "empty string is an invalid path", ExtractAll(tmpfile, "").Error())
	}

	// target zipfile is empty
	{
		assert.Equal(t, "empty string is an invalid path", ExtractAll("", "").Error())
	}
}

func TestExtractAll(t *testing.T) {
	prepTmpDir()
	sys.Copy("../../net", tmpDir)

	// Create the new zipfile
	src := path.Join(tmpDir, "net")
	err := Create(tmpfile, src)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tmpfile))

	// Remove zip target files
	sys.RemoveAll(src)
	assert.False(t, sys.Exists(src))

	// Now extract the files and validate
	err = ExtractAll(tmpfile, tmpDir)
	assert.Nil(t, err)
	paths, err := sys.AllPaths(tmpDir)
	assert.Nil(t, err)
	result := n.S(paths).Map(func(x n.O) n.O { return sys.SlicePath(x.(string), -3, -1) }).ToStrs()
	expected := []string{
		"n/test/temp",
		"test/temp/.tmp",
		"test/temp/agent",
		"temp/agent/agent.go",
		"test/temp/mech",
		"temp/mech/example",
		"mech/example/mech.go",
		"temp/mech/mech.go",
		"temp/mech/mech_test.go",
		"temp/mech/page.go",
		"test/temp/net.go",
		"test/temp/net_test.go",
	}
	assert.Equal(t, expected, result)
}

func TestExtractPrefixedZip(t *testing.T) {
	prepTmpDir()

	// Copy zip to temp dir
	sys.Copy(testZipfile1, tmpDir)

	// Now extract the files and validate
	zipfile := path.Join(tmpDir, "file1.zip")
	err := ExtractAll(zipfile, tmpDir)
	assert.Nil(t, err)
	paths, err := sys.AllPaths(tmpDir)
	assert.Nil(t, err)
	result := n.S(paths).Map(func(x n.O) n.O { return sys.SlicePath(x.(string), -3, -1) }).ToStrs()
	expected := []string{
		"n/test/temp",
		"test/temp/LICENSE.txt",
		"test/temp/README.md",
		"test/temp/_metadata",
		"temp/_metadata/verified_contents.json",
		"test/temp/contentscript.js",
		"test/temp/file1.zip",
		"test/temp/icon128.png",
		"test/temp/icon16.png",
		"test/temp/icon19.png",
		"test/temp/icon38.png",
		"test/temp/manifest.json",
	}
	assert.Equal(t, expected, result)
}

func TestExtractPrefixedZip2(t *testing.T) {
	prepTmpDir()

	// Copy zip to temp dir
	sys.Copy(testZipfile2, tmpDir)

	// Now extract the files and validate
	zipfile := path.Join(tmpDir, "file2.zip")
	err := ExtractAll(zipfile, tmpDir)
	assert.Nil(t, err)
}

func TestTrimPrefix(t *testing.T) {
	prepTmpDir()
	assert.Nil(t, sys.Copy(testZipfile1, tempZipfile1))

	// empty target name
	{
		assert.Equal(t, "empty string is an invalid path", TrimPrefix("").Error())
	}

	// happy
	{
		// Trim 566 extra bytes at front of zip
		err := TrimPrefix(tempZipfile1)
		assert.Nil(t, err)

		// Now read the zipfile and check the prefix bytes
		data, err := sys.ReadBytes(tempZipfile1)
		assert.Nil(t, err)
		assert.Equal(t, []byte{0x50, 0x4B}, data[0:2])
	}
}

func prepTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}

// OneShotForceZipCloseError patches *archive/zip.Writer.Close to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceZipCloseError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*zip.Writer)(nil)), "Close", func(*zip.Writer) error {
		patch.Unpatch()
		return os.ErrInvalid
	})
}

// OneShotForceZipCreateError patches *archive/zip.Writer.Create to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceZipCreateError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*zip.Writer)(nil)), "Create", func(*zip.Writer, string) (io.Writer, error) {
		patch.Unpatch()
		return nil, os.ErrInvalid
	})
}

// OneShotForceZipFileOpenError patches *archive/zip.File.Open to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceZipFileOpenError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*zip.File)(nil)), "Open", func(*zip.File) (io.ReadCloser, error) {
		patch.Unpatch()
		return nil, os.ErrInvalid
	})
}

// OneShotForceZipOpenReaderError patches *archive/zip.OpenReader to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceZipOpenReaderError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(zip.OpenReader, func(name string) (*zip.ReadCloser, error) {
		patch.Unpatch()
		return nil, os.ErrInvalid
	})
}
