package n

import (
	"testing"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestYaml(t *testing.T) {
	{
		// Get non existing string
		rawYaml := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.Yaml("foo").Any())
	}
	{
		// Get non existing nested string
		rawYaml := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.Yaml("foo.foo").Any())
	}
	{
		// Get string from map
		rawYaml := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "two", q.Yaml("1.2").A())
	}
	{
		// Get string from nested map
		rawYaml := `1:
  2:
    3: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "three", q.Yaml("1.2.3").A())
	}
	{
		// Get map from map
		rawYaml := `1:
  2: two`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		expected := map[string]interface{}{"2": "two"}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.Yaml("1").M())
	}
	{
		// Get map from map from map
		rawYaml := `1:
  2:
    3: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		expected := map[string]interface{}{"3": "three"}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.Yaml("1.2").M())
	}
	{
		// Get slice from map
		rawYaml := `foo:
  - 1
  - 2
  - 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, []string{"1", "2", "3"}, q.Yaml("foo").Strs())
	}
}

func TestYamlWithKeyIndexing(t *testing.T) {
	{
		// Select map from slice from map
		rawYaml := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		expected := map[string]interface{}{"name": 2.0}

		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, expected, q.Yaml("foo.[name:2]").M())
	}
	{
		// Bad key
		rawYaml := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.Yaml("fee.[name:2]").Any())
	}
	{
		// Bad sub key
		rawYaml := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.Yaml("foo.[fee:2]").Any())
	}
	{
		// Missing target
		rawYaml := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)

		q := Q(data)
		assert.True(t, q.Any())
		assert.False(t, q.Yaml("foo.[name:5]").Any())
	}
	{
		// Continue keying in after slice: one
		rawYaml := `foo:
  - name: one
  - name: two
  - name: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "one", q.Yaml("foo.[name:one].name").O())
	}
	// 	{
	// TODO: implement this ???
	// 		// Continue keying in after slice: two
	// 		rawYaml := `foo:
	//   - name:
	//       bar: frodo
	//       foo: blah
	//   - name: two
	//   - name: three`
	// 		data := map[string]interface{}{}
	// 		yaml.Unmarshal([]byte(rawYaml), &data)
	// 		q := Q(data)
	// 		assert.True(t, q.Any())
	// 		assert.Equal(t, "frodo", q.Yaml("foo.[name.bar:frodo].name.bar").O())
	// 	}
}

func TestYamlWithSliceIndexing(t *testing.T) {
	{
		rawYaml := `foo:
  - name: 1
  - name: 2
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		{
			expected := map[string]interface{}{"name": 1.0}
			assert.Equal(t, expected, q.Yaml("foo.[0]").M())
		}
		{
			expected := map[string]interface{}{"name": 2.0}
			assert.Equal(t, expected, q.Yaml("foo.[1]").M())
		}
		{
			expected := map[string]interface{}{"name": 3.0}
			assert.Equal(t, expected, q.Yaml("foo.[2]").M())
		}
		{
			expected := map[string]interface{}{"name": 3.0}
			assert.Equal(t, expected, q.Yaml("foo.[-1]").M())
		}
		{
			expected := map[string]interface{}{"name": 2.0}
			assert.Equal(t, expected, q.Yaml("foo.[-2]").M())
		}
		{
			expected := map[string]interface{}{"name": 1.0}
			assert.Equal(t, expected, q.Yaml("foo.[-3]").M())
		}
	}
	{
		// Select first element when only one
		rawYaml := `foo:
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		expected := map[string]interface{}{"name": 3.0}
		assert.Equal(t, expected, q.Yaml("foo.[0]").M())
		assert.Equal(t, expected, q.Yaml("foo.[-1]").M())
	}
	{
		// Select first element when only one
		rawYaml := `foo:
  - name: 3`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		expected := map[string]interface{}{"name": 3.0}
		assert.Equal(t, expected, q.Yaml("foo.[0]").M())
		assert.Equal(t, expected, q.Yaml("foo.[-1]").M())
	}
	{
		// Continue keying in after slice: one
		rawYaml := `foo:
  - name: one
  - name: two
  - name: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "one", q.Yaml("foo.[0].name").O())
	}
	{
		// Continue keying in after slice: two
		rawYaml := `foo:
  - name:
      bar: frodo
  - name: two
  - name: three`
		data := map[string]interface{}{}
		yaml.Unmarshal([]byte(rawYaml), &data)
		q := Q(data)
		assert.True(t, q.Any())
		assert.Equal(t, "frodo", q.Yaml("foo.[0].name.bar").O())
	}
}

func TestYamlReplace(t *testing.T) {
	{
		rawYaml := `foo:
  - name: bar
    valueFrom:
      secretKeyRef:
        name: <% template "foo" . %>
        key: keybar`
		data := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(rawYaml), &data)
		assert.Nil(t, err)
		expected := map[string]interface{}{
			"foo": []interface{}{
				map[string]interface{}{
					"name": "bar", "valueFrom": map[string]interface{}{
						"secretKeyRef": map[string]interface{}{"name": "{{ template \"foo\" . }}", "key": "keybar"},
					}}}}

		values := map[string]string{"<%": "{{", "%>": "}}"}
		assert.Equal(t, expected, YamlReplace(data, values))
	}
	{
		rawYaml := `deployment:
  initContainers:
    - name: init-mysql
      env:
        - name: MYSQL
          valueFrom:
            secretKeyRef:
              name: <% template "foo" . %>
              key: password`
		q, err := FromYaml(rawYaml)
		assert.Nil(t, err)
		{
			data := q.Yaml("deployment.initContainers").S()
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
			assert.Equal(t, expected, YamlReplace(data, values))
		}
		{
			data := q.Yaml("deployment.initContainers").SAMap()
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
			assert.Equal(t, expected, YamlReplace(data, values))
		}
	}
}

func TestYamlSetInsertRoot(t *testing.T) {
	{
		// key path doesn't exist so it gets created
		raw := `spec:
  template:
    spec: initContainers
`
		data := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(raw), &data)
		assert.Nil(t, err)

		inserted, err := Q(data).YamlSet("line1.line2", "foo")
		assert.Nil(t, err)

		expected := map[string]interface{}{
			"line1": map[string]interface{}{"line2": "foo"},
			"spec":  map[string]interface{}{"template": map[string]interface{}{"spec": "initContainers"}},
		}
		assert.Equal(t, expected, inserted.M())
	}
}

func TestYamlSetInsertNested(t *testing.T) {
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
  imagePullPolicy: Always
`
	yamlPayload := []map[string]interface{}{}
	err = yaml.Unmarshal([]byte(rawPayload), &yamlPayload)
	assert.Nil(t, err)

	// Test inserted data + payload
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.initContainers", yamlPayload)
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"initContainers": []map[string]interface{}{
				map[string]interface{}{"name": "bar", "image": "busybox:1.25.0", "imagePullPolicy": "Always"},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetInsertByIndex(t *testing.T) {
	// Match insert payload
	rawData := `spec:
  template:
     spec:
       containers:
       - name: foo
         image: foo:latest
`
	// Test that the raw data is unmarshalable
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	// Test inserted data + payload
	data := map[string]interface{}{"name": "bar", "image": "bar:1.2.3"}
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[1]", data)
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "foo", "image": "foo:latest"},
				map[string]interface{}{"name": "bar", "image": "bar:1.2.3"},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetOverrideByIndex(t *testing.T) {
	// Match insert payload
	rawData := `spec:
  template:
     spec:
       containers:
       - name: foo
         image: foo:latest
`
	// Test that the raw data is unmarshalable
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	// Test inserted data + payload
	data := map[string]interface{}{"name": "bar", "image": "bar:1.2.3"}
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[0]", data)
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "bar", "image": "bar:1.2.3"},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetOverrideByName(t *testing.T) {
	// Match insert payload
	rawData := `spec:
  template:
     spec:
       containers:
       - name: foo
         image: foo:latest
`
	// Test that the raw data is unmarshalable
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	// Test inserted data + payload
	data := map[string]interface{}{"name": "bar", "image": "bar:1.2.3"}
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[name:foo]", data)

	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "bar", "image": "bar:1.2.3"},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetUpdateListByIndex(t *testing.T) {
	// Match insert payload
	rawData := `spec:
  template:
    spec:
      containers:
      - name: foo
        image: fobar:latest
`
	// Test that the raw data is unmarshalable
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	// Test inserted data + payload
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[0].image", "foobar:1.2.3")
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "foo", "image": "foobar:1.2.3"},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetUpdateListItemByName(t *testing.T) {
	// Match insert payload
	rawData := `spec:
  template:
    spec:
      containers:
      - name: foo
        image: fobar:latest
`
	// Test that the raw data is unmarshalable
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	// Test inserted data + payload
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[name:foo].image", "foobar:1.2.3")
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "foo", "image": "foobar:1.2.3"},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetDeepNesting(t *testing.T) {
	rawData := `spec:
  template:
    spec:
      containers:
      - name: foo
        env:
        - name: var1
          value: foobar1
        - name: var2
          value: foobar2
        - name: var3
          value: foobar3
`
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[0].env.[name:var2].value", "foofoo")
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "foo", "env": []interface{}{
					map[string]interface{}{"name": "var1", "value": "foobar1"},
					map[string]interface{}{"name": "var2", "value": "foofoo"},
					map[string]interface{}{"name": "var3", "value": "foobar3"},
				}},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetInsertIndexAtBegining(t *testing.T) {
	rawData := `spec:
  template:
    spec:
      containers:
      - name: foo
        env:
        - name: var1
          value: foobar1
`
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	data := map[string]interface{}{"name": "var5", "value": "foobar5"}
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[0].env.[[0]]", data)
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "foo", "env": []interface{}{
					map[string]interface{}{"name": "var5", "value": "foobar5"},
					map[string]interface{}{"name": "var1", "value": "foobar1"},
				}},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetInsertIndexMiddle(t *testing.T) {
	rawData := `spec:
  template:
    spec:
      containers:
      - name: foo
        env:
        - name: var1
          value: foobar1
        - name: var2
          value: foobar2
`
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	data := map[string]interface{}{"name": "var5", "value": "foobar5"}
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[0].env.[[1]]", data)
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "foo", "env": []interface{}{
					map[string]interface{}{"name": "var1", "value": "foobar1"},
					map[string]interface{}{"name": "var5", "value": "foobar5"},
					map[string]interface{}{"name": "var2", "value": "foobar2"},
				}},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}

func TestYamlSetInsertAtName(t *testing.T) {
	rawData := `spec:
  template:
    spec:
      containers:
      - name: foo
        env:
        - name: var1
          value: foobar1
        - name: var2
          value: foobar2
`
	yamlData := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(rawData), &yamlData)
	assert.Nil(t, err)

	data := map[string]interface{}{"name": "var5", "value": "foobar5"}
	inserted, err := Q(yamlData).YamlSet("spec.template.spec.containers.[0].env.[[name:var2]]", data)
	assert.Nil(t, err)
	expected := map[string]interface{}{
		"spec": map[string]interface{}{"template": map[string]interface{}{"spec": map[string]interface{}{
			"containers": []interface{}{
				map[string]interface{}{"name": "foo", "env": []interface{}{
					map[string]interface{}{"name": "var1", "value": "foobar1"},
					map[string]interface{}{"name": "var5", "value": "foobar5"},
					map[string]interface{}{"name": "var2", "value": "foobar2"},
				}},
			},
		}}}}
	assert.Equal(t, expected, inserted.M())
}
