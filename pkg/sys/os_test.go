package sys

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

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
	os.Symlink("../second", symlink)

	// Copy first dir to dst without following links
	{
		beforeInfo, err := Lstat(secondDir)
		assert.Nil(t, err)

		dstDir, _ := Abs(path.Join(tmpDir, "dst"))
		Copy(firstDir, dstDir)

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
	os.Symlink(secondDir, symlink)

	// Copy first dir to dst without following links
	{
		beforeInfo, err := Lstat(secondDir)
		assert.Nil(t, err)

		dstDir, _ := Abs(path.Join(tmpDir, "dst"))
		Copy(firstDir, dstDir)

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

func TestCopyWithFileParentDoentExist(t *testing.T) {
	// test/temp/foo/bar/README.md does not exist and neither does its parent
	// so foo/bar will be created then Copy README.md to bar will be a clone
	cleanTmpDir()
	src := "./README.md"
	dst := path.Join(tmpDir, "foo/bar/readme")

	assert.False(t, Exists(dst))
	Copy(src, dst)
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
	CopyFile(src, dst)
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
	cleanTmpDir()

	// Copy regular file
	foo := path.Join(tmpDir, "foo")

	assert.False(t, Exists(foo))
	CopyFile(readme, foo)
	assert.True(t, Exists(foo))

	srcMD5, err := MD5(readme)
	assert.Nil(t, err)
	dstMD5, err := MD5(foo)
	assert.Nil(t, err)
	assert.Equal(t, srcMD5, dstMD5)

	// Overwrite file
	CopyFile(testfile, foo)
	srcMD5, err = MD5(testfile)
	assert.Nil(t, err)
	dstMD5, err = MD5(foo)
	assert.Nil(t, err)
	assert.Equal(t, srcMD5, dstMD5)
}

func TestExists(t *testing.T) {
	assert.False(t, Exists("bob"))
	assert.True(t, Exists(readme))
}

func TestExtractString(t *testing.T) {
	cleanTmpDir()

	// Write out test data
	err := WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// No match
	match, err := ExtractString(tmpfile, `foobar`)
	assert.Nil(t, err)
	assert.Equal(t, "", match)

	// Single Match - whole string
	match, err = ExtractString(tmpfile, `(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	assert.Equal(t, "pkgver=76.0.3809.100", match)

	// Single Match
	match, err = ExtractString(tmpfile, `(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)

	// Multiple Match - only sees the first one
	match, err = ExtractString(tmpfile, `.*pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)
}

func TestExtractStrings(t *testing.T) {
	cleanTmpDir()

	// Write out test data
	err := WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// No match
	matches, err := ExtractStrings(tmpfile, `foobar`)
	assert.Nil(t, err)
	assert.Nil(t, matches)

	// Single Match substring
	matches, err = ExtractStrings(tmpfile, `(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100"}, matches)

	// Single Match wholestring
	matches, err = ExtractStrings(tmpfile, `(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100"}, matches)

	// Multiple Match substrings
	matches, err = ExtractStrings(tmpfile, `.*pkgver=(.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100", "foo"}, matches)

	// Multiple Match whole strings
	matches, err = ExtractStrings(tmpfile, `.*(pkgver=.*)`)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100", "pkgver=foo"}, matches)
}

func TestExtractStringP(t *testing.T) {
	cleanTmpDir()

	// Write out test data
	err := WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// No match
	rx, err := regexp.Compile(`foobar`)
	assert.Nil(t, err)
	match, err := ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "", match)

	// Single Match - whole string
	rx, err = regexp.Compile(`(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	match, err = ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "pkgver=76.0.3809.100", match)

	// Single Match
	rx, err = regexp.Compile(`(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	match, err = ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)

	// Multiple Match - only sees the first one
	rx, err = regexp.Compile(`.*pkgver=(.*)`)
	assert.Nil(t, err)
	match, err = ExtractStringP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, "76.0.3809.100", match)
}

func TestExtractStringsP(t *testing.T) {
	cleanTmpDir()

	// Write out test data
	err := WriteString(tmpfile, `# test data
pkgname=chromium
pkgver=76.0.3809.100
  pkgver=foo
pkgrel=1
_launcher_ver=6
pkgdesc="Chromium focused on privacy and the removal of Google Orwellian tracking"
arch=('x86_64')
url="https://www.chromium.org/Home"
`)
	assert.Nil(t, err)

	// No match
	rx, err := regexp.Compile(`foobar`)
	assert.Nil(t, err)
	matches, err := ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Nil(t, matches)

	// Single Match substring
	rx, err = regexp.Compile(`(?m)^pkgver=(.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100"}, matches)

	// Single Match wholestring
	rx, err = regexp.Compile(`(?m)^(pkgver=.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100"}, matches)

	// Multiple Match substrings
	rx, err = regexp.Compile(`.*pkgver=(.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"76.0.3809.100", "foo"}, matches)

	// Multiple Match whole strings
	rx, err = regexp.Compile(`.*(pkgver=.*)`)
	assert.Nil(t, err)
	matches, err = ExtractStringsP(tmpfile, rx)
	assert.Nil(t, err)
	assert.Equal(t, []string{"pkgver=76.0.3809.100", "pkgver=foo"}, matches)
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

func TestReadBytes(t *testing.T) {
	cleanTmpDir()

	// Write out test data
	err := WriteString(tmpfile, "this is a test")
	assert.Nil(t, err)

	// Read the file back in and validate
	data, err := ReadBytes(tmpfile)
	assert.Nil(t, err)
	assert.Equal(t, "this is a test", string(data))
}

func TestReadString(t *testing.T) {
	cleanTmpDir()

	// Write out test data
	err := WriteString(tmpfile, "this is a test")
	assert.Nil(t, err)

	// Read the file back in and validate
	data, err := ReadString(tmpfile)
	assert.Nil(t, err)
	assert.Equal(t, "this is a test", data)
}

func TestReadLines(t *testing.T) {
	lines, err := ReadLines(testfile)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(lines))
}

func TestReadYaml(t *testing.T) {
	cleanTmpDir()

	// Write out a test yaml file
	yamldata1 := "foo:\n  bar:\n    - 1\n    - 2\n"
	data1 := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(yamldata1), &data1)
	assert.Nil(t, err)

	// Write out the data structure as yaml to disk
	err = WriteYaml(tmpfile, data1)
	assert.Nil(t, err)

	// Read the file back into memory and compare data structure
	var data2 map[string]interface{}
	data2, err = ReadYaml(tmpfile)

	assert.Equal(t, data1, data2)
}

func TestSize(t *testing.T) {
	assert.Equal(t, int64(604), Size(testfile))

}

func TestTouch(t *testing.T) {
	cleanTmpDir()

	// Doesn't exist so create
	assert.False(t, Exists(tmpfile))
	_, err := Touch(tmpfile)
	assert.Nil(t, err)
	assert.True(t, Exists(tmpfile))

	// Truncate and re-create it
	_, err = Touch(tmpfile)
	assert.Nil(t, err)
}

func TestWriteFile(t *testing.T) {
	cleanTmpDir()

	// Read and write file
	data, err := ioutil.ReadFile(testfile)
	assert.Nil(t, err)
	err = WriteBytes(tmpfile, data)
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

func TestWriteYaml(t *testing.T) {
	cleanTmpDir()

	// Invalid data structure test
	err := WriteYaml(tmpfile, "foo")
	assert.Equal(t, "invalid data structure to marshal - string", err.Error())
	err = WriteYaml(tmpfile, []byte("foo"))
	assert.Equal(t, "invalid data structure to marshal - []uint8", err.Error())

	// Convert yaml string into a data structure
	yamldata1 := "foo:\n  bar:\n    - 1\n    - 2\n"
	data1 := &map[string]interface{}{}
	err = yaml.Unmarshal([]byte(yamldata1), data1)
	assert.Nil(t, err)

	// Write out the data structure as yaml to disk
	err = WriteYaml(tmpfile, data1)
	assert.Nil(t, err)

	// Read the file back into memory and compare data structure
	var yamldata2 []byte
	yamldata2, err = ioutil.ReadFile(tmpfile)
	assert.Nil(t, err)
	data2 := &map[string]interface{}{}
	err = yaml.Unmarshal(yamldata2, data2)
	assert.Nil(t, err)

	assert.Equal(t, data1, data2)
}

func cleanTmpDir() {
	if Exists(tmpDir) {
		RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
}
