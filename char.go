package n

// Char wraps the Go rune providing a way to distinguish it from an int32
// where as a rune is indistinguishable from an int32. Provides convenience
// methods on par with other rapid development languages.
type Char rune

// // C is an alias to NewChar for brevity
// func C(obj interface{}) *Str {
// 	return NewChar(obj)
// }

// NewChar creates a new chart from the given obj. Will always be non nil.
// Supports: string, *string, rune, *rune, byte, *byte
func NewChar(obj interface{}) *Char {
	return ToChar(obj)
}

// NewCharV creates a new chart from the given obj. Will always be non nil.
// Allows for empty Char with a Null value
func NewCharV(obj ...interface{}) *Char {
	new := Char(0)
	return &new
}

// Object interface methods
//--------------------------------------------------------------------------------------------------

// A is an alias of String for brevity
func (p *Char) A() string {
	return p.String()
}

// Equal returns true if the given *Char is value equal to this *Char.
func (p *Char) Equal(obj interface{}) bool {
	other := ToChar(obj)
	if p == nil {
		return false
	}
	return *p == *other
}

// G returns the underlying data structure as a builtin Go type
func (p *Char) G() rune {
	return p.O().(rune)
}

// Less returns true if the given *Char is less than this *Char.
func (p *Char) Less(obj interface{}) bool {
	other := ToChar(obj)
	if p == nil {
		return false
	}
	return p.A() < other.A()
}

// O returns the underlying data structure as is
func (p *Char) O() interface{} {
	if p == nil {
		return rune(0)
	}
	return rune(*p)
}

// Nil tests if the object is nil
func (p *Char) Nil() bool {
	if p == nil {
		return true
	}
	return false
}

// Null tests if the char is a rune(0)
func (p *Char) Null() bool {
	if p == nil {
		return false
	}
	return rune(*p) == rune(0)
}

// String returns a string representation of the Object, implements Stringer interface.
func (p *Char) String() string {
	if p == nil || *p == Char(0) {
		return ""
	}
	return string(*p)
}
