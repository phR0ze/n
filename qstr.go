package n

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	// ReGraphicalOnly is a regex to filter on graphical runes only
	ReGraphicalOnly = regexp.MustCompile(`[^[:graph:]]+`)
)

// QStr implementes the Queryable Interface and integrates with other queryable types.
// It provides a plethora of convenience methods to work with string types.
type QStr struct {
	v string
}

// A instantiates a new QStr optionally seeding it with the given value
// which may be of any type and will be converted to a string.
func A(str interface{}) (result *QStr) {
	switch x := str.(type) {
	case *QStr:
		result = x
	case string:
		result = &QStr{v: x}
	case rune:
		result = &QStr{v: string(x)}
	case []rune:
		result = &QStr{v: string(x)}
	case []byte:
		result = &QStr{v: string(x)}
	case int:
		result = &QStr{v: strconv.Itoa(x)}
	default:
		result = &QStr{v: fmt.Sprintf("%v", x)}
	}
	return
}

// Type returns the identifier for this queryable type.
// Implements the queryable interface.
func (q *QStr) Type() QType {
	return QStrType
}

// O returns the underlying data structure
// Implements the queryable interface.
func (q *QStr) O() interface{} {
	return q.v
}

// A exports the QStr as an external string
func (q *QStr) A() string {
	return q.v
}

// B exports the QStr as an external []byte
func (q *QStr) B() []byte {
	return []byte(q.v)
}

// Q casts the QStr as a Queryable
func (q *QStr) Q() Queryable {
	return Queryable(q)
}

// At returns the rune at the given index location. Allows for negative notation
func (q *QStr) At(i int) rune {
	if i < 0 {
		i = len(q.v) + i
	}
	if i >= 0 && i < len(q.v) {
		return rune(q.v[i])
	}
	panic("Index out of QStr bounds")
}

// Contains checks if the given target is contained in this string
func (q *QStr) Contains(target string) bool {
	return strings.Contains(q.v, target)
}

// ContainsAll checks if all the given targets are contained in this string
func (q *QStr) ContainsAll(targets ...string) bool {
	for i := range targets {
		if !strings.Contains(q.v, targets[i]) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any of the targets are contained in this string
func (q *QStr) ContainsAny(targets ...string) bool {
	for i := range targets {
		if strings.Contains(q.v, targets[i]) {
			return true
		}
	}
	return false
}

// Empty returns true if the pointer is nil, string is empty or whitespace only
func (q *QStr) Empty() bool {
	return q == nil || q.TrimSpace().v == ""
}

// HasAnyPrefix checks if the string has any of the given prefixes
func (q *QStr) HasAnyPrefix(prefixes ...string) bool {
	for i := range prefixes {
		if strings.HasPrefix(q.v, prefixes[i]) {
			return true
		}
	}
	return false
}

// HasAnySuffix checks if the string has any of the given suffixes
func (q *QStr) HasAnySuffix(suffixes ...string) bool {
	for i := range suffixes {
		if strings.HasSuffix(q.v, suffixes[i]) {
			return true
		}
	}
	return false
}

// HasPrefix checks if the string has the given prefix
func (q *QStr) HasPrefix(prefix string) bool {
	return strings.HasPrefix(q.v, prefix)
}

// HasSuffix checks if the string has the given suffix
func (q *QStr) HasSuffix(suffix string) bool {
	return strings.HasSuffix(q.v, suffix)
}

// Len returns the length of the string
func (q *QStr) Len() int {
	return len(q.v)
}

// Nil tests if the queryable is nil.
// Implements Queryable interface
func (q *QStr) Nil() bool {
	if q == nil {
		return true
	}
	return false
}

// Replace wraps strings.Replace and allows for chaining and defaults
func (q *QStr) Replace(old, new string, ns ...int) *QStr {
	n := -1
	if len(ns) > 0 {
		n = ns[0]
	}
	return A(strings.Replace(q.v, old, new, n))
}

// SpaceLeft returns leading whitespace
func (q *QStr) SpaceLeft() *QStr {
	spaces := []rune{}
	for _, r := range q.v {
		if unicode.IsSpace(r) {
			spaces = append(spaces, r)
		} else {
			break
		}
	}
	return A(spaces)
}

// SpaceRight returns trailing whitespace
func (q *QStr) SpaceRight() *QStr {
	spaces := []rune{}
	for i := len(q.v) - 1; i > 0; i-- {
		if unicode.IsSpace(rune(q.v[i])) {
			spaces = append(spaces, rune(q.v[i]))
		} else {
			break
		}
	}
	return A(spaces)
}

// Split creates a new QSlice from the split string.
// Optional 'delim' defaults to space allows for changing the split delimiter.
func (q *QStr) Split(delim ...string) *QSlice {
	_delim := " "
	if len(delim) > 0 {
		_delim = delim[0]
	}
	return S(strings.Split(q.v, _delim))
}

// SplitOn splits the string on the first occurance of the delim.
// The delim is included in the first component.
func (q *QStr) SplitOn(delim string) (first, second string) {
	// if q.v != "" {
	// 	s := q.Split(delim)
	// 	if s.Len() > 0 {
	// 		first = s.First().A()
	// 		if strings.Contains(q.v, delim) {
	// 			first += delim
	// 		}
	// 	}
	// 	if s.Len() > 1 {
	// 		second = s.Slice(1, -1).Join(delim).A()
	// 	}
	// }
	return
}

// // ToASCII with given string
// func (q *strN) ToASCII() *strN {
// 	return A(ReGraphicalOnly.ReplaceAllString(q.v, " "))
// }

// // TrimPrefix trims the given prefix off the string
// func (q *strN) TrimPrefix(prefix string) *strN {
// 	return A(strings.TrimPrefix(q.v, prefix))
// }

// TrimSpace pass through to strings.TrimSpace
func (q *QStr) TrimSpace() *QStr {
	return A(strings.TrimSpace(q.v))
}

// // TrimSpaceLeft trims leading whitespace
// func (q *strN) TrimSpaceLeft() *strN {
// 	return A(strings.TrimLeftFunc(q.v, unicode.IsSpace))
// }

// // TrimSpaceRight trims trailing whitespace
// func (q *strN) TrimSpaceRight() *strN {
// 	return A(strings.TrimRightFunc(q.v, unicode.IsSpace))
// }

// // TrimSuffix trims the given suffix off the string
// func (q *strN) TrimSuffix(suffix string) *strN {
// 	return A(strings.TrimSuffix(q.v, suffix))
// }

// // YamlType converts the given string into a type expected in Yaml.
// // Quotes signifies a string.
// // No quotes signifies an int.
// // true or false signifies a bool.
// func (q *strN) YamlType() interface{} {
// 	if q.HasAnyPrefix("\"", "'") && q.HasAnySuffix("\"", "'") {
// 		return q.v[1 : len(q.v)-1]
// 	} else if q.v == "true" || q.v == "false" {
// 		if b, err := strconv.ParseBool(q.v); err == nil {
// 			return b
// 		}
// 	} else if f, err := strconv.ParseFloat(q.v, 32); err == nil {
// 		return f
// 	}
// 	return q.v
// }
