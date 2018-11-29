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
func (str *strNub) A() string {
	return str.v
}

// Q creates a queryable from the string
func (str *strNub) Q() *Queryable {
	return Q(str.v)
}

// Contains checks if the given target is contained in this string
func (str *strNub) Contains(target string) bool {
	return strings.Contains(str.v, target)
}

// ContainsAny checks if any of the targets are contained in this string
func (str *strNub) ContainsAny(targets ...string) bool {
	for i := range targets {
		if strings.Contains(str.v, targets[i]) {
			return true
		}
	}
	return false
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (str *strNub) HasAnyPrefix(prefixes ...string) bool {
	for i := range prefixes {
		if strings.HasPrefix(str.v, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (str *strNub) HasAnySuffix(suffixes ...string) bool {
	for i := range suffixes {
		if strings.HasSuffix(str.v, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (str *strNub) HasPrefix(prefix string) bool {
	return strings.HasPrefix(str.v, prefix)
}

// HasSuffix checks if the string has the given suffix
func (str *strNub) HasSuffix(suffix string) bool {
	return strings.HasSuffix(str.v, suffix)
}

// Split creates a new nub from the split string
func (str *strNub) Split(delim string) *strSliceNub {
	return StrSlice(strings.Split(str.v, delim))
}

// TrimPrefix trims the given prefix off the string
func (str *strNub) TrimPrefix(prefix string) *strNub {
	return A(strings.TrimPrefix(str.v, prefix))
}

// TrimSuffix trims the given suffix off the string
func (str *strNub) TrimSuffix(suffix string) *strNub {
	return A(strings.TrimSuffix(str.v, suffix))
}
