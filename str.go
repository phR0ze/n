package n

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type strN struct {
	v string
}

// A provides a new empty string nub
// handling both string and []byte
func A(str interface{}) (result *strN) {
	result = &strN{v: ""}
	switch x := str.(type) {
	case string:
		result.v = x
	case []byte:
		result.v = string(x)
	case int:
		strconv.Itoa(x)
	default:
		result.v = fmt.Sprintf("%v", x)
	}
	return
}

// B exports object as []byte
func (q *strN) B() []byte {
	return []byte(q.v)
}

// A exports object invoking deferred execution
func (q *strN) A() string {
	return q.v
}

// Q creates a queryable from the string
func (q *strN) Q() *Queryable {
	return Q(q.v)
}

// Contains checks if the given target is contained in this string
func (q *strN) Contains(target string) bool {
	return strings.Contains(q.v, target)
}

// ContainsAny checks if any of the targets are contained in this string
func (q *strN) ContainsAny(targets ...string) bool {
	for i := range targets {
		if strings.Contains(q.v, targets[i]) {
			return true
		}
	}
	return false
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (q *strN) HasAnyPrefix(prefixes ...string) bool {
	for i := range prefixes {
		if strings.HasPrefix(q.v, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (q *strN) HasAnySuffix(suffixes ...string) bool {
	for i := range suffixes {
		if strings.HasSuffix(q.v, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (q *strN) HasPrefix(prefix string) bool {
	return strings.HasPrefix(q.v, prefix)
}

// HasSuffix checks if the string has the given suffix
func (q *strN) HasSuffix(suffix string) bool {
	return strings.HasSuffix(q.v, suffix)
}

// Len returns the length of the string
func (q *strN) Len() int {
	return len(q.v)
}

// Replace allows for chaining and default to all instances
func (q *strN) Replace(new, old string, ns ...int) *strN {
	n := -1
	if len(ns) > 0 {
		n = ns[0]
	}
	return A(strings.Replace(q.v, new, old, n))
}

// SpaceLeft returns leading whitespace
func (q *strN) SpaceLeft() string {
	spaces := []rune{}
	for _, r := range q.v {
		if unicode.IsSpace(r) {
			spaces = append(spaces, r)
		} else {
			break
		}
	}
	return string(spaces)
}

// Split creates a new nub from the split string
func (q *strN) Split(delim string) *strSliceN {
	return S(strings.Split(q.v, delim)...)
}

// TrimPrefix trims the given prefix off the string
func (q *strN) TrimPrefix(prefix string) *strN {
	return A(strings.TrimPrefix(q.v, prefix))
}

// TrimSpace pass through to strings.TrimSpace
func (q *strN) TrimSpace() *strN {
	return A(strings.TrimSpace(q.v))
}

// TrimSpaceLeft trims leading whitespace
func (q *strN) TrimSpaceLeft() *strN {
	return A(strings.TrimLeftFunc(q.v, unicode.IsSpace))
}

// TrimSpaceRight trims trailing whitespace
func (q *strN) TrimSpaceRight() *strN {
	return A(strings.TrimRightFunc(q.v, unicode.IsSpace))
}

// TrimSuffix trims the given suffix off the string
func (q *strN) TrimSuffix(suffix string) *strN {
	return A(strings.TrimSuffix(q.v, suffix))
}

// YamlIndent detects the indent string being used and returns it
func (q *strN) YamlIndent() string {
	scanner := bufio.NewScanner(bytes.NewReader([]byte(q.v)))
	for scanner.Scan() {
		line := scanner.Text()
		space := A(line).SpaceLeft()
		if space != "" {
			return space
		}
	}
	return "  "
}

// YamlSet inserts the block of text at the specified yaml location.
// This operation is done on raw text to avoid Yaml syntax errors when working
// with Helm/Go templating syntax that hasn't been processed yet.
// Supports a primitive dot delimited keying notation.
// e.g. spec.template.spec.initContainers
func (q *strN) YamlSet(key, block string) (data bytes.Buffer, err error) {
	depth := ""
	indent := q.YamlIndent()

	// Primitive dot notation keying
	ok := false
	keys := A(key).Split(".")
	if key, ok = keys.TakeFirst(); !ok {
		return
	}
	keyNotFound := false
	if !q.Contains(keys.Last().A()) {
		keyNotFound = true
	}

	// Scan through data line by line writing to data
	writeBlock := false
	writer := bufio.NewWriter(&data)
	scanner := bufio.NewScanner(bytes.NewReader([]byte(q.v)))
	for scanner.Scan() {
		line := scanner.Text()
		a := A(string(line)).TrimSpaceLeft()
		curDepth := A(string(line)).SpaceLeft()

		// Write out the block now so we can use curDepth
		if writeBlock {
			writeBlock = false
			if !keys.Any() && keyNotFound {
				writer.WriteString(curDepth + key + ":\n")
			}
			blockScanner := bufio.NewScanner(bytes.NewReader([]byte(block)))
			for blockScanner.Scan() {
				writer.WriteString(curDepth + blockScanner.Text())
				writer.WriteString("\n")
			}
		}

		// Hit - flag write block for next loop
		if ok && depth == curDepth && a.HasPrefix(key+":") {
			depth += indent
			if key, ok = keys.TakeFirst(); !ok || (!keys.Any() && keyNotFound) {

				// Inserts
				writeBlock = true
			}
		}

		// Write out the current line
		writer.WriteString(line)
		writer.WriteString("\n")
	}

	writer.Flush()
	return
}

// YamlType converts the given string into a type expected in Yaml.
// Quotes signifies a string.
// No quotes signifies an int.
// true or false signifies a bool.
func (q *strN) YamlType() interface{} {
	if q.HasAnyPrefix("\"", "'") && q.HasAnySuffix("\"", "'") {
		return q.v[1 : len(q.v)-1]
	} else if q.v == "true" || q.v == "false" {
		if b, err := strconv.ParseBool(q.v); err == nil {
			return b
		}
	} else if f, err := strconv.ParseFloat(q.v, 32); err == nil {
		return f
	}
	return q.v
}
