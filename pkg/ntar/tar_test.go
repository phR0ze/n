package ntar

import (
	"path"
	"testing"

	"github.com/phR0ze/n/pkg/nos"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"

func TestCreate(t *testing.T) {
	prepTmpDir()

	// Create the new tarball
	src := path.Join(tmpDir, "ncli")
	err := Create(tmpfile, src)
	assert.Nil(t, err)
	assert.True(t, nos.Exists(tmpfile))

	// Remove tarball target files
	nos.RemoveAll(src)
	assert.False(t, nos.Exists(src))

	// Extract tarball
	err = ExtractAll(tmpfile, tmpDir)
	assert.True(t, nos.Exists(path.Join(src, "cli.go")))
	assert.True(t, nos.Exists(path.Join(src, "cli_test.go")))
}

func prepTmpDir() {
	if nos.Exists(tmpDir) {
		nos.RemoveAll(tmpDir)
	}
	nos.MkdirP(tmpDir)
	nos.Copy("../ncli", tmpDir)
}
