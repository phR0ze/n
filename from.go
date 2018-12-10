package n

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"

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
		buff := convHelmTplToYaml(fileBytes)
		data := map[string]interface{}{}
		err = yaml.Unmarshal(buff, &data)
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

// convHelmTplToYaml converts the given helm templat to a valid yaml file.
// It does this by replacing templating with placeholders.
func convHelmTplToYaml(fileBytes []byte) []byte {
	i := 0
	prevDepth := ""
	var buff bytes.Buffer
	validYamlModifier := "validYamlModifier"

	scanner := bufio.NewScanner(bytes.NewReader(fileBytes))
	for scanner.Scan() {
		line := scanner.Text()
		a := A(line)
		if a.TrimSpaceLeft().HasPrefix("{{") {
			line = fmt.Sprintf("%s%s-%d: %s", prevDepth, validYamlModifier, i, url.QueryEscape(line))
			i++
		} else if a.Contains("{{") {
			//if pieces.Len() > 1 {
			//	line = fmt.Sprintf("%s:%s", pieces.First().A(), url.QueryEscape(pieces.Slice(1, -1).Join(":").A()))
			//}
		} else {
			if !a.Empty() && !a.TrimSpaceLeft().HasPrefix("#") {
				prevDepth = A(line).SpaceLeft()
			}
		}
		buff.WriteString(line)
		buff.WriteString("\n")
		fmt.Println(line)
	}

	return buff.Bytes()
}
