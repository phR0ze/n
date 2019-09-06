package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/bouk/monkey"
	"github.com/phR0ze/n/pkg/sys"
	"github.com/phR0ze/n/pkg/test"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"
var testfile = "../../../test/testfile"

func TestCreate(t *testing.T) {
	prepTmpDir()

	// force io.Copy failure
	{
		test.OneShotForceIOCopyError()
		err := Create(tmpfile, testfile)
		assert.True(t, strings.Contains(err.Error(), "failed to copy data from reader to writer for tar target"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// force write header failure
	{
		OneShotForceTarWriterError()
		err := Create(tmpfile, testfile)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to write target file header"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// force file info header error
	{
		OneShotForceTarHeaderError()
		err := Create(tmpfile, testfile)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create target file"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
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
		assert.True(t, strings.Contains(err.Error(), "for tarball: open"))
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
		patch := test.ForceOSCloseError()
		OneShotForceGzipCloseError()
		OneShotForceTarCloseError()
		err = Create(tmpfile, srcDir)
		patch.Unpatch()
		assert.True(t, strings.HasPrefix(err.Error(), "failed to close file writer: failed to close gzip writer: failed to close tarball writer: failed to read directory"))
		assert.True(t, strings.Contains(err.Error(), "to add files from: failed to open directory"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(srcDir, 0755))
		assert.Nil(t, sys.RemoveAll(srcDir))
		assert.Nil(t, sys.Remove(tmpfile))
	}

	// attempt to create tar in read only directory
	{
		// Create new destination directory
		dstDir := path.Join(tmpDir, "dst")
		_, err := sys.MkdirP(dstDir)
		assert.Nil(t, err)

		// Now set the destination to be readonly
		assert.Nil(t, os.Chmod(dstDir, 0444))

		// Now attempt to create the tar in the read only directory
		err = Create(path.Join(dstDir, "tar"), tmpfile)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create tarfile"))
		assert.True(t, strings.Contains(err.Error(), ": open /"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(dstDir, 0755))
		assert.Nil(t, sys.RemoveAll(dstDir))
	}

	// force glob failure
	{
		dstDir := path.Join(tmpDir, "dst")
		_, err := sys.MkdirP(dstDir)
		assert.Nil(t, err)

		test.OneShotForceFilePathGlobError()
		err = Create(path.Join(dstDir, "tar"), tmpfile)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to get glob for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))

		// Correct permission and remove
		assert.Nil(t, sys.RemoveAll(dstDir))
	}

	// no sources found
	{
		err := Create(tmpfile, path.Join(tmpDir, "bogus"))
		assert.True(t, strings.HasPrefix(err.Error(), "failed to get any sources for"))
	}

	// source empty
	{
		err := Create(tmpfile, "")
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// tarfile empty
	{
		err := Create("", "")
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}
}

func TestCreateFromGlobWithExtractCheck(t *testing.T) {
	clearTmpDir()

	// Create target files in dir
	dir1, err := sys.MkdirP(path.Join(tmpDir, "dir1"))
	assert.Nil(t, err)
	file1, err := sys.CopyFile(testfile, path.Join(dir1, "file1"))
	assert.Nil(t, err)
	_, err = sys.CopyFile(testfile, path.Join(dir1, "file2"))
	assert.Nil(t, err)
	bob1, err := sys.CopyFile(testfile, path.Join(dir1, "bob1"))
	assert.Nil(t, err)

	tarball := path.Join(tmpDir, "test.tgz")
	assert.False(t, sys.Exists(tarball))
	err = Create(tarball, path.Join(dir1, "*1"))
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tarball))

	// Now extract the newly created tarball
	dir2 := path.Join(tmpDir, "dir2")
	err = ExtractAll(tarball, dir2)
	assert.Nil(t, err)

	// Now validate the new files
	newfiles, err := sys.AllFiles(dir2)
	assert.Nil(t, err)
	assert.Len(t, newfiles, 2)
	assert.True(t, strings.HasSuffix(newfiles[0], "bob1"))
	assert.True(t, strings.HasSuffix(newfiles[1], "file1"))

	// Now do an MD5 sum of all files
	{
		md5new, err := sys.MD5(newfiles[0])
		assert.Nil(t, err)
		md5old, err := sys.MD5(bob1)
		assert.Nil(t, err)
		assert.Equal(t, md5new, md5old)
	}
	{
		md5new, err := sys.MD5(newfiles[1])
		assert.Nil(t, err)
		md5old, err := sys.MD5(file1)
		assert.Nil(t, err)
		assert.Equal(t, md5new, md5old)
	}
}

func TestCreateFromDirectoryOfFilesWithExtractCheck(t *testing.T) {
	prepTmpDir()

	// Create a tarball to work with
	newsrc := path.Join(tmpDir, "net")
	tarball := path.Join(tmpDir, "test.tgz")
	assert.False(t, sys.Exists(tarball))
	err := Create(tarball, newsrc)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tarball))

	// Rename the original source files to old
	oldsrc := path.Join(tmpDir, "old")
	_, err = sys.Move(newsrc, oldsrc)
	assert.Nil(t, err)

	// Now extract the newly created tarball
	err = ExtractAll(tarball, tmpDir)
	assert.Nil(t, err)

	// Now do a MD5 sum of all files and ensure integrity exists
	newfiles, err := sys.AllFiles(newsrc)
	assert.Nil(t, err)
	oldfiles, err := sys.AllFiles(oldsrc)
	assert.Nil(t, err)
	assert.Equal(t, len(newfiles), len(oldfiles))
	for i := 0; i < len(newfiles); i++ {
		md5new, err := sys.MD5(newfiles[i])
		assert.Nil(t, err)
		md5old, err := sys.MD5(oldfiles[i])
		assert.Nil(t, err)
		assert.Equal(t, md5new, md5old)
	}
}

func TestCreateFromSingleFileWithExtractCheck(t *testing.T) {
	clearTmpDir()

	tarball := path.Join(tmpDir, "test.tgz")
	assert.False(t, sys.Exists(tarball))
	err := Create(tarball, testfile)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tarball))

	// Now extract the newly created tarball
	err = ExtractAll(tarball, tmpDir)
	assert.Nil(t, err)

	// Now validate the new files
	newfiles, err := sys.AllFiles(tmpDir)
	assert.Nil(t, err)
	assert.Len(t, newfiles, 2)
	assert.True(t, strings.HasSuffix(newfiles[0], "test.tgz"))
	assert.True(t, strings.HasSuffix(newfiles[1], "testfile"))

	// Now do an MD5 sum of all files
	md5new, err := sys.MD5(newfiles[1])
	assert.Nil(t, err)
	md5old, err := sys.MD5(testfile)
	assert.Nil(t, err)
	assert.Equal(t, md5new, md5old)
}

func TestExtractAllIntegrityCheckExternal(t *testing.T) {
	prepTmpDir()

	// Create a tarball to work with
	newsrc := path.Join(tmpDir, "net")
	tarball := path.Join(tmpDir, "test.tgz")
	assert.False(t, sys.Exists(tarball))
	sys.ExecOut(`tar -C %s -cvzf %s net`, tmpDir, tarball)
	assert.True(t, sys.Exists(tarball))

	// Rename the original source files to old
	oldsrc := path.Join(tmpDir, "old")
	_, err := sys.Move(newsrc, oldsrc)
	assert.Nil(t, err)

	// Now extract the newly created tarball
	err = ExtractAll(tarball, tmpDir)
	assert.Nil(t, err)

	// Now do a MD5 sum of all files and ensure integrity exists
	newfiles, err := sys.AllFiles(newsrc)
	assert.Nil(t, err)
	oldfiles, err := sys.AllFiles(oldsrc)
	assert.Nil(t, err)
	assert.Equal(t, len(newfiles), len(oldfiles))
	for i := 0; i < len(newfiles); i++ {
		md5new, err := sys.MD5(newfiles[i])
		assert.Nil(t, err)
		md5old, err := sys.MD5(oldfiles[i])
		assert.Nil(t, err)
		assert.Equal(t, md5new, md5old)
	}
}

func TestExtractAll(t *testing.T) {
	prepTmpDir()

	// Create tarball to work with
	src := path.Join(tmpDir, "net")
	err := Create(tmpfile, src)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tmpfile))

	// force os.Chtimes error
	{
		test.OneShotForceOSChtimesError()
		err := ExtractAll(tmpfile, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to set file access times for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// force os.Chmod error
	{
		test.OneShotForceOSChmodError()
		err := ExtractAll(tmpfile, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to set file mode for"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// force os.Close error
	{
		test.OneShotForceOSCloseError()
		err := ExtractAll(tmpfile, tmpDir)
		assert.Equal(t, "failed to close file: invalid argument", err.Error())
	}

	// force os.Copy error
	{
		test.OneShotForceOSCloseError()
		test.OneShotForceIOCopyError()
		err := ExtractAll(tmpfile, tmpDir)
		assert.Equal(t, "failed to close file: failed to copy data from tar to disk: invalid argument", err.Error())
	}

	// force os.Create error
	{
		test.OneShotForceOSCreateError()
		err := ExtractAll(tmpfile, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create file"))
		assert.True(t, strings.HasSuffix(err.Error(), "from tarfile: invalid argument"))
	}

	// force gzip.NewReader error
	{
		OneShotForceGzipReaderError()
		err = ExtractAll(tmpfile, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open gzip reader from"))
		assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	}

	// attempt to read writeonly tarfile
	{
		assert.Nil(t, os.Chmod(tmpfile, 0222))

		err = ExtractAll(tmpfile, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open tarfile"))
		assert.True(t, strings.Contains(err.Error(), "for reading: open"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		assert.Nil(t, os.Chmod(tmpfile, 0644))
	}

	// source empty
	{
		err := ExtractAll(tmpfile, "")
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}

	// tarfile empty
	{
		err := ExtractAll("", "")
		assert.Equal(t, "empty string is an invalid path", err.Error())
	}
}

func clearTmpDir() {
	sys.RemoveAll(tmpDir)
	sys.MkdirP(tmpDir)
}

func prepTmpDir() {
	clearTmpDir()
	sys.Copy("../../net", tmpDir)
}

// OneShotForceGzipCloseError patches *compress/gzip.Writer.Close to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceGzipCloseError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*gzip.Writer)(nil)), "Close", func(*gzip.Writer) error {
		patch.Unpatch()
		return os.ErrInvalid
	})
}

// OneShotForceGzipReaderError patches *compress/gzip.NewReader to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceGzipReaderError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(gzip.NewReader, func(io.Reader) (*gzip.Reader, error) {
		patch.Unpatch()
		return nil, os.ErrInvalid
	})
}

// OneShotForceTarCloseError patches *archive/tar.Writer.Close to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceTarCloseError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*tar.Writer)(nil)), "Close", func(*tar.Writer) error {
		patch.Unpatch()
		return os.ErrInvalid
	})
}

// OneShotForceTarHeaderError patches *archive/tar.FileInfoHeader to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceTarHeaderError() {
	var patch *monkey.PatchGuard
	patch = monkey.Patch(tar.FileInfoHeader, func(os.FileInfo, string) (*tar.Header, error) {
		patch.Unpatch()
		return nil, os.ErrInvalid
	})
}

// OneShotForceTarWriterError patches *archive/tar.Writer.WriteHeader to return an error. Once it has been triggered it
// removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// e.g. go test -gcflags=-l ./pkg/sys
func OneShotForceTarWriterError() {
	var patch *monkey.PatchGuard
	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*tar.Writer)(nil)), "WriteHeader", func(*tar.Writer, *tar.Header) error {
		patch.Unpatch()
		return os.ErrInvalid
	})
}
