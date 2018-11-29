package nub

import (
	"strings"
)

type strNub struct {
	v string
}

// A provides a new empty string nub
func A(str string) *strNub {
	return &strNub{v: str}
}

// A materializes object invoking deferred execution
func (q *strNub) A() string {
	return q.v
}

// Q creates a queryable from the string
func (q *strNub) Q() *Queryable {
	return Q(q.v)
}

// Contains checks if the given target is contained in this string
func (q *strNub) Contains(target string) bool {
	return strings.Contains(q.v, target)
}

// ContainsAny checks if any of the targets are contained in this string
func (q *strNub) ContainsAny(targets ...string) bool {
	for i := range targets {
		if strings.Contains(q.v, targets[i]) {
			return true
		}
	}
	return false
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (q *strNub) HasAnyPrefix(prefixes ...string) bool {
	for i := range prefixes {
		if strings.HasPrefix(q.v, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (q *strNub) HasAnySuffix(suffixes ...string) bool {
	for i := range suffixes {
		if strings.HasSuffix(q.v, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (q *strNub) HasPrefix(prefix string) bool {
	return strings.HasPrefix(q.v, prefix)
}

// HasSuffix checks if the string has the given suffix
func (q *strNub) HasSuffix(suffix string) bool {
	return strings.HasSuffix(q.v, suffix)
}

// Split creates a new nub from the split string
func (q *strNub) Split(delim string) *strSliceNub {
	return StrSlice(strings.Split(q.v, delim))
}

// TrimPrefix trims the given prefix off the string
func (q *strNub) TrimPrefix(prefix string) *strNub {
	return A(strings.TrimPrefix(q.v, prefix))
}

// TrimSuffix trims the given suffix off the string
func (q *strNub) TrimSuffix(suffix string) *strNub {
	return A(strings.TrimSuffix(q.v, suffix))
}

// YAMLType converts the given string into a type expected in YAML.
// Quotes signifies a string.
// No quotes signifies an int.
// true or false signifies a bool.
func (q *strNub) YAMLType(t string) interface{} {
	//if A(str)
	return nil
}
