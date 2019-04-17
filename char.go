package n

// Char wraps the Go rune providing a way to distinguish it from an int32
// where as a rune is indistinguishable from an int32. Provides convenience
// methods on par with other rapid development languages.
type Char rune

// // C is an alias to NewChar for brevity
// func C(obj interface{}) *Str {
// 	return NewChar(obj)
// }

// NewChar creates a new chart from the given obj
// Supports: string, *string, rune, *rune, byte, *byte
func NewChar(obj interface{}) *Char {
	str := ""
	var new Char
	switch x := obj.(type) {
	case nil:
	case Str, *Str:
		str = string(Indirect(x).(Str))
	case string, *string:
		str = Indirect(x).(string)
	case byte, *byte:
		str = string(Indirect(x).(byte))
	case []byte, *[]byte:
		str = string(Indirect(x).([]byte))
	case rune, *rune:
		str = string(Indirect(x).(rune))
	default:
		str = ToString(obj)
	}
	if len(str) > 0 {
		new = Char(str[0])
	}
	return &new
}

// Object interface methods
//--------------------------------------------------------------------------------------------------

// O returns the underlying data structure as is
func (p *Char) O() interface{} {
	if p == nil {
		return nil
	}
	return *p
}

// Nil tests if the object is nil
func (p *Char) Nil() bool {
	if p == nil {
		return true
	}
	return false
}

// String returns a string representation of the Object, implements Stringer interface.
func (p *Char) String() string {
	if p == nil || *p == Char(0) {
		return ""
	}
	return string(*p)
}
