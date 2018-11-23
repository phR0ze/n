package nub

import "strings"

type strNub struct {
	Raw string
}

// Str creates a new nub from the given string
func Str(slice string) *strNub {
	return &strNub{Raw: slice}
}

// Contains checks if the given target is contained in this string
func (str *strNub) Contains(target string) bool {
	return strings.Contains(str.Raw, target)
}

// ContainsAny checks if any of the targets are contained in this string
func (str *strNub) ContainsAny(targets []string) bool {
	for i := range targets {
		if strings.Contains(str.Raw, targets[i]) {
			return true
		}
	}
	return false
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (str *strNub) HasAnyPrefix(prefixes []string) bool {
	for i := range prefixes {
		if strings.HasPrefix(str.Raw, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (str *strNub) HasAnySuffix(suffixes []string) bool {
	for i := range suffixes {
		if strings.HasSuffix(str.Raw, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (str *strNub) HasPrefix(prefix string) bool {
	return strings.HasPrefix(str.Raw, prefix)
}

// HasSuffix checks if the string has the given suffix
func (str *strNub) HasSuffix(prefix string) bool {
	return strings.HasSuffix(str.Raw, prefix)
}

// Split creates a new nub from the split string
func (str *strNub) Split(delim string) *strSliceNub {
	return StrSlice(strings.Split(str.Raw, delim))
}
