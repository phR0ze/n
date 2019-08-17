package net

import (
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"
var testfile = "../../test/testfile"
var readme = "../../README.md"

func TestDownloadFile(t *testing.T) {
	cleanTmpDir()
	dst, err := DownloadFile("https://www.google.com", tmpfile)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(dst))
}

func TestDirURL(t *testing.T) {
	assert.Equal(t, "https://foobar.com", DirURL("https://foobar.com/bob"))
	assert.Equal(t, "https://foobar.com/bob/foo", DirURL("https://foobar.com/bob/foo/bar"))
}

func TestJoinURL(t *testing.T) {
	assert.Equal(t, "http://foobar.com/blah/bar", JoinURL("HttP://foobar.com", "/blah", "bar"))
	assert.Equal(t, "http://foobar.com/blah/bar", JoinURL("HttP://foobar.com", "blah", "bar"))
	assert.Equal(t, "https://foobar.com/blah/bar", JoinURL("HttPs://foobar.com", "blah", "bar"))
	assert.Equal(t, "ftp://foobar.com/blah/bar", JoinURL("FTP://foobar.com", "blah", "bar"))
}

func TestNormalizeURL(t *testing.T) {
	assert.Equal(t, "https://foobar", NormalizeURL("Https://foobar"))
}

func TestPing(t *testing.T) {
	err := Ping(TCP, "www.google.com:80")
	assert.Nil(t, err)

	err = Ping(TCP, "blah:80")
	assert.NotNil(t, err)
}

func TestSplitURL(t *testing.T) {
	assert.Equal(t, []string{"http://", "foobar.com", "blah", "bar"}, SplitURL("HttP://foobar.com/blah/bar"))
}

func cleanTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
