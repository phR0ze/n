package n

import (
	"fmt"
	"testing"

	"github.com/phR0ze/n/pkg/sys"
	yaml "github.com/phR0ze/yaml/v2"
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
		assert.Equal(t, M().Add("foo", "bar"), m)
	}

	// More complicated file
	{
		// Write out the yaml to read in
		data := "{\n  \"foo\": {\n    \"value\": 1\n  }\n}"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the json and validate
		m := LoadJSON(tmpFile)
		assert.Equal(t, M().Add("foo", map[string]interface{}{"value": float64(1)}), m)
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
		assert.Equal(t, M().Add("foo", "bar"), m)
	}

	// More complicated file
	{
		// Write out the yaml to read in
		data := "{\n  \"foo\": {\n    \"value\": 1\n  }\n}"
		sys.WriteBytes(tmpFile, []byte(data))

		// Load the json and validate
		m, err := LoadJSONE(tmpFile)
		assert.Nil(t, err)
		assert.Equal(t, M().Add("foo", map[string]interface{}{"value": float64(1)}), m)
	}
}

func TestLoadYAML(t *testing.T) {
	clearTmpDir()

	// Load yaml file
	{
		// Write out the yaml to read in
		data1 := "b: b1\na: a1\n"
		assert.NoError(t, sys.WriteBytes(tmpFile, []byte(data1)))

		// Load yaml and modify target
		m := LoadYAML(tmpFile).Update("a", "a2")
		assert.NoError(t, sys.Remove(tmpFile))
		assert.False(t, sys.Exists(tmpFile))

		// Write YAML out to disk and read back in and assert order
		assert.NoError(t, m.WriteYAML(tmpFile))
		data2, err := sys.ReadBytes(tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "b: b1\na: a2\n", string(data2))
	}
}

func TestLoadYAMLE(t *testing.T) {
	clearTmpDir()

	// Anchors and aliases are handled correctly
	{
		m := map[string]interface{}{}
		data1 := `
foo:
  <<: &foobar1
    bar1: val1
  <<: &foobar2
    bar2: val2
  <<: &foobar3
    bar3: val3
  foo1:
    - <<: *foobar2
      blah1:
        <<: *foobar1
  foo2:
    - <<: *foobar2
      <<: *foobar3
      blah2: true
`
		err := yaml.Unmarshal([]byte(data1), &m)
		assert.NoError(t, err)

		// Validate complete structure
		assert.Equal(t, 1, len(m))
		foo := m["foo"].(map[interface{}]interface{})
		assert.Equal(t, 5, len(foo))
		assert.Equal(t, "val1", foo["bar1"])
		assert.Equal(t, "val2", foo["bar2"])
		assert.Equal(t, "val3", foo["bar3"])

		foo1slice := foo["foo1"].([]interface{})
		assert.Equal(t, 1, len(foo1slice))
		foo1 := foo1slice[0].(map[interface{}]interface{})
		assert.Equal(t, 2, len(foo1))
		assert.Equal(t, "val2", foo1["bar2"])
		assert.Equal(t, "val1", foo1["blah1"].(map[interface{}]interface{})["bar1"])

		foo2slice := foo["foo2"].([]interface{})
		assert.Equal(t, 1, len(foo2slice))
		foo2 := foo2slice[0].(map[interface{}]interface{})
		assert.Equal(t, 3, len(foo2))
		assert.Equal(t, "val2", foo2["bar2"])
		assert.Equal(t, "val3", foo2["bar3"])
		assert.Equal(t, true, foo2["blah2"])

		// Validate complete structure
		fn := func(x *StringMap) {
			assert.Equal(t, 1, x.Len())
			assert.Equal(t, "val1", x.Query("foo.bar1").A())
			assert.Equal(t, "val2", x.Query("foo.bar2").A())
			assert.Equal(t, "val3", x.Query("foo.bar3").A())
			assert.Equal(t, "val2", x.Query("foo.foo1.[0].bar2").A())
			assert.Equal(t, "val1", x.Query("foo.foo1.[0].blah1.bar1").A())
			assert.Equal(t, "val2", x.Query("foo.foo2.[0].bar2").A())
			assert.Equal(t, "val3", x.Query("foo.foo2.[0].bar3").A())
			assert.Equal(t, true, x.Query("foo.foo2.[0].blah2").O())
		}
		fn(MV(m))

		// Now repeat with yaml.StringMap load
		x, err := ToStringMapE(data1)
		assert.NoError(t, err)
		fn(x)
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
