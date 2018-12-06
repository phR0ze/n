package n

import (
	"strconv"
	"strings"
	"unicode"
)

type strN struct {
	v string
}

// A provides a new empty string nub
func A(str string) *strN {
	return &strN{v: str}
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

// Split creates a new nub from the split string
func (q *strN) Split(delim string) *strSliceN {
	return S(strings.Split(q.v, delim)...)
}

// TrimPrefix trims the given prefix off the string
func (q *strN) TrimPrefix(prefix string) *strN {
	return A(strings.TrimPrefix(q.v, prefix))
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

// YAMLType converts the given string into a type expected in YAML.
// Quotes signifies a string.
// No quotes signifies an int.
// true or false signifies a bool.
func (q *strN) YAMLType() interface{} {
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
