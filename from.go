package n

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/tmpl"
)

// Load a yaml file as a Numerable
func Load(filepath string) (result *OldNumerable, err error) {
	var yamlBytes []byte
	if yamlBytes, err = ioutil.ReadFile(filepath); err == nil {
		data := map[string]interface{}{}
		if err = yaml.Unmarshal(yamlBytes, &data); err == nil {
			result = Q(data)
		}
	}
	return
}

// FromYaml return a numerable from the given Yaml
func FromYaml(yml string) (result *OldNumerable, err error) {
	data := map[string]interface{}{}
	if err = yaml.Unmarshal([]byte(yml), &data); err == nil {
		result = Q(data)
	}
	return
}

// FromYamlTmplFile loads a yaml file from disk and processes any templating
// provided by the tmpl package returning an unmarshaled yaml block numerable.
func FromYamlTmplFile(filepath string, vars map[string]string) *OldNumerable {
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
	return Nil()
}
