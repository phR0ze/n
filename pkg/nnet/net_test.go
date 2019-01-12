package nnet

import (
	"testing"

	"github.com/phR0ze/n/pkg/nos"
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
	assert.True(t, nos.Exists(dst))
}

func cleanTmpDir() {
	if nos.Exists(tmpDir) {
		nos.RemoveAll(tmpDir)
	}
	nos.MkdirP(tmpDir)
}
