package nub

import (
	"reflect"
	"strings"
)

// A provides a new empty Queryable string
func A() *Queryable {
	obj := string("")
	ref := reflect.ValueOf(obj)
	return &Queryable{O: obj, ref: &ref, Iter: strIter(ref, obj)}
}

func strIter(ref reflect.Value, obj interface{}) func() Iterator {
	return func() Iterator {
		i := 0
		len := ref.Len()
		return func() (item interface{}, ok bool) {
			if ok = i < len; ok {
				item = ref.Index(i).Interface()
				i++
			}
			return
		}
	}
}

// Str materializes the result into a string
func (q *Queryable) Str() string {
	return q.O.(string)
}

// Strs materializes the results into a string slice
func (q *Queryable) Strs() (result []string) {
	q.Each(func(item interface{}) {
		result = append(result, item.(string))
	})
	return
}

// TODO: Need to refactor below here
type strNub struct {
	raw string
}

// NewStr creates a new nub
func NewStr() *strNub {
	return Str("")
}

// Str creates a new nub from the given string
func Str(slice string) *strNub {
	return &strNub{raw: slice}
}

// Contains checks if the given target is contained in this string
func (str *strNub) Contains(target string) bool {
	return strings.Contains(str.raw, target)
}

// ContainsAny checks if any of the targets are contained in this string
func (str *strNub) ContainsAny(targets []string) bool {
	for i := range targets {
		if strings.Contains(str.raw, targets[i]) {
			return true
		}
	}
	return false
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (str *strNub) HasAnyPrefix(prefixes []string) bool {
	for i := range prefixes {
		if strings.HasPrefix(str.raw, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (str *strNub) HasAnySuffix(suffixes []string) bool {
	for i := range suffixes {
		if strings.HasSuffix(str.raw, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (str *strNub) HasPrefix(prefix string) bool {
	return strings.HasPrefix(str.raw, prefix)
}

// HasSuffix checks if the string has the given suffix
func (str *strNub) HasSuffix(prefix string) bool {
	return strings.HasSuffix(str.raw, prefix)
}

// Split creates a new nub from the split string
func (str *strNub) Split(delim string) *strSliceNub {
	return StrSlice(strings.Split(str.raw, delim))
}

// M materializes object invoking deferred execution
func (str *strNub) M() string {
	return str.raw
}
