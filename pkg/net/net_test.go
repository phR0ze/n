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

func TestPing(t *testing.T) {
	err := Ping(TCP, "www.google.com:80")
	assert.Nil(t, err)

	err = Ping(TCP, "blah:80")
	assert.NotNil(t, err)
}

func cleanTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
