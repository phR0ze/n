package zip

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var testZipfile1 = "../../../test/file1.zip"
var testZipfile2 = "../../../test/file2.zip"

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"
var tempZipfile1 = "../../../test/temp/file1.zip"
var tempZipfile2 = "../../../test/temp/file2.zip"
var testfile = "../../../test/testfile"

func TestCreateSad(t *testing.T) {
	clearTmpDir()

	// // force zip.Create error
	// {
	// 	// Create new source directory/file
	// 	srcDir := path.Join(tmpDir, "src")
	// 	_, err := sys.MkdirP(srcDir)
	// 	assert.Nil(t, err)
	// 	_, err = sys.Touch(path.Join(srcDir, "file"))
	// 	assert.Nil(t, err)

	// 	// Now attempt to create th zip but force the io.Copy error
	// 	OneShotForceZipCreateError()
	// 	err = Create(tmpfile, srcDir)
	// 	assert.True(t, strings.HasPrefix(err.Error(), "failed to add target file"))
	// 	assert.True(t, strings.HasSuffix(err.Error(), "to zip: invalid argument"))

	// 	// Clean up
	// 	assert.Nil(t, sys.RemoveAll(srcDir))
	// 	assert.Nil(t, sys.Remove(tmpfile))
	// }

	// // force io.Copy error
	// {
	// 	// Create new source directory/file
	// 	srcDir := path.Join(tmpDir, "src")
	// 	_, err := sys.MkdirP(srcDir)
	// 	assert.Nil(t, err)
	// 	_, err = sys.Touch(path.Join(srcDir, "file"))
	// 	assert.Nil(t, err)

	// 	// Now attempt to create th zip but force the io.Copy error
	// 	test.OneShotForceIOCopyError()
	// 	err = Create(tmpfile, srcDir)
	// 	assert.True(t, strings.HasPrefix(err.Error(), "failed to copy data from reader to writer for zip target"))
	// 	assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))

	// 	// Clean up
	// 	assert.Nil(t, sys.RemoveAll(srcDir))
	// 	assert.Nil(t, sys.Remove(tmpfile))
	// }

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
		assert.True(t, strings.Contains(err.Error(), "to add files from: failed to open directory"))
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

	// // force glob failure
	// {
	// 	dstDir := path.Join(tmpDir, "dst")
	// 	_, err := sys.MkdirP(dstDir)
	// 	assert.Nil(t, err)

	// 	test.OneShotForceOSCloseError()
	// 	test.OneShotForceFilePathGlobError()
	// 	OneShotForceZipCloseError()
	// 	err = Create(path.Join(dstDir, "tar"), tmpfile)
	// 	assert.True(t, strings.HasPrefix(err.Error(), "failed to close file writer: failed to close zipfile writer: failed to get glob for"))
	// 	assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))

	// 	// Correct permission and remove
	// 	assert.Nil(t, sys.RemoveAll(dstDir))
	// }

	// no sources found
	{
		err := Create(tmpfile, path.Join(tmpDir, "bogus"))
		assert.True(t, strings.HasPrefix(err.Error(), "failed to get any sources for"))
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

	zipfile := path.Join(tmpDir, "test.zip")
	assert.False(t, sys.Exists(zipfile))
	err = Create(zipfile, path.Join(dir1, "*1"))
	assert.Nil(t, err)
	assert.True(t, sys.Exists(zipfile))

	// Now extract the newly created tarball
	dir2 := path.Join(tmpDir, "dir2")
	err = ExtractAll(zipfile, dir2)
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

	// Create a zipfile to work with
	newsrc := path.Join(tmpDir, "net")
	zipfile := path.Join(tmpDir, "test.zip")
	assert.False(t, sys.Exists(zipfile))
	err := Create(zipfile, newsrc)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(zipfile))

	// Rename the original source files to old
	oldsrc := path.Join(tmpDir, "old")
	_, err = sys.Move(newsrc, oldsrc)
	assert.Nil(t, err)

	// Now extract the newly created zipfile
	err = ExtractAll(zipfile, tmpDir)
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

	zipfile := path.Join(tmpDir, "test.zip")
	assert.False(t, sys.Exists(zipfile))
	err := Create(zipfile, testfile)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(zipfile))

	// Now extract the newly created zipfile
	err = ExtractAll(zipfile, tmpDir)
	assert.Nil(t, err)

	// Now validate the new files
	newfiles, err := sys.AllFiles(tmpDir)
	assert.Nil(t, err)
	assert.Len(t, newfiles, 2)
	assert.True(t, strings.HasSuffix(newfiles[0], "test.zip"))
	assert.True(t, strings.HasSuffix(newfiles[1], "testfile"))

	// Now do an MD5 sum of all files
	md5new, err := sys.MD5(newfiles[1])
	assert.Nil(t, err)
	md5old, err := sys.MD5(testfile)
	assert.Nil(t, err)
	assert.Equal(t, md5new, md5old)
}

func TestExtractAllSad(t *testing.T) {
	clearTmpDir()

	// // force os.Chtimes error
	// {
	// 	assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
	// 	test.OneShotForceOSChtimesError()
	// 	err := ExtractAll(tempZipfile1, tmpDir)
	// 	assert.True(t, strings.HasPrefix(err.Error(), "failed to set file access times for"))
	// 	assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	// }

	// // force os.Chmod error
	// {
	// 	assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
	// 	test.OneShotForceOSChmodError()
	// 	err := ExtractAll(tempZipfile1, tmpDir)
	// 	assert.True(t, strings.HasPrefix(err.Error(), "failed to set file mode for"))
	// 	assert.True(t, strings.HasSuffix(err.Error(), ": invalid argument"))
	// }

	// // force Close error
	// {
	// 	// Pretrim and skip trim
	// 	assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
	// 	TrimPrefix(tempZipfile1)
	// 	ForceNilFromTrimPrefix()

	// 	test.OneShotForceOSCloseError()
	// 	err := ExtractAll(tempZipfile1, tmpDir)
	// 	assert.Equal(t, "failed to close zipfile writer: invalid argument", err.Error())
	// }

	// // force os.Copy error
	// {
	// 	// Pretrim and skip trim
	// 	assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
	// 	TrimPrefix(tempZipfile1)
	// 	ForceNilFromTrimPrefix()

	// 	test.OneShotForceIOCopyError()
	// 	test.OneShotForceOSCloseError()
	// 	err := ExtractAll(tempZipfile1, tmpDir)
	// 	assert.Equal(t, "failed to close zipfile writer: failed to copy data from zip to disk: invalid argument", err.Error())
	// }

	// // force zip.FileOpen error
	// {
	// 	// Pretrim and skip trim
	// 	assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
	// 	TrimPrefix(tempZipfile1)
	// 	ForceNilFromTrimPrefix()

	// 	OneShotForceZipFileOpenError()
	// 	test.OneShotForceOSCloseError()
	// 	err := ExtractAll(tempZipfile1, tmpDir)
	// 	assert.True(t, strings.HasPrefix(err.Error(), "failed to close zipfile writer: failed to open zipfile target"))
	// 	assert.True(t, strings.HasSuffix(err.Error(), "for reading: invalid argument"))
	// }

	// attempt to write to writeonly destination path
	{
		assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
		// Create write only destination
		dir, err := sys.MkdirP(path.Join(tmpDir, "dir"))
		assert.Nil(t, err)
		assert.Nil(t, os.Chmod(dir, 0222))

		// Now attempt to read the write only zip
		err = ExtractAll(tempZipfile1, dir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to create file"))
		assert.True(t, strings.Contains(err.Error(), "from zipfile: open"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))

		// Correct permission and remove
		assert.Nil(t, os.Chmod(dir, 0755))
		assert.Nil(t, sys.RemoveAll(dir))
	}

	// force zip.OpenReader error
	// {
	// 	assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
	// 	OneShotForceZipOpenReaderError()
	// 	err := ExtractAll(tempZipfile1, tmpDir)
	// 	assert.True(t, strings.HasPrefix(err.Error(), "failed to open zipfile"))
	// 	assert.True(t, strings.HasSuffix(err.Error(), "for reading: invalid argument"))
	// }

	// invalid zipfile
	{
		assert.Nil(t, sys.Copy(testZipfile1, tmpDir))
		err := ExtractAll(tmpfile, tmpDir)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open zipfile '.tmp' to detect zip data"))
		assert.True(t, strings.HasSuffix(err.Error(), ": no such file or directory"))
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

func TestExtractPrefixedZip(t *testing.T) {
	clearTmpDir()

	// Copy zip to temp dir
	sys.Copy(testZipfile1, tmpDir)

	// Now extract the files and validate
	zipfile := path.Join(tmpDir, "file1.zip")
	err := ExtractAll(zipfile, tmpDir)
	assert.Nil(t, err)
	paths, err := sys.AllPaths(tmpDir)
	assert.Nil(t, err)

	result := []string{}
	for _, path := range paths {
		result = append(result, sys.SlicePath(path, -3, -1))
	}

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
	clearTmpDir()

	// Copy zip to temp dir
	sys.Copy(testZipfile2, tmpDir)

	// Now extract the files and validate
	zipfile := path.Join(tmpDir, "file2.zip")
	err := ExtractAll(zipfile, tmpDir)
	assert.Nil(t, err)
}

func TestTrimPrefix(t *testing.T) {
	clearTmpDir()
	assert.Nil(t, sys.Copy(testZipfile1, tempZipfile1))

	// happy
	{
		// Trim 566 extra bytes at front of zip
		err := TrimPrefix(tempZipfile1)
		assert.Nil(t, err)

		// Now read the zipfile and check the prefix bytes
		data, err := sys.ReadBytes(tempZipfile1)
		assert.Nil(t, err)
		assert.Equal(t, []byte{0x50, 0x4B}, data[0:2])
		assert.Nil(t, sys.Copy(testZipfile1, tempZipfile1))
	}

	// // force os.File.Truncate failure and fail close
	// {
	// 	test.OneShotForceOSCloseError()
	// 	test.OneShotForceOSTruncateError(os.ErrInvalid)
	// 	err := TrimPrefix(tempZipfile1)
	// 	assert.Equal(t, "failed to close file writer: failed to truncate zipfile 'file1.zip': invalid argument", err.Error())
	// 	assert.Nil(t, sys.Copy(testZipfile1, tempZipfile1))
	// }

	// // force os.File.WriteAt failure
	// {
	// 	test.OneShotForceOSWriteAtError(os.ErrInvalid)
	// 	err := TrimPrefix(tempZipfile1)
	// 	assert.Equal(t, "failed to write shifted data to zipfile 'file1.zip': invalid argument", err.Error())
	// 	assert.Nil(t, sys.Copy(testZipfile1, tempZipfile1))
	// }

	// // force os.File.ReadAt failure
	// {
	// 	test.OneShotForceOSReadAtError(os.ErrInvalid)
	// 	err := TrimPrefix(tempZipfile1)
	// 	assert.Equal(t, "failed to read from zipfile 'file1.zip' to shift data: invalid argument", err.Error())
	// }

	// // force os.File.Read failure io.EOF
	// {
	// 	test.OneShotForceOSReadError(io.EOF)
	// 	err := TrimPrefix(tempZipfile1)
	// 	assert.Equal(t, "unable to identify 'file1.zip' as a valid zipfile", err.Error())
	// }

	// // force io.Read failure non io.EOF
	// {
	// 	test.OneShotForceOSReadError(os.ErrInvalid)
	// 	err := TrimPrefix(tempZipfile1)
	// 	assert.Equal(t, "failed to read from zipfile 'file1.zip' for zip identification: invalid argument", err.Error())
	// }

	// try to open a write only file
	{
		assert.Nil(t, os.Chmod(tempZipfile1, 0222))
		err := TrimPrefix(tempZipfile1)
		assert.True(t, strings.HasPrefix(err.Error(), "failed to open zipfile 'file1.zip'"))
		assert.True(t, strings.HasSuffix(err.Error(), ": permission denied"))
		assert.Nil(t, os.Chmod(tempZipfile1, 0644))
	}

	// empty target name
	{
		assert.Equal(t, "empty string is an invalid path", TrimPrefix("").Error())
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

// // OneShotForceZipCloseError patches *archive/zip.Writer.Close to return an error. Once it has been triggered it
// // removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// // e.g. go test -gcflags=-l ./pkg/sys
// func OneShotForceZipCloseError() {
// 	var patch *monkey.PatchGuard
// 	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*zip.Writer)(nil)), "Close", func(*zip.Writer) error {
// 		patch.Unpatch()
// 		return os.ErrInvalid
// 	})
// }

// // OneShotForceZipCreateError patches *archive/zip.Writer.Create to return an error. Once it has been triggered it
// // removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// // e.g. go test -gcflags=-l ./pkg/sys
// func OneShotForceZipCreateError() {
// 	var patch *monkey.PatchGuard
// 	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*zip.Writer)(nil)), "Create", func(*zip.Writer, string) (io.Writer, error) {
// 		patch.Unpatch()
// 		return nil, os.ErrInvalid
// 	})
// }

// // OneShotForceZipFileOpenError patches *archive/zip.File.Open to return an error. Once it has been triggered it
// // removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// // e.g. go test -gcflags=-l ./pkg/sys
// func OneShotForceZipFileOpenError() {
// 	var patch *monkey.PatchGuard
// 	patch = monkey.PatchInstanceMethod(reflect.TypeOf((*zip.File)(nil)), "Open", func(*zip.File) (io.ReadCloser, error) {
// 		patch.Unpatch()
// 		return nil, os.ErrInvalid
// 	})
// }

// // OneShotForceZipOpenReaderError patches *archive/zip.OpenReader to return an error. Once it has been triggered it
// // removes the patch and operates as per normal. This patch requires the -gcflags=-l to operate correctly
// // e.g. go test -gcflags=-l ./pkg/sys
// func OneShotForceZipOpenReaderError() {
// 	var patch *monkey.PatchGuard
// 	patch = monkey.Patch(zip.OpenReader, func(name string) (*zip.ReadCloser, error) {
// 		patch.Unpatch()
// 		return nil, os.ErrInvalid
// 	})
// }

// func ForceNilFromTrimPrefix() {
// 	var patch *monkey.PatchGuard
// 	patch = monkey.Patch(TrimPrefix, func(string) error {
// 		patch.Unpatch()
// 		return nil
// 	})
// }
