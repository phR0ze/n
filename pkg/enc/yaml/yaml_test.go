package yaml

import (
	"io/ioutil"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"
var testfile = "../../../test/testfile"

func TestReadYAML(t *testing.T) {
	clearTmpDir()

	// Write out a test yaml file
	yamldata1 := "foo:\n  bar:\n    - 1\n    - 2\n"
	data1 := map[string]interface{}{}
	err := Unmarshal([]byte(yamldata1), &data1)
	assert.Nil(t, err)

	// Write out the data structure as yaml to disk
	err = WriteYAML(tmpfile, data1)
	assert.Nil(t, err)

	// Read the file back into memory and compare data structure
	var data2 map[string]interface{}
	data2, err = ReadYAML(tmpfile)

	assert.Equal(t, data1, data2)
}

func TestWriteYAML(t *testing.T) {
	clearTmpDir()

	// Invalid data structure test
	err := WriteYAML(tmpfile, "foo")
	assert.Equal(t, "invalid data structure to marshal - string", err.Error())
	err = WriteYAML(tmpfile, []byte("foo"))
	assert.Equal(t, "invalid data structure to marshal - []uint8", err.Error())

	// Convert yaml string into a data structure
	yamldata1 := "foo:\n  bar:\n    - 1\n    - 2\n"
	data1 := &map[string]interface{}{}
	err = Unmarshal([]byte(yamldata1), data1)
	assert.Nil(t, err)

	// Write out the data structure as yaml to disk
	err = WriteYAML(tmpfile, data1)
	assert.Nil(t, err)

	// Read the file back into memory and compare data structure
	var yamldata2 []byte
	yamldata2, err = ioutil.ReadFile(tmpfile)
	assert.Nil(t, err)
	data2 := &map[string]interface{}{}
	err = Unmarshal(yamldata2, data2)
	assert.Nil(t, err)

	assert.Equal(t, data1, data2)
}

func clearTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
