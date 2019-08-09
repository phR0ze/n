package tar

import (
	"path"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../test/temp"
var tmpfile = "../../test/temp/.tmp"

func TestCreate(t *testing.T) {
	prepTmpDir()

	// Create the new tarball
	src := path.Join(tmpDir, "cli")
	err := Create(src, tmpfile)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tmpfile))

	// Remove tarball target files
	sys.RemoveAll(src)
	assert.False(t, sys.Exists(src))

	// Extract tarball
	err = ExtractAll(tmpfile, tmpDir)
	assert.True(t, sys.Exists(path.Join(src, "cli.go")))
	assert.True(t, sys.Exists(path.Join(src, "cli_test.go")))
}

func prepTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
	sys.Copy("../cli", tmpDir)
}
