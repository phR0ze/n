package json

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

var tmpDir = "../../../test/temp"
var tmpfile = "../../../test/temp/.tmp"
var testfile = "../../../test/testfile"

func TestReadJSON(t *testing.T) {
	cleanTmpDir()

	// Write out a test json file
	jsondata := `{
  "foo": {
    "bar": [
      1,
      2
    ]
  }
}
`
	fmt.Println(jsondata)
	data1 := map[string]interface{}{}
	err := Unmarshal([]byte(jsondata), &data1)
	assert.Nil(t, err)

	// Write out the data structure as json to disk
	err = WriteJSON(tmpfile, data1)
	assert.Nil(t, err)

	// Read the file back into memory and compare data structure
	var data2 map[string]interface{}
	data2, err = ReadJSON(tmpfile)

	assert.Equal(t, data1, data2)
}

func TestWriteJSON(t *testing.T) {
	cleanTmpDir()

	// Invalid data structure test
	err := WriteJSON(tmpfile, "foo")
	assert.Equal(t, "invalid data structure to marshal - string", err.Error())
	err = WriteJSON(tmpfile, []byte("foo"))
	assert.Equal(t, "invalid data structure to marshal - []uint8", err.Error())

	// Convert json string into a data structure
	jsondata1 := "{\n  \"foo\": {\n    \"bar\": [\n      1,\n      2\n    ]\n  }\n}\n"
	data1 := &map[string]interface{}{}
	err = Unmarshal([]byte(jsondata1), data1)
	assert.Nil(t, err)

	// Write out the data structure as yaml to disk
	err = WriteJSON(tmpfile, data1)
	assert.Nil(t, err)

	// Read the file back into memory and compare data structure
	var jsondata2 []byte
	jsondata2, err = ioutil.ReadFile(tmpfile)
	assert.Nil(t, err)
	data2 := &map[string]interface{}{}
	err = Unmarshal(jsondata2, data2)
	assert.Nil(t, err)

	assert.Equal(t, data1, data2)
}

func cleanTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
