package n

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/tmpl"
)

// Load a yaml file as a Queryable
func Load(filepath string) (result *Queryable, err error) {
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

// FromYamlTmplFile loads a yaml file from disk and processes any templating
// provided by the tmpl package returning an unmarshaled yaml block queryable.
func FromYamlTmplFile(filepath string, vars map[string]string) *Queryable {
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
	return N()
}
