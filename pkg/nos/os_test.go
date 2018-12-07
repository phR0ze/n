package nos

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpFile = "../../test/temp/.tmp"
var readme = "../../README.md"

func TestCopy(t *testing.T) {
	{
		cleanTmpDir()
		Copy("../../pkg", tmpDir)
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

func TestIsDir(t *testing.T) {
	assert.False(t, IsDir(readme))
	assert.True(t, IsDir("../.."))
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile(readme))
	assert.False(t, IsFile("../.."))
}

func TestMD5(t *testing.T) {
	if Exists(tmpFile) {
		os.Remove(tmpFile)
	}
	f, _ := os.Create(tmpFile)
	defer f.Close()
	f.WriteString(`This is a test of the emergency broadcast system.`)

	expected := "067a8c38325b12159844261d16e5cb13"
	result, _ := MD5(tmpFile)
	assert.Equal(t, expected, result)
}

func TestMkdirP(t *testing.T) {
	if Exists(tmpDir) {
		os.RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
	assert.True(t, Exists(tmpDir))
}

func cleanTmpDir() {
	if Exists(tmpDir) {
		os.RemoveAll(tmpDir)
	}
	MkdirP(tmpDir)
}
