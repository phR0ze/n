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
	"github.com/phR0ze/n"
	"github.com/phR0ze/n/pkg/sys"
	"github.com/phR0ze/n/pkg/test"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"

func TestCreate(t *testing.T) {
	prepTmpDir()

	// force tar InfoHeader error
	{
		// Create new source directory/file
		srcDir := path.Join(tmpDir, "src")
		_, err := sys.MkdirP(srcDir)
		assert.Nil(t, err)
		_, err = sys.Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Now attempt to create the tar but force the io.Copy error
		OneShotForceTarHeaderError()
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create target file header"))
		assert.True(t, strings.HasSuffix(err.Error(), "for tarball: invalid argument"))

		// Clean up
		assert.Nil(t, sys.RemoveAll(srcDir))
		assert.Nil(t, sys.Remove(tmpfile))
	}

	// force tar WriteHeader error
	{
		// Create new source directory/file
		srcDir := path.Join(tmpDir, "src")
		_, err := sys.MkdirP(srcDir)
		assert.Nil(t, err)
		_, err = sys.Touch(path.Join(srcDir, "file"))
		assert.Nil(t, err)

		// Now attempt to create the tar but force the io.Copy error
		OneShotForceTarWriterError()
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to write target file header"))
		assert.True(t, strings.HasSuffix(err.Error(), "for tarball: invalid argument"))

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

		// Now attempt to create the tar but force the io.Copy error
		test.OneShotForceIOCopyError()
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to copy data from reader to writer for tar target"))
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
		test.OneShotForceOSCloseError()
		OneShotForceGzipCloseError()
		OneShotForceTarCloseError()
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to close file writer: failed to close gzip writer: failed to close tarball writer: failed to open target file"))
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
		err = Create(tmpfile, srcDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to read directory"))
		assert.True(t, strings.Contains(err.Error(), "to add files from: open"))
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

	// happy
	{
		// Create the new tarball
		src := path.Join(tmpDir, "net")
		err := Create(tmpfile, src)
		assert.Nil(t, err)
		assert.True(t, sys.Exists(tmpfile))

		// Remove tarball target files
		sys.RemoveAll(src)
		assert.False(t, sys.Exists(src))

		// Extract tarball
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

func prepTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
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
