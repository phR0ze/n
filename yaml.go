package n

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/tmpl"
)

// FromYAMLFile a yaml/json file as a str map
// returns nil on failure of any kind
func FromYAMLFile(filepath string) (result *Queryable, err error) {
	var yamlBytes []byte
	if yamlBytes, err = ioutil.ReadFile(filepath); err == nil {
		data := map[string]interface{}{}
		if err = yaml.Unmarshal(yamlBytes, &data); err == nil {
			result = Q(data)
		}
	}
	return
}

// FromYAML return a queryable from the given YAML
func FromYAML(yml string) (result *Queryable, err error) {
	data := map[string]interface{}{}
	if err = yaml.Unmarshal([]byte(yml), &data); err == nil {
		result = Q(data)
	}
	return
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
func LoadTmplFile(filepath string, startTag, endTag string, vars map[string]string) string {
	if data, err := ioutil.ReadFile(filepath); err == nil {
		if tpl, err := tmpl.New(string(data), startTag, endTag); err == nil {
			if result, err := tpl.Process(vars); err == nil {
				return result
			}
		}
	}
	return ""
}

// YAML gets data by key which can be dot delimited
// returns nil queryable on errors or keys not found
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
			if v == nil {
				if i, err := strconv.Atoi(k); err == nil {
					result = q.At(i)
				} else {
					panic(errors.New("Failed to convert index to an int"))
				}
			} else {
				for i := range x {
					if m, ok := x[i].(map[string]interface{}); ok {
						if entry, ok := m[k]; ok {
							if v == entry {
								result = Q(m)
								break
							}
						}
					}
				}
			}
		}
		if keys.Len() != 0 && result != nil && result.Any() {
			result = result.YAML(keys.Join(".").A())
		}
	}
	if result == nil {
		result = N()
	}
	return
}

// YAMLReplace recursively makes string substitutions
func YAMLReplace(data interface{}, values map[string]string) (result interface{}) {
	switch x := data.(type) {
	case map[string]interface{}:
		for k, v := range x {
			switch y := v.(type) {
			case string:
				for o, n := range values {
					y = strings.Replace(y, o, n, -1)
				}
				x[k] = y
			case []interface{}, map[string]interface{}:
				x[k] = YAMLReplace(v, values)
			}
		}
		result = data
	case []map[string]interface{}:
		resultSlice := []map[string]interface{}{}
		for i := range x {
			v := YAMLReplace(x[i], values)
			resultSlice = append(resultSlice, v.(map[string]interface{}))
		}
		result = resultSlice
	case []interface{}:
		resultSlice := []interface{}{}
		for i := range x {
			var value interface{}
			switch y := x[i].(type) {
			case string:
				for o, n := range values {
					y = strings.Replace(y, o, n, -1)
				}
				value = y
			case []interface{}, map[string]interface{}:
				value = YAMLReplace(y, values)
			default:
				value = y
			}
			resultSlice = append(resultSlice, value)
		}
		result = resultSlice
	default:
		result = data
	}
	return
}
