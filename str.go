package nub

import "strings"

// StrNub ...
type StrNub struct {
	Raw string
}

// Str creates a new nub
func Str(slice string) *StrNub {
	return &StrNub{Raw: slice}
}

// Contains calls through to underlying string
func (str *StrNub) Contains(target string) bool {
	return strings.Contains(str.Raw, target)
}

// ContainsAny provides a reusable check for the given targets
func (str *StrNub) ContainsAny(targets []string) bool {
	for i := range targets {
		if strings.Contains(str.Raw, targets[i]) {
			return true
		}
	}
	return false
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (str *StrNub) HasAnyPrefix(prefixes []string) bool {
	for i := range prefixes {
		if strings.HasPrefix(str.Raw, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (str *StrNub) HasAnySuffix(suffixes []string) bool {
	for i := range suffixes {
		if strings.HasSuffix(str.Raw, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (str *StrNub) HasPrefix(prefix string) bool {
	return strings.HasPrefix(str.Raw, prefix)
}

// HasSuffix checks if the string has the given suffix
func (str *StrNub) HasSuffix(prefix string) bool {
	return strings.HasSuffix(str.Raw, prefix)
}

// Split creates a new nub from the split string
func (str *StrNub) Split(delim string) *StrSliceNub {
	return StrSlice(strings.Split(str.Raw, delim))
}
