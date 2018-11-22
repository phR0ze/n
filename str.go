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

// Split creates a new nub from the split string
func (str *StrNub) Split(delim string) *StrSliceNub {
	return StrSlice(strings.Split(str.Raw, delim))
}
