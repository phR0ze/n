package ntmpl

import (
	"testing"

	"github.com/phR0ze/n/pkg/nos"
	"github.com/stretchr/testify/assert"
)

var tmpFile = "../../test/temp/.tmp"

func TestLoad(t *testing.T) {
	data := `labels:
  chart: {{ name }}:{{ version }}
  release: {{ Release.Name }}
  heritage: {{ Release.Service }}`
	nos.WriteFile(tmpFile, []byte(data))

	expected := `labels:
  chart: foo:1.0.2
  release: babble
  heritage: fish`

	result, err := Load(tmpFile, "{{ ", " }}", map[string]string{
		"name":            "foo",
		"version":         "1.0.2",
		"Release.Name":    "babble",
		"Release.Service": "fish",
	})
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestTagSpaces(t *testing.T) {
	{
		// README.md example
		data := `labels:
	  chart: {{ name }}:{{ version }}
	  release: {{ Release.Name }}
	  heritage: {{ Release.Service }}`
		expected := `labels:
	  chart: foo:1.0.2
	  release: babble
	  heritage: fish`
		{
			// Tags with spaces
			tpl, err := New(data, "{{ ", " }}")
			assert.Nil(t, err)
			assert.Equal(t, []string{"name", "version", "Release.Name", "Release.Service"}, tpl.tags)
			result, err := tpl.Process(map[string]string{
				"name":            "foo",
				"version":         "1.0.2",
				"Release.Name":    "babble",
				"Release.Service": "fish",
			})
			assert.Nil(t, err)
			assert.Equal(t, expected, result)
		}
		{
			// Tags with no spaces
			tpl, err := New(data, "{{", "}}")
			assert.Equal(t, []string{"name", "version", "Release.Name", "Release.Service"}, tpl.tags)
			assert.NotNil(t, tpl)
			assert.Nil(t, err)
			result, err := tpl.Process(map[string]string{
				"name":            "foo",
				"version":         "1.0.2",
				"Release.Name":    "babble",
				"Release.Service": "fish",
			})
			assert.Nil(t, err)
			assert.Equal(t, expected, result)
		}
	}
}

func TestNoStartTag(t *testing.T) {
	tpl, err := New("foobar", "", "}}")
	assert.Nil(t, tpl)
	assert.NotNil(t, err)
}

func TestNoEndTag(t *testing.T) {
	tpl, err := New("foobar", "{{", "")
	assert.Nil(t, tpl)
	assert.NotNil(t, err)
}

func TestDataHasNoTags(t *testing.T) {
	tpl, err := New("foobar", "{{", "}}")
	assert.Nil(t, tpl)
	assert.NotNil(t, err)
}

func TestEmptyTag(t *testing.T) {
	tpl, err := New("foo{{}}bar", "{{", "}}")
	assert.Nil(t, err)
	result, err := tpl.Process(map[string]string{"": "111", "aaa": "bbb"})
	assert.Nil(t, err)
	assert.Equal(t, "foo111bar", result)
}

func TestSpaceTag(t *testing.T) {
	tpl, err := New("foo{{ }}bar", "{{", "}}")
	assert.Nil(t, err)
	result, err := tpl.Process(map[string]string{"": "111", "aaa": "bbb"})
	assert.Nil(t, err)
	assert.Equal(t, "foo111bar", result)
}

func TestOnlyTag(t *testing.T) {
	tpl, err := New("[foo]", "[", "]")
	assert.Nil(t, err)
	result, err := tpl.Process(map[string]string{"foo": "111", "aaa": "bbb"})
	assert.Nil(t, err)
	assert.Equal(t, "111", result)
}

func TestStartWithTag(t *testing.T) {
	tpl, err := New("[foo]barbaz", "[", "]")
	assert.Nil(t, err)
	result, err := tpl.Process(map[string]string{"foo": "111", "aaa": "bbb"})
	assert.Nil(t, err)
	assert.Equal(t, "111barbaz", result)
}

func TestEndWithTag(t *testing.T) {
	tpl, err := New("foobar[foo]", "[", "]")
	assert.Nil(t, err)
	result, err := tpl.Process(map[string]string{"foo": "111", "aaa": "bbb"})
	assert.Nil(t, err)
	assert.Equal(t, "foobar111", result)
}

// func TestTemplateReset(t *testing.T) {
// 	template := "foo{bar}baz"
// 	tpl := Noob(template, "{", "}")
// 	s := tpl.ExecuteString(map[string]interface{}{"bar": "111"})
// 	result := "foo111baz"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}

// 	template = "[xxxyyyzz"
// 	if err := tpl.Reset(template, "[", "]"); err == nil {
// 		t.Fatalf("expecting error for unclosed tag on %q", template)
// 	}

// 	template = "[xxx]yyy[zz]"
// 	if err := tpl.Reset(template, "[", "]"); err != nil {
// 		t.Fatalf("unexpected error: %s", err)
// 	}
// 	s = tpl.ExecuteString(map[string]interface{}{"xxx": "11", "zz": "2222"})
// 	result = "11yyy2222"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestDuplicateTags(t *testing.T) {
// 	template := "[foo]bar[foo][foo]baz"
// 	tpl := Noob(template, "[", "]")

// 	s := tpl.ExecuteString(map[string]interface{}{"foo": "111", "aaa": "bbb"})
// 	result := "111bar111111baz"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestMultipleTags(t *testing.T) {
// 	template := "foo[foo]aa[aaa]ccc"
// 	tpl := Noob(template, "[", "]")

// 	s := tpl.ExecuteString(map[string]interface{}{"foo": "111", "aaa": "bbb"})
// 	result := "foo111aabbbccc"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestLongDelimiter(t *testing.T) {
// 	template := "foo{{{foo}}}bar"
// 	tpl := Noob(template, "{{{", "}}}")

// 	s := tpl.ExecuteString(map[string]interface{}{"foo": "111", "aaa": "bbb"})
// 	result := "foo111bar"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestIdenticalDelimiter(t *testing.T) {
// 	template := "foo@foo@foo@aaa@"
// 	tpl := Noob(template, "@", "@")

// 	s := tpl.ExecuteString(map[string]interface{}{"foo": "111", "aaa": "bbb"})
// 	result := "foo111foobbb"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestDlimitersWithDistinctSize(t *testing.T) {
// 	template := "foo<?phpaaa?>bar<?phpzzz?>"
// 	tpl := Noob(template, "<?php", "?>")

// 	s := tpl.ExecuteString(map[string]interface{}{"zzz": "111", "aaa": "bbb"})
// 	result := "foobbbbar111"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestEmptyValue(t *testing.T) {
// 	template := "foobar[foo]"
// 	tpl := Noob(template, "[", "]")

// 	s := tpl.ExecuteString(map[string]interface{}{"foo": "", "aaa": "bbb"})
// 	result := "foobar"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestNoValue(t *testing.T) {
// 	template := "foobar[foo]x[aaa]"
// 	tpl := Noob(template, "[", "]")

// 	s := tpl.ExecuteString(map[string]interface{}{"aaa": "bbb"})
// 	result := "foobarxbbb"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

// func TestNoEndDelimiter(t *testing.T) {
// 	template := "foobar[foo"
// 	_, err := New(template, "[", "]")
// 	if err == nil {
// 		t.Fatalf("expected non-nil error. got nil")
// 	}

// 	expectPanic(t, func() { Noob(template, "[", "]") })
// }

// func TestUnsupportedValue(t *testing.T) {
// 	template := "foobar[foo]"
// 	tpl := Noob(template, "[", "]")

// 	expectPanic(t, func() {
// 		tpl.ExecuteString(map[string]interface{}{"foo": 123, "aaa": "bbb"})
// 	})
// }

// func TestMixedValues(t *testing.T) {
// 	template := "foo[foo]bar[bar]baz[baz]"
// 	tpl := Noob(template, "[", "]")

// 	s := tpl.ExecuteString(map[string]interface{}{
// 		"foo": "111",
// 		"bar": []byte("bbb"),
// 		"baz": TagFunc(func(w io.Writer, tag string) (int, error) { return w.Write([]byte(tag)) }),
// 	})
// 	result := "foo111barbbbbazbaz"
// 	if s != result {
// 		t.Fatalf("unexpected template value %q. Expected %q", s, result)
// 	}
// }

func TestProcess(t *testing.T) {
	testProcess(t, "", "", true)
	testProcess(t, "a", "a", true)
	testProcess(t, "abc", "abc", true)
	testProcess(t, "{foo}", "xxxx", false)
	testProcess(t, "a{foo}", "axxxx", false)
	testProcess(t, "{foo}a", "xxxxa", false)
	testProcess(t, "a{foo}bc", "axxxxbc", false)
	testProcess(t, "{foo}{foo}", "xxxxxxxx", false)
	testProcess(t, "{foo}bar{foo}", "xxxxbarxxxx", false)

	// unclosed tag
	testProcess(t, "{unclosed", "{unclosed", true)
	testProcess(t, "{{unclosed", "{{unclosed", true)
	testProcess(t, "{un{closed", "{un{closed", true)

	// test unknown tag (they get removed from the final output)
	testProcess(t, "{unknown}", "", false)
	testProcess(t, "{foo}q{unexpected}{missing}bar{foo}", "xxxxqbarxxxx", false)
}

func testProcess(t *testing.T, template, expected string, shouldErr bool) {
	tpl, err := New(template, "{", "}")
	if shouldErr {
		assert.Nil(t, tpl)
		assert.NotNil(t, err)
	} else {
		assert.Nil(t, err)
		result, err := tpl.Process(map[string]string{"foo": "xxxx"})
		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	}
}
