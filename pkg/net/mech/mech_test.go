package mech

import (
	"fmt"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"
var testfile = "../../../test/testfile"
var readme = "../../../README.md"

func TestDownload(t *testing.T) {
	cleanTmpDir()
	dst, err := Download("https://www.google.com", tmpfile)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(dst))
}

func TestPageLinks(t *testing.T) {
	cleanTmpDir()

	page, err := Get("https://www.google.com")
	assert.Nil(t, err)
	links, err := page.Links()
	assert.NotNil(t, err)
	for _, link := range links {
		fmt.Println(link)
	}
}

func cleanTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
