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

// Split creates a new nub from the split string
func (str *StrNub) Split(delim string) *StrSliceNub {
	return StrSlice(strings.Split(str.Raw, delim))
}
