package n

import (
	// 	"fmt"
	// 	"regexp"
	// 	"strconv"
	// 	"strings"
	// 	"unicode"

	"strings"

	"github.com/pkg/errors"
)

// var (
// 	// ReGraphicalOnly is a regex to filter on graphical runes only
// 	ReGraphicalOnly = regexp.MustCompile(`[^[:graph:]]+`)
// )

// Str wraps the Go string to include convenience methods on par with other
// rapid development languages.
type Str string

// A is an alias to NewStr
func A(str interface{}) *Str {
	return NewStr(str)
}

// NewStr creates a new *Str which will never be nil
func NewStr(str interface{}) *Str {
	var new Str
	switch x := str.(type) {
	case Str:
		new = x
	case *Str:
		new = *x
	case rune:
		new = Str(x)
	case []rune:
		new = Str(x)
	default:
		new = Str(Obj(str).ToString())
	}
	return &new
}

// A exports the Str as a Go string
func (p *Str) A() string {
	if p == nil {
		return ""
	}
	return string(*p)
}

// AsciiOnly bool
// Equal(str *Str) bool
// Clear() *Str
// Concat(str *Str) *Str

// All checks if all the given strs are contained in this Str
func (p *Str) All(strs []string) bool {
	return p.AllV(strs...)
}

// AllV checks if all the given variadic strs are contained in this Str
func (p *Str) AllV(strs ...string) bool {
	if p == nil {
		return false
	}
	for i := range strs {
		if !strings.Contains(string(*p), strs[i]) {
			return false
		}
	}
	return true
}

// Any checks if any of the given strs are contained in this Str
func (p *Str) Any(strs []string) bool {
	return p.AnyV(strs...)
}

// AnyV checks if any of the given variadic strs are contained in this Str
func (p *Str) AnyV(strs ...string) bool {
	if p == nil {
		return false
	}
	for i := range strs {
		if strings.Contains(string(*p), strs[i]) {
			return true
		}
	}
	return false
}

// At returns the element at the given index location. Allows for negative notation.
func (p *Str) At(i int) rune {
	r, _ := p.AtE(i)
	return r
}

// AtE returns the element at the given index location. Allows for negative notation.
func (p *Str) AtE(i int) (r rune, err error) {
	if p == nil {
		err = errors.Errorf("Str is nil")
		return
	}
	if i = absIndex(len(*p), i); i == -1 {
		err = errors.Errorf("index out of Str bounds")
		return
	}
	r = rune(string(*p)[i])
	return
}

// B exports the Str as a Go []byte
func (p *Str) B() []byte {
	if p == nil {
		return []byte("")
	}
	return []byte(*p)
}

// Contains checks if the given str is contained in this Str
func (p *Str) Contains(str string) bool {
	if p == nil {
		return false
	}
	return strings.Contains(string(*p), str)
}

// func ContainsAny(s, chars string) bool
// func ContainsRune(s string, r rune) bool
// func Count(s, substr string) int
// func EqualFold(s, t string) bool
// func Fields(s string) []string
// func FieldsFunc(s string, f func(rune) bool) []string
// func HasPrefix(s, prefix string) bool
// func HasSuffix(s, suffix string) bool
// func Index(s, substr string) int
// func IndexAny(s, chars string) int
// func IndexByte(s string, c byte) int
// func IndexFunc(s string, f func(rune) bool) int
// func IndexRune(s string, r rune) int
// func Join(a []string, sep string) string
// func LastIndex(s, substr string) int
// func LastIndexAny(s, chars string) int
// func LastIndexByte(s string, c byte) int
// func LastIndexFunc(s string, f func(rune) bool) int
// func Map(mapping func(rune) rune, s string) string
// func Repeat(s string, count int) string
// func Replace(s, old, new string, n int) string
// func ReplaceAll(s, old, new string) string
// func Split(s, sep string) []string
// func SplitAfter(s, sep string) []string
// func SplitAfterN(s, sep string, n int) []string
// func SplitN(s, sep string, n int) []string
// func Title(s string) string
// func ToLower(s string) string
// func ToLowerSpecial(c unicode.SpecialCase, s string) string
// func ToTitle(s string) string
// func ToTitleSpecial(c unicode.SpecialCase, s string) string
// func ToUpper(s string) string
// func ToUpperSpecial(c unicode.SpecialCase, s string) string
// func Trim(s string, cutset string) string
// func TrimFunc(s string, f func(rune) bool) string
// func TrimLeft(s string, cutset string) string
// func TrimLeftFunc(s string, f func(rune) bool) string
// func TrimPrefix(s, prefix string) string
// func TrimRight(s string, cutset string) string
// func TrimRightFunc(s string, f func(rune) bool) string
// func TrimSpace(s string) string
// func TrimSuffix(s, suffix string) string

// // Empty returns true if the pointer is nil, string is empty or whitespace only
// func (q *QStr) Empty() bool {
// 	return q == nil || q.TrimSpace().v == ""
// }

// // HasAnyPrefix checks if the string has any of the given prefixes
// func (q *QStr) HasAnyPrefix(prefixes ...string) bool {
// 	for i := range prefixes {
// 		if strings.HasPrefix(q.v, prefixes[i]) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // HasAnySuffix checks if the string has any of the given suffixes
// func (q *QStr) HasAnySuffix(suffixes ...string) bool {
// 	for i := range suffixes {
// 		if strings.HasSuffix(q.v, suffixes[i]) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // HasPrefix checks if the string has the given prefix
// func (q *QStr) HasPrefix(prefix string) bool {
// 	return strings.HasPrefix(q.v, prefix)
// }

// // HasSuffix checks if the string has the given suffix
// func (q *QStr) HasSuffix(suffix string) bool {
// 	return strings.HasSuffix(q.v, suffix)
// }

// // Len returns the length of the string
// func (q *QStr) Len() int {
// 	return len(q.v)
// }

// // Nil tests if the numerable is nil.
// // Implements Numerable interface
// func (q *QStr) Nil() bool {
// 	if q == nil {
// 		return true
// 	}
// 	return false
// }

// O returns the underlying data structure as is
func (p *Str) O() interface{} {
	if p == nil {
		return ""
	}
	return string(*p)
}

// // Replace wraps strings.Replace and allows for chaining and defaults
// func (q *QStr) Replace(old, new string, ns ...int) *QStr {
// 	n := -1
// 	if len(ns) > 0 {
// 		n = ns[0]
// 	}
// 	return A(strings.Replace(q.v, old, new, n))
// }

// // SpaceLeft returns leading whitespace
// func (q *QStr) SpaceLeft() *QStr {
// 	spaces := []rune{}
// 	for _, r := range q.v {
// 		if unicode.IsSpace(r) {
// 			spaces = append(spaces, r)
// 		} else {
// 			break
// 		}
// 	}
// 	return A(spaces)
// }

// // SpaceRight returns trailing whitespace
// func (q *QStr) SpaceRight() *QStr {
// 	spaces := []rune{}
// 	for i := len(q.v) - 1; i > 0; i-- {
// 		if unicode.IsSpace(rune(q.v[i])) {
// 			spaces = append(spaces, rune(q.v[i]))
// 		} else {
// 			break
// 		}
// 	}
// 	return A(spaces)
// }

// // // Split creates a new NSlice from the split string.
// // // Optional 'delim' defaults to space allows for changing the split delimiter.
// // func (q *QStr) Split(delim ...string) *NSlice {
// // 	_delim := " "
// // 	if len(delim) > 0 {
// // 		_delim = delim[0]
// // 	}
// // 	return S(strings.Split(q.v, _delim))
// // }

// // SplitOn splits the string on the first occurance of the delim.
// // The delim is included in the first component.
// func (q *QStr) SplitOn(delim string) (first, second string) {
// 	// if q.v != "" {
// 	// 	s := q.Split(delim)
// 	// 	if s.Len() > 0 {
// 	// 		first = s.First().A()
// 	// 		if strings.Contains(q.v, delim) {
// 	// 			first += delim
// 	// 		}
// 	// 	}
// 	// 	if s.Len() > 1 {
// 	// 		second = s.Slice(1, -1).Join(delim).A()
// 	// 	}
// 	// }
// 	return
// }

// String returns a string representation of this Slice, implements the Stringer interface
func (p *Str) String() string {
	return p.A()
}

// // // ToASCII with given string
// // func (q *strN) ToASCII() *strN {
// // 	return A(ReGraphicalOnly.ReplaceAllString(q.v, " "))
// // }

// // // TrimPrefix trims the given prefix off the string
// // func (q *strN) TrimPrefix(prefix string) *strN {
// // 	return A(strings.TrimPrefix(q.v, prefix))
// // }

// // TrimSpace pass through to strings.TrimSpace
// func (q *QStr) TrimSpace() *QStr {
// 	return A(strings.TrimSpace(q.v))
// }

// // // TrimSpaceLeft trims leading whitespace
// // func (q *strN) TrimSpaceLeft() *strN {
// // 	return A(strings.TrimLeftFunc(q.v, unicode.IsSpace))
// // }

// // // TrimSpaceRight trims trailing whitespace
// // func (q *strN) TrimSpaceRight() *strN {
// // 	return A(strings.TrimRightFunc(q.v, unicode.IsSpace))
// // }

// // // TrimSuffix trims the given suffix off the string
// // func (q *strN) TrimSuffix(suffix string) *strN {
// // 	return A(strings.TrimSuffix(q.v, suffix))
// // }

// // // YamlType converts the given string into a type expected in Yaml.
// // // Quotes signifies a string.
// // // No quotes signifies an int.
// // // true or false signifies a bool.
// // func (q *strN) YamlType() interface{} {
// // 	if q.HasAnyPrefix("\"", "'") && q.HasAnySuffix("\"", "'") {
// // 		return q.v[1 : len(q.v)-1]
// // 	} else if q.v == "true" || q.v == "false" {
// // 		if b, err := strconv.ParseBool(q.v); err == nil {
// // 			return b
// // 		}
// // 	} else if f, err := strconv.ParseFloat(q.v, 32); err == nil {
// // 		return f
// // 	}
// // 	return q.v
// // }
