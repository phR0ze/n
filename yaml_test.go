package n

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestYAML(t *testing.T) {
	{
		// Get non existing string
		rawYAML := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.YAML("foo").Any())
	}
	{
		// Get non existing nested string
		rawYAML := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.YAML("foo.foo").Any())
	}
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
}

func TestYAMLWithKeyIndexing(t *testing.T) {
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
	{
		// Bad key
		rawYAML := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.YAML("fee.[name:2]").Any())
	}
	{
		// Bad sub key
		rawYAML := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.YAML("foo.[fee:2]").Any())
	}
	{
		// Missing target
		rawYAML := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.YAML("foo.[name:5]").Any())
	}
	{
		// Continue keying in after slice: one
		rawYAML := `foo:
  - name: one
  - name: two
  - name: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "one", q.YAML("foo.[name:one].name").O())
	}
	// 	{
	// TODO: implement this ???
	// 		// Continue keying in after slice: two
	// 		rawYAML := `foo:
	//   - name:
	//       bar: frodo
	//       foo: blah
	//   - name: two
	//   - name: three`
	// 		data := map[string]interface{}{}
	// 		yaml.Unmarshal([]byte(rawYAML), &data)
	// 		q := Q(data)
	// 		assert.True(t, q.Any())
	// 		assert.Equal(t, "frodo", q.YAML("foo.[name.bar:frodo].name.bar").O())
	// 	}
}

func TestYAMLWithSliceIndexing(t *testing.T) {
	{
		rawYAML := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		{
			expected := map[string]interface{}{"name": 1.0}
			assert.Equal(t, expected, q.YAML("foo.[0]").M())
		}
		{
			expected := map[string]interface{}{"name": 2.0}
			assert.Equal(t, expected, q.YAML("foo.[1]").M())
		}
		{
			expected := map[string]interface{}{"name": 3.0}
			assert.Equal(t, expected, q.YAML("foo.[2]").M())
		}
		{
			expected := map[string]interface{}{"name": 3.0}
			assert.Equal(t, expected, q.YAML("foo.[-1]").M())
		}
		{
			expected := map[string]interface{}{"name": 2.0}
			assert.Equal(t, expected, q.YAML("foo.[-2]").M())
		}
		{
			expected := map[string]interface{}{"name": 1.0}
			assert.Equal(t, expected, q.YAML("foo.[-3]").M())
		}
	}
	{
		// Select first element when only one
		rawYAML := `foo:
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		expected := map[string]interface{}{"name": 3.0}
		assert.Equal(t, expected, q.YAML("foo.[0]").M())
		assert.Equal(t, expected, q.YAML("foo.[-1]").M())
	}
	{
		// Continue keying in after slice: one
		rawYAML := `foo:
  - name: one
  - name: two
  - name: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "one", q.YAML("foo.[0].name").O())
	}
	{
		// Continue keying in after slice: two
		rawYAML := `foo:
  - name:
      bar: frodo
  - name: two
  - name: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYAML), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "frodo", q.YAML("foo.[0].name.bar").O())
	}
}

func TestYAMLReplace(t *testing.T) {
	{
		rawYAML := `foo:
  - name: bar
    valueFrom:
      secretKeyRef:
        name: <% template "foo" . %>
        key: keybar`
		data := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(rawYAML), &data)
		assert.Nil(t, err)
		expected := map[string]interface{}{
			"foo": []interface{}{
				map[string]interface{}{
					"name": "bar", "valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{"name": "{{ template \"foo\" . }}", "key": "keybar"},
					}}}}

		values := map[string]string{"<%": "{{", "%>": "}}"}
		assert.Equal(t, expected, YAMLReplace(data, values))
	}
	{
		rawYAML := `deployment:
  initContainers:
    - name: init-mysql
      env:
        - name: MYSQL
          valueFrom:
            secretKeyRef:
              name: <% template "foo" . %>
              key: password`
		q, err := FromYAML(rawYAML)
		assert.Nil(t, err)
		{
			data := q.YAML("deployment.initContainers").S()
			values := map[string]string{"<%": "{{", "%>": "}}"}
			expected := []interface{}{
				map[string]interface{}{
					"name": "init-mysql",
					"env": []interface{}{
						map[string]interface{}{
							"name": "MYSQL", "valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{"key": "password", "name": "{{ template \"foo\" . }}"},
							}},
					}},
			}
			assert.Equal(t, expected, YAMLReplace(data, values))
		}
		{
			data := q.YAML("deployment.initContainers").SAMap()
			values := map[string]string{"<%": "{{", "%>": "}}"}
			expected := []map[string]interface{}{
				map[string]interface{}{
					"name": "init-mysql",
					"env": []interface{}{
						map[string]interface{}{
							"name": "MYSQL", "valueFrom": map[string]interface{}{
								"secretKeyRef": map[string]interface{}{"key": "password", "name": "{{ template \"foo\" . }}"},
							}},
					}},
			}
			assert.Equal(t, expected, YAMLReplace(data, values))
		}
	}
}
