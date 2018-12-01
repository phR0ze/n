package n

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/tmpl"
)

// FromYAMLFile a yaml/json file as a str map
// returns nil on failure of any kind
func FromYAMLFile(filepath string) *Queryable {
	if yamlFile, err := ioutil.ReadFile(filepath); err == nil {
		data := map[string]interface{}{}
		if err := yaml.Unmarshal(yamlFile, &data); err == nil {
			return Q(data)
		}
	}
	return nil
}

// FromYAML return a queryable from the given YAML
func FromYAML(yml string) *Queryable {
	data := map[string]interface{}{}
	if err := yaml.Unmarshal([]byte(yml), &data); err == nil {
		return Q(data)
	}
	return nil
}

// FromYAMLTmplFile loads a yaml file from disk and processes any templating
func FromYAMLTmplFile(filepath string, vars map[string]string) *Queryable {
	if data, err := ioutil.ReadFile(filepath); err == nil {
		if tpl, err := tmpl.New(string(data), "{{", "}}"); err == nil {
			if result, err := tpl.Process(vars); err == nil {
				m := map[string]interface{}{}
				if err := yaml.Unmarshal([]byte(result), &m); err == nil {
					return Q(m)
				}
			}
		}
	}
	return nil
}

// LoadTmplFile loads a yaml file from disk and processes any templating
func LoadTmplFile(filepath string, vars map[string]string) string {
	if data, err := ioutil.ReadFile(filepath); err == nil {
		if tpl, err := tmpl.New(string(data), "{{", "}}"); err == nil {
			if result, err := tpl.Process(vars); err == nil {
				return result
			}
		}
	}
	return ""
}

// YAML gets data by key which can be dot delimited
func (q *Queryable) YAML(key string) (result *Queryable) {
	keys := A(key).Split(".")
	if key, ok := keys.TakeFirst(); ok {
		switch x := q.v.Interface().(type) {
		case map[string]interface{}:
			if !A(key).ContainsAny(":", "[", "]") {
				if v, ok := x[key]; ok {
					result = Q(v)
				}
			}
		case []interface{}:
			k, v := A(key).TrimPrefix("[").TrimSuffix("]").Split(":").YAMLPair()
			for i := range x {
				if m, ok := x[i].(map[string]interface{}); ok {
					if entry, ok := m[k]; ok {
						if v == entry {
							return Q(m)
						}
					}
				}
			}
		}
		if keys.Len() != 0 && result.Any() {
			result = result.YAML(keys.Join(".").A())
		}
	}
	if result == nil {
		result = N()
	}
	return
}
