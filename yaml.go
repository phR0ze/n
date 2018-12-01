package n

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/tmpl"
)

// LoadYAML a yaml/json file as a str map
// returns nil on failure of any kind
func LoadYAML(filepath string) *Queryable {
	if yamlFile, err := ioutil.ReadFile(filepath); err == nil {
		data := map[string]interface{}{}
		yaml.Unmarshal(yamlFile, &data)
		return Q(data)
	}
	return nil
}

// LoadYAMLTmpl loads a yaml file from disk and processes any templating
func LoadYAMLTmpl(filepath string, vars map[string]string) *Queryable {
	if data, err := ioutil.ReadFile(filepath); err == nil {
		if tpl, err := tmpl.New(string(data), "{{", "}}"); err == nil {
			if result, err := tpl.Process(vars); err == nil {
				return Q(result)
			}
		}
	}
	return N()
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
