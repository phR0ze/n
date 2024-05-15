package yaml

import (
	"os"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	yaml "github.com/phR0ze/yaml/v2"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"

func TestReadYAML(t *testing.T) {
	clearTmpDir()

	// Write out a test yaml file
	yamldata1 := "foo:\n  bar:\n    - 1\n    - 2\n"
	data1 := map[string]interface{}{}
	err := Unmarshal([]byte(yamldata1), &data1)
	assert.NoError(t, err)

	// Write out the data structure as yaml to disk
	err = WriteYAML(tmpfile, data1)
	assert.NoError(t, err)

	// Read the file back into memory and compare data structure
	var data2 map[string]interface{}
	data2, err = ReadYAML(tmpfile)
	assert.NoError(t, err)

	assert.Equal(t, data1, data2)
}

func TestWriteYAML(t *testing.T) {
	clearTmpDir()

	// Invalid data structure test
	{
		err := WriteYAML(tmpfile, "foo")
		assert.Equal(t, "invalid data structure to marshal - string", err.Error())
		err = WriteYAML(tmpfile, []byte("foo"))
		assert.Equal(t, "invalid data structure to marshal - []uint8", err.Error())
	}

	// Convert yaml string into a data structure
	{
		yamldata1 := "foo:\n  bar:\n    - 1\n    - 2\n"
		data1 := &map[string]interface{}{}
		err := Unmarshal([]byte(yamldata1), data1)
		assert.NoError(t, err)

		// Write out the data structure as yaml to disk
		err = WriteYAML(tmpfile, data1)
		assert.NoError(t, err)

		// Read the file back into memory and compare data structure
		var yamldata2 []byte
		yamldata2, err = os.ReadFile(tmpfile)
		assert.NoError(t, err)
		data2 := &map[string]interface{}{}
		err = Unmarshal(yamldata2, data2)
		assert.NoError(t, err)

		assert.Equal(t, data1, data2)
	}
}

func TestOrderedReadWriteYaml(t *testing.T) {
	clearTmpDir()

	// Write out the data structure as yaml to disk
	data1 := "b: b1\na: a1\n"
	yamldata1 := yaml.MapSlice{}
	err := Unmarshal([]byte(data1), &yamldata1)
	assert.NoError(t, err)
	err = WriteYAML(tmpfile, yamldata1)
	assert.NoError(t, err)

	// Read the file back into memory and compare raw string
	var buffer []byte
	buffer, err = os.ReadFile(tmpfile)
	assert.Nil(t, err)
	data2 := string(buffer)
	assert.Equal(t, data1, data2)

	// Now compare data structures
	yamldata2 := yaml.MapSlice{}
	err = Unmarshal(buffer, &yamldata2)
	assert.NoError(t, err)
	assert.Equal(t, yamldata1, yamldata2)

	// Now convert back and check again
	data3, err := Marshal(yamldata2)
	assert.NoError(t, err)
	assert.Equal(t, data1, string(data3))
}

func clearTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
