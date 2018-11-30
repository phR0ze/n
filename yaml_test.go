package n

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestYAML(t *testing.T) {
	{
		// Get string from map
		rawYAML := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "two", q.YAML("1.2").A())
	}
	{
		// Get string from nested map
		rawYAML := `1:
  2:
    3: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "three", q.YAML("1.2.3").A())
	}
	{
		// Get map from map
		rawYAML := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := map[string]interface{}{"2": "two"}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.YAML("1").M())
	}
	{
		// Get map from map from map
		rawYAML := `1:
  2:
    3: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := map[string]interface{}{"3": "three"}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.YAML("1.2").M())
	}
	{
		// Get slice from map
		rawYAML := `foo:
  - 1
  - 2
  - 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, []string{"1", "2", "3"}, q.YAML("foo").Strs())
	}
	{
		// Select map from slice from map
		rawYAML := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		expected := map[string]interface{}{"name": 2.0}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.YAML("foo.[name:2]").M())
	}
}
