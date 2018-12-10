package n

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"

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
		tmpBytes := A(fileBytes).Replace("{{", "<<").Replace("}}", ">>").B()

		// Now convert unnamed variables to something that unmarshals
		i := 0
		prevDepth := ""
		validYamlModifier := "validYamlModifier"
		scanner := bufio.NewScanner(bytes.NewReader(tmpBytes))
		for scanner.Scan() {
			line := scanner.Text()
			a := A(line)
			if a.TrimSpaceLeft().HasPrefix("<<") {
				newline := fmt.Sprintf("%s%s-%d: %s\n", prevDepth, validYamlModifier, i, line)
				buff.WriteString(newline)
				fmt.Printf(newline)
				i++
			} else {
				buff.WriteString(line)
				buff.WriteString("\n")
				if !a.Empty() && !a.TrimSpaceLeft().HasPrefix("#") {
					prevDepth = A(line).SpaceLeft()
				}
				fmt.Println(line)
			}
		}
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
