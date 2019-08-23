package zip

import (
	"path"
	"testing"

	"github.com/phR0ze/n"
	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"

func TestCreate(t *testing.T) {
	prepTmpDir()

	// Create the new zipfile
	src := path.Join(tmpDir, "net")
	err := Create(tmpfile, src)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tmpfile))

	// Remove zip target files
	sys.RemoveAll(src)
	assert.False(t, sys.Exists(src))

	// Call out and extract the files
	_, err = sys.ExecOut("unzip %s -d %s", tmpfile, tmpDir)
	assert.Nil(t, err)
	paths, err := sys.AllPaths(tmpDir)
	assert.Nil(t, err)
	result := n.S(paths).Map(func(x n.O) n.O { return sys.SlicePath(x.(string), -3, -1) }).ToStrs()
	expected := []string{
		"n/test/temp",
		"test/temp/.tmp",
		"test/temp/agent",
		"temp/agent/agent.go",
		"test/temp/mech",
		"temp/mech/example",
		"mech/example/mech.go",
		"temp/mech/mech.go",
		"temp/mech/mech_test.go",
		"temp/mech/page.go",
		"test/temp/net.go",
		"test/temp/net_test.go",
	}
	assert.Equal(t, expected, result)
}

func TestExtractAll(t *testing.T) {
	prepTmpDir()

	// Create the new zipfile
	src := path.Join(tmpDir, "net")
	err := Create(tmpfile, src)
	assert.Nil(t, err)
	assert.True(t, sys.Exists(tmpfile))

	// Remove zip target files
	sys.RemoveAll(src)
	assert.False(t, sys.Exists(src))

	// Now extract the files and validate
	err = ExtractAll(tmpfile, tmpDir)
	assert.Nil(t, err)
	paths, err := sys.AllPaths(tmpDir)
	assert.Nil(t, err)
	result := n.S(paths).Map(func(x n.O) n.O { return sys.SlicePath(x.(string), -3, -1) }).ToStrs()
	expected := []string{
		"n/test/temp",
		"test/temp/.tmp",
		"test/temp/agent",
		"temp/agent/agent.go",
		"test/temp/mech",
		"temp/mech/example",
		"mech/example/mech.go",
		"temp/mech/mech.go",
		"temp/mech/mech_test.go",
		"temp/mech/page.go",
		"test/temp/net.go",
		"test/temp/net_test.go",
	}
	assert.Equal(t, expected, result)
}

func prepTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
	sys.Copy("../../net", tmpDir)
}
