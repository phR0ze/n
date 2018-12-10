package n

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestStrA(t *testing.T) {
	assert.Equal(t, "test", A("test").A())
}

func TestStrQ(t *testing.T) {
	assert.Equal(t, "test", A("test").Q().A())
}

func TestStrContains(t *testing.T) {
	assert.True(t, A("test").Contains("tes"))
	assert.False(t, A("test").Contains("bob"))
}

func TestStrContainsAny(t *testing.T) {
	assert.True(t, A("test").ContainsAny("tes"))
	assert.True(t, A("test").ContainsAny("f", "t"))
	assert.False(t, A("test").ContainsAny("f", "b"))
}

func TestStrHasAnyPrefix(t *testing.T) {
	assert.True(t, A("test").HasAnyPrefix("tes"))
	assert.True(t, A("test").HasAnyPrefix("bob", "tes"))
	assert.False(t, A("test").HasAnyPrefix("bob"))
}

func TestStrHasAnySuffix(t *testing.T) {
	assert.True(t, A("test").HasAnySuffix("est"))
	assert.True(t, A("test").HasAnySuffix("bob", "est"))
	assert.False(t, A("test").HasAnySuffix("bob"))
}

func TestStrHasPrefix(t *testing.T) {
	assert.True(t, A("test").HasPrefix("tes"))
}

func TestStrHasSuffix(t *testing.T) {
	assert.True(t, A("test").HasSuffix("est"))
}

func TestStrSplit(t *testing.T) {
	assert.Equal(t, []string{"1", "2"}, A("1.2").Split(".").S())
}

func TestStrSplitOn(t *testing.T) {
	{
		first, second := A("").SplitOn(":")
		assert.Equal(t, "", first)
		assert.Equal(t, "", second)
	}
	{
		first, second := A("foo").SplitOn(":")
		assert.Equal(t, "foo", first)
		assert.Equal(t, "", second)
	}
	{
		first, second := A("foo:").SplitOn(":")
		assert.Equal(t, "foo:", first)
		assert.Equal(t, "", second)
	}
	{
		first, second := A(":foo").SplitOn(":")
		assert.Equal(t, ":", first)
		assert.Equal(t, "foo", second)
	}
	{
		first, second := A("foo: bar").SplitOn(":")
		assert.Equal(t, "foo:", first)
		assert.Equal(t, " bar", second)
	}
	{
		first, second := A("foo: bar:frodo").SplitOn(":")
		assert.Equal(t, "foo:", first)
		assert.Equal(t, " bar:frodo", second)
	}
}

func TestStrSpaceLeft(t *testing.T) {
	assert.Equal(t, "", A("").SpaceLeft())
	assert.Equal(t, "", A("bob").SpaceLeft())
	assert.Equal(t, "  ", A("  bob").SpaceLeft())
	assert.Equal(t, "    ", A("    bob").SpaceLeft())
	assert.Equal(t, "\n", A("\nbob").SpaceLeft())
	assert.Equal(t, "\t", A("\tbob").SpaceLeft())
}

func TestStrTrimPrefix(t *testing.T) {
	assert.Equal(t, "test]", A("[test]").TrimPrefix("[").A())
}

func TestStrTrimSpace(t *testing.T) {
	{
		//Left
		assert.Equal(t, "bob", A("bob").TrimSpaceLeft().A())
		assert.Equal(t, "bob", A("  bob").TrimSpaceLeft().A())
		assert.Equal(t, "bob  ", A("  bob  ").TrimSpaceLeft().A())
		assert.Equal(t, 3, A("  bob").TrimSpaceLeft().Len())
	}
	{
		// Right
		assert.Equal(t, "bob", A("bob").TrimSpaceRight().A())
		assert.Equal(t, "bob", A("bob  ").TrimSpaceRight().A())
		assert.Equal(t, "  bob", A("  bob  ").TrimSpaceRight().A())
		assert.Equal(t, 3, A("bob  ").TrimSpaceRight().Len())
	}
}

func TestStrTrimSuffix(t *testing.T) {
	assert.Equal(t, "[test", A("[test]").TrimSuffix("]").A())
}

func TestYamlIndent(t *testing.T) {
	{
		// Two spaces
		raw := `spec:
  template:
    spec: initContainers`
		data := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(raw), &data)
		assert.Nil(t, err)
		assert.Equal(t, "  ", A(raw).YamlIndent())
	}
	{
		// Four spaces
		raw := `spec:
    template:
        spec:
            initContainers:
            - foo
`
		data := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(raw), &data)
		assert.Nil(t, err)
		assert.Equal(t, "    ", A(raw).YamlIndent())
	}
}

func TestYamlInsertEmpty(t *testing.T) {
	{
		// No match nothing done
		raw := `spec:
  template:
    spec: initContainers
`
		data := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(raw), &data)
		assert.Nil(t, err)

		inserted, err := A(raw).YamlSet("line1.line2", "foo")
		assert.Nil(t, err)
		assert.Equal(t, raw, inserted.String())
	}
}

func TestYamlInsert(t *testing.T) {
	{
		// Match insert payload
		rawData := `spec:
  template:
    spec:
      initContainers:
      - name: foo
`
		// Test that the raw data is unmarshalable
		yamlData := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(rawData), &yamlData)
		assert.Nil(t, err)

		rawPayload := `- name: bar
  image: "busybox:1.25.0"
  imagePullPolicy: foobar
`
		// Test that the raw payload is unmarshalable
		yamlPayload := []map[string]interface{}{}
		err = yaml.Unmarshal([]byte(rawPayload), &yamlPayload)
		assert.Nil(t, err)

		// Marshal to get proper formatting
		bytesPayload, err := yaml.Marshal(yamlPayload)
		assert.Nil(t, err)

		// Test inserted data + payload
		inserted, err := A(rawData).YamlSet("spec.template.spec.initContainers", string(bytesPayload))
		assert.Nil(t, err)
		expected := `spec:
  template:
    spec:
      initContainers:
      - image: busybox:1.25.0
        imagePullPolicy: foobar
        name: bar
      - name: foo
`
		assert.Equal(t, expected, inserted.String())

		// Test if the new inserted data can be Unmarshaled
		backToYaml := map[string]interface{}{}
		err = yaml.Unmarshal([]byte(inserted.Bytes()), &backToYaml)
		assert.Nil(t, err)
	}
}

func TestYamlInsertCreate(t *testing.T) {
	// Match insert payload
	rawData := `spec:
  template:
    spec:
      imagePullSecrets:
      - name: foo
`
	// Test that the raw data is unmarshalable
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	rawPayload := `- name: bar
  image: "busybox:1.25.0"
  imagePullPolicy: foobar
`
	// Test that the raw payload is unmarshalable
	yamlPayload := []map[string]interface{}{}
	err = yaml.Unmarshal([]byte(rawPayload), &yamlPayload)
	assert.Nil(t, err)

	// Marshal to get proper formatting
	bytesPayload, err := yaml.Marshal(yamlPayload)
	assert.Nil(t, err)

	// Test inserted data + payload
	inserted, err := A(rawData).YamlSet("spec.template.spec.initContainers", string(bytesPayload))
	assert.Nil(t, err)
	expected := `spec:
  template:
    spec:
      initContainers:
      - image: busybox:1.25.0
        imagePullPolicy: foobar
        name: bar
      imagePullSecrets:
      - name: foo
`
	assert.Equal(t, expected, inserted.String())

	// Test if the new inserted data can be Unmarshaled
	backToYaml := map[string]interface{}{}
	err = yaml.Unmarshal([]byte(inserted.Bytes()), &backToYaml)
	assert.Nil(t, err)
}

func TestYamlInsertReplace(t *testing.T) {
	{
		rawData := `spec:
  template:
    spec:
      containers:
        name: foo
`
		// Test that the raw data is unmarshalable
		yamlData := map[string]interface{}{}
		assert.Nil(t, yaml.Unmarshal([]byte(rawData), &yamlData))

		// Test that the name was replaced
		updated, err := A(rawData).YamlSet("spec.template.spec.containers.name", "bar")
		assert.Nil(t, err)
		expected := `spec:
  template:
    spec:
      containers:
        name: foo
`
		assert.Equal(t, expected, updated.String())
	}
}

func TestYamlType(t *testing.T) {
	{
		// string
		assert.Equal(t, "test", A("\"test\"").YamlType())
		assert.Equal(t, "test", A("'test'").YamlType())
		assert.Equal(t, "1", A("\"1\"").YamlType())
		assert.Equal(t, "1", A("'1'").YamlType())
	}
	{
		// int
		assert.Equal(t, 1.0, A("1").YamlType())
		assert.Equal(t, 0.0, A("0").YamlType())
		assert.Equal(t, 25.0, A("25").YamlType())
	}
	{
		// bool
		assert.Equal(t, true, A("true").YamlType())
		assert.Equal(t, false, A("false").YamlType())
	}
	{
		// default
		assert.Equal(t, "True", A("True").YamlType())
	}
}
