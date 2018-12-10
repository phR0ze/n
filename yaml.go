package n

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/ntmpl"
)

// FromYamlFile a yaml/json file as a str map
// returns nil on failure of any kind
func FromYamlFile(filepath string) (result *Queryable, err error) {
	var yamlBytes []byte
	if yamlBytes, err = ioutil.ReadFile(filepath); err == nil {
		data := map[string]interface{}{}
		if err = yaml.Unmarshal(yamlBytes, &data); err == nil {
			result = Q(data)
		}
	}
	return
}

// FromYaml return a queryable from the given Yaml
func FromYaml(yml string) (result *Queryable, err error) {
	data := map[string]interface{}{}
	if err = yaml.Unmarshal([]byte(yml), &data); err == nil {
		result = Q(data)
	}
	return
}

// FromHelmFile loads a helm yaml file from disk and converts templating to valid yaml
// before working with it.
func FromHelmFile(filepath string) (result *Queryable, err error) {
	result = N()
	var fileBytes []byte
	if fileBytes, err = ioutil.ReadFile(filepath); err == nil {
		var buff bytes.Buffer

		// First convert named variables to something that unmarshals
		tmpBytes := A(fileBytes).Replace("{{", "<%").Replace("}}", "%>").B()

		// Now convert unnamed variables to something that unmarshals
		//i := 0
		//prevDepth := ""
		//validYaml
		scanner := bufio.NewScanner(bytes.NewReader(tmpBytes))
		for scanner.Scan() {
			line := scanner.Text()

			if A(line).TrimSpaceLeft().HasPrefix("<%") {
				//buff.WriteString(fmt.Sprintf("%s", prevDepth, line))
				buff.WriteString("#")
				buff.WriteString(line)
				buff.WriteString("\n")
			} else {
				buff.WriteString(line)
				buff.WriteString("\n")
			}

			//prevDepth := A(line).SpaceLeft()
		}

		// Convert to Yaml
		data := map[string]interface{}{}
		err = yaml.Unmarshal(buff.Bytes(), &data)
	}
	return
}

// FromYamlTmplFile loads a yaml file from disk and processes any templating
// provided by the ntmpl package returning an unmarshaled yaml block queryable.
func FromYamlTmplFile(filepath string, vars map[string]string) *Queryable {
	if data, err := ioutil.ReadFile(filepath); err == nil {
		if tpl, err := ntmpl.New(string(data), "{{", "}}"); err == nil {
			if result, err := tpl.Process(vars); err == nil {
				m := map[string]interface{}{}
				if err := yaml.Unmarshal([]byte(result), &m); err == nil {
					return Q(m)
				}
			}
		}
	}
	return N()
}

// Yaml gets data by key which can be dot delimited
// returns nil queryable on errors or keys not found
func (q *Queryable) Yaml(key string) (result *Queryable) {
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
			k, v := A(key).TrimPrefix("[").TrimSuffix("]").Split(":").YamlPair()
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
			result = result.Yaml(keys.Join(".").A())
		}
	}
	if result == nil {
		result = N()
	}
	return
}

// YamlReplace recursively makes string substitutions
func YamlReplace(data interface{}, values map[string]string) (result interface{}) {
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
				x[k] = YamlReplace(v, values)
			}
		}
		result = data
	case []map[string]interface{}:
		resultSlice := []map[string]interface{}{}
		for i := range x {
			v := YamlReplace(x[i], values)
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
				value = YamlReplace(y, values)
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
