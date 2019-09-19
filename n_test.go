package n

import (
	"fmt"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	"github.com/stretchr/testify/assert"
)

const nines7 = 9999999
const nines6 = 999999
const nines5 = 99999
const nines4 = 9999
const nines3 = 999

type bob struct {
	o interface{}
}

var tmpDir = "test/temp"
var tmpFile = "test/temp/.tmp"

func ExampleEitherStr() {
	fmt.Println(EitherStr("", "default"))
	// Output: default
}

func TestEitherStr(t *testing.T) {
	// Test default set
	{
		assert.Equal(t, "test", EitherStr("", "test"))
	}

	// Test value set
	{
		assert.Equal(t, "foo", EitherStr("foo", "test"))
	}
}

func ExampleSetOnEmpty() {
	result := ""
	fmt.Println(SetOnEmpty(&result, "foo"))
	// Output: foo
}

func TestSetOnEmpty(t *testing.T) {
	result := ""
	assert.Equal(t, "foo", SetOnEmpty(&result, "foo"))
	assert.Equal(t, "foo", SetOnEmpty(&result, "bar"))
	assert.Equal(t, "bar", SetOnEmpty(nil, "bar"))
}

func TestSetOnTrueA(t *testing.T) {
	result := ""
	assert.Equal(t, "foo", SetOnTrueA(&result, "foo", true))
	assert.Equal(t, "foo", SetOnTrueA(&result, "bar", true))
	assert.Equal(t, "bar", SetOnTrueA(nil, "bar", true))
	assert.Equal(t, "foo", SetOnTrueA(&result, "bar", false))
}

func ExampleSetOnFalseB() {
	result := false
	fmt.Println(SetOnFalseB(&result, true, false))
	// Output: true
}

func TestSetOnFalseB(t *testing.T) {
	result := false
	assert.Equal(t, true, SetOnFalseB(&result, true, false))
	assert.Equal(t, true, SetOnFalseB(&result, false, true))
	assert.Equal(t, true, SetOnFalseB(&result, true, false))
}

func ExampleSetOnTrueB() {
	result := false
	fmt.Println(SetOnTrueB(&result, true, true))
	// Output: true
}

func TestSetOnTrueB(t *testing.T) {
	result := false
	assert.Equal(t, true, SetOnTrueB(&result, true, true))
	assert.Equal(t, false, SetOnTrueB(&result, false, true))
	assert.Equal(t, false, SetOnTrueB(&result, true, false))
}

func TestRange(t *testing.T) {
	assert.Equal(t, []int{0}, Range(0, 0))
	assert.Equal(t, []int{0, 1}, Range(0, 1))
	assert.Equal(t, []int{3, 4, 5, 6, 7, 8}, Range(3, 8))
}

func TestLoadJSON(t *testing.T) {
	clearTmpDir()

	// Load simple file
	{
		// Write out the yaml to read in
		data := "{\n  \"foo\": \"bar\"\n}"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the json and validate
		m := LoadJSON(tmpFile)
		assert.Equal(t, &StringMap{"foo": "bar"}, m)
	}

	// More complicated file
	{
		// Write out the yaml to read in
		data := "{\n  \"foo\": {\n    \"value\": 1\n  }\n}"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the json and validate
		m := LoadJSON(tmpFile)
		assert.Equal(t, &StringMap{"foo": map[string]interface{}{"value": float64(1)}}, m)
	}
}

func TestLoadJSONE(t *testing.T) {
	clearTmpDir()

	// Load simple file
	{
		// Write out the yaml to read in
		data := "{\n  \"foo\": \"bar\"\n}"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the json and validate
		m, err := LoadJSONE(tmpFile)
		assert.Nil(t, err)
		assert.Equal(t, &StringMap{"foo": "bar"}, m)
	}

	// More complicated file
	{
		// Write out the yaml to read in
		data := "{\n  \"foo\": {\n    \"value\": 1\n  }\n}"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the json and validate
		m, err := LoadJSONE(tmpFile)
		assert.Nil(t, err)
		assert.Equal(t, &StringMap{"foo": map[string]interface{}{"value": float64(1)}}, m)
	}
}

func TestLoadYAML(t *testing.T) {
	clearTmpDir()

	// Load yaml file
	{
		// Write out the yaml to read in
		data := "foo:\n  bar: 1\n"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the yaml and validate
		m := LoadYAML(tmpFile)
		assert.Equal(t, &StringMap{"foo": map[string]interface{}{"bar": float64(1)}}, m)
	}
}

func TestLoadYAMLE(t *testing.T) {
	clearTmpDir()

	// Load yaml file
	{
		// Write out the yaml to read in
		data := "foo:\n  bar: 1\n"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the yaml and validate
		m, err := LoadYAMLE(tmpFile)
		assert.Nil(t, err)
		assert.Equal(t, &StringMap{"foo": map[string]interface{}{"bar": float64(1)}}, m)
	}
}

func rangeObject(min, max int) []Object {
	result := make([]Object, max-min+1)
	for i := range result {
		result[i] = Object{min + i}
	}
	return result
}

func rangeInterObject(min, max int) []interface{} {
	result := make([]interface{}, max-min+1)
	for i := range result {
		result[i] = Object{min + i}
	}
	return result
}

// rangeO creates slice of the given range of numbers inclusive
func rangeO(min, max int) []interface{} {
	result := make([]interface{}, max-min+1)
	for i := range result {
		result[i] = min + i
	}
	return result
}

func clearTmpDir() {
	if sys.Exists(tmpDir) {
		sys.RemoveAll(tmpDir)
	}
	sys.MkdirP(tmpDir)
}
