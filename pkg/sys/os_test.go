package sys

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

func TestCopyFollowLinks(t *testing.T) {
	cleanTmpDir()

	// Create first directory
	// temp/first/f0,f1,f2,f3,f4
	firstDir, _ := Abs(path.Join(tmpDir, "first"))
	MkdirP(firstDir)
	for i := 0; i < 5; i++ {
		target := path.Join(firstDir, fmt.Sprintf("f%d", i))
		Touch(target)
	}

	// Create second dir files
	// temp/second/s0,s1,s2,s3,s4
	secondDir, _ := Abs(path.Join(tmpDir, "second"))
	MkdirP(secondDir)
	for i := 0; i < 5; i++ {
		target := path.Join(secondDir, fmt.Sprintf("s%d", i))
		Touch(target)
	}

	// Create sysmlink in first dir to second dir
	symlink := path.Join(tmpDir, "first", "second")
	os.Symlink(secondDir, symlink)

	// Get all the paths following symlinks
	paths, err := AllPaths(firstDir)
	assert.Nil(t, err)

	assert.Equal(t, []string{}, paths)

	// // Compute results using filepath.Walk
	// paths2, err := filepathWalkAllPaths(secondDir)
	// assert.Nil(t, err)

	// // Compare AllFiles to filepathWalkAllFiles
	// assert.Equal(t, paths, paths2)

	// // Now ensure that links are followed
	// paths, err = AllPaths(secondDir)
	// assert.Nil(t, err)
	// assert.Equal(t, expected, paths)
}

func TestCopy(t *testing.T) {
	{
		// test/temp/pkg does not exist
		// so Copy sys to pkg will be a clone
		cleanTmpDir()
		src := "."
		dst := path.Join(tmpDir, "pkg")

		Copy(src, dst)
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

		Copy(src, dst)
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

func TestCopyGlob(t *testing.T) {
	{
		cleanTmpDir()
		dst := path.Join(tmpDir)
		Copy("./*", dst)

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

func TestCopyFile(t *testing.T) {
	{
		cleanTmpDir()
		foo := path.Join(tmpDir, "foo")

		assert.False(t, Exists(foo))
		CopyFile(readme, foo)
		assert.True(t, Exists(foo))

		srcMD5, _ := MD5(readme)
		dstMD5, _ := MD5(foo)
		assert.Equal(t, srcMD5, dstMD5)
	}
}

func TestExists(t *testing.T) {
	assert.False(t, Exists("bob"))
	assert.True(t, Exists(readme))
}

func TestMD5(t *testing.T) {
	if Exists(tmpfile) {
		Remove(tmpfile)
	}
	f, _ := os.Create(tmpfile)
	defer f.Close()
	f.WriteString(`This is a test of the emergency broadcast system.`)

	expected := "067a8c38325b12159844261d16e5cb13"
	result, _ := MD5(tmpfile)
	assert.Equal(t, expected, result)
}

func TestMkdirP(t *testing.T) {
	if Exists(tmpDir) {
		RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
	assert.True(t, Exists(tmpDir))
}

func TestMove(t *testing.T) {

	// Copy file in to tmpDir then rename in same location
	cleanTmpDir()
	Copy(testfile, tmpDir)
	newTestFile := path.Join(tmpDir, "testfile")

	srcMd5, _ := MD5(newTestFile)
	assert.True(t, Exists(newTestFile))
	assert.False(t, Exists(tmpfile))
	err := Move(newTestFile, tmpfile)
	assert.Nil(t, err)
	assert.True(t, Exists(tmpfile))
	dstMd5, _ := MD5(tmpfile)
	assert.False(t, Exists(newTestFile))
	assert.Equal(t, srcMd5, dstMd5)

	// Now create a sub directory and move it there
	subDir := path.Join(tmpDir, "sub")
	MkdirP(subDir)
	err = Move(tmpfile, subDir)
	assert.Nil(t, err)
	assert.False(t, Exists(tmpfile))
	assert.True(t, Exists(path.Join(subDir, path.Base(tmpfile))))
	dstMd5, _ = MD5(path.Join(subDir, path.Base(tmpfile)))
	assert.Equal(t, srcMd5, dstMd5)
}

func TestPwd(t *testing.T) {
	assert.Equal(t, "sys", path.Base(Pwd()))
}

func TestReadLines(t *testing.T) {
	lines, err := ReadLines(testfile)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(lines))
}

func TestSize(t *testing.T) {
	assert.Equal(t, int64(604), Size(testfile))

}

func TestTouch(t *testing.T) {
	cleanTmpDir()
	assert.False(t, Exists(tmpfile))
	assert.Nil(t, Touch(tmpfile))
	assert.True(t, Exists(tmpfile))
	assert.Nil(t, Touch(tmpfile))
}

func TestWriteFile(t *testing.T) {
	cleanTmpDir()

	// Read and write file
	data, err := ioutil.ReadFile(testfile)
	assert.Nil(t, err)
	err = WriteFile(tmpfile, data)
	assert.Nil(t, err)

	// Test the resulting file
	data2, err := ioutil.ReadFile(tmpfile)
	assert.Nil(t, err)
	assert.Equal(t, data, data2)
}

func TestWriteStream(t *testing.T) {
	var expectedData []byte
	expectedData, err := ioutil.ReadFile(testfile)
	assert.Nil(t, err)

	// No file exists
	{
		cleanTmpDir()

		// Read and write file
		reader, err := os.Open(testfile)
		assert.Nil(t, err)
		err = WriteStream(reader, tmpfile)
		assert.Nil(t, err)

		// Test the resulting file
		var data []byte
		data, err = ioutil.ReadFile(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, expectedData, data)
	}

	// Overwrite and truncate file
	{
		// Read and write file
		reader, err := os.Open(testfile)
		assert.Nil(t, err)
		err = WriteStream(reader, tmpfile)
		assert.Nil(t, err)

		// Test the resulting file
		var data []byte
		data, err = ioutil.ReadFile(testfile)
		assert.Nil(t, err)
		assert.Equal(t, expectedData, data)
	}
}

func TestWriteLines(t *testing.T) {
	cleanTmpDir()
	lines, err := ReadLines(testfile)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(lines))
	err = WriteLines(tmpfile, lines)
	assert.Nil(t, err)
	{
		lines2, err := ReadLines(tmpfile)
		assert.Nil(t, err)
		assert.Equal(t, lines, lines2)
	}
}

func cleanTmpDir() {
	if Exists(tmpDir) {
		RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
}
