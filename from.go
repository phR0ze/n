package n

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/phR0ze/n/pkg/nos"
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
	var lines []string
	if lines, err = nos.ReadLines(filepath); err == nil {
		buff := convHelmTplToYaml(lines)
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
func convHelmTplToYaml(lines []string) []byte {
	cnt := 0
	prevDepth := ""
	var buff bytes.Buffer

	type yamlType int
	type yamlItem struct {
		typ yamlType
		val string
	}
	const (
		typeMap yamlType = iota
		typeList
		typeEmpty
		typeTplLine
		typeMapItem
		typeListItem
	)

	// Tokenize yaml
	prevType := typeEmpty
	items := []yamlItem{}
	for i := range lines {
		a := A(lines[i])
		if !a.Empty() && !a.TrimSpaceLeft().HasPrefix("#") {
			typ := typeEmpty
			if a.TrimSpaceLeft().HasPrefix("{{") {
				typ = typeTplLine
			} else {
				dash, colon := strings.Index(lines[i], "-"), strings.Index(lines[i], ":")
				if (colon > -1 && dash == -1) || (colon > -1 && colon < dash) {
					_, s := a.SplitOn(":")
					if A(s).Empty() {
						// map
						// list
						typ = typeMap
					} else {
						// inline map
						// inline list
						// map item
						typ = typeMapItem
					}
				} else if (dash > -1 && colon == -1) || (dash > -1 && dash < colon) {
					typ = typeListItem
					if items[i-1].typ == typeMap {
						items[i-1].typ = typeList
					}
				}
			}
			items = append(items, yamlItem{typ, lines[i]})
		}
	}

	validYamlModifier := "validYamlModifier"
	for i := range lines {
		line := lines[i]

		a := A(line)
		mapType := false
		if !a.Empty() && !a.TrimSpaceLeft().HasPrefix("#") {
			if a.TrimSpaceLeft().HasPrefix("{{") {
				// Escape template line
				line = fmt.Sprintf("%s%s-%d: %s", prevDepth, validYamlModifier, cnt, helmEscape(line))
				cnt++
			} else {

				// Escape template value
				if a.Contains("{{") {
					first, second := a.SplitOn(": ")
					if second == "" {
						first, second = a.SplitOn("- ")
					}
					line = fmt.Sprintf("%s%s", first, helmEscape(second))
				} else {
					prevDepth = A(line).SpaceLeft()
				}
			}
		}
		buff.WriteString(line)
		buff.WriteString("\n")
		fmt.Println(line)
	}

	return buff.Bytes()
}

// Escape helm templating
func helmEscape(in string) (out string) {
	return fmt.Sprintf("\"%s\"", A(in).
		Replace(":", "[[C]]").
		Replace("-", "[[T]]").
		Replace("\"", "[[DQ]]").
		Replace("{{", "[[").
		Replace("}}", "]]").
		A())
}
