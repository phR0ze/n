package n

import (
	"time"

	"github.com/pkg/errors"
)

// Object is a wrapper around an interface{} value providing a number of export methods
// for casting and converting to other types via the excellent cast package.
type Object struct {
	o interface{} // value
}

// Obj creates a new Object from the given obj appending the Object methods
func Obj(obj interface{}) *Object {
	return &Object{obj}
}

// Object interface methods
//--------------------------------------------------------------------------------------------------

// O returns the underlying data structure as is
func (p *Object) O() interface{} {
	if p.Nil() {
		return nil
	}
	return p.o
}

// Nil tests if the numerable is nil
func (p *Object) Nil() bool {
	if p == nil || p.o == nil {
		return true
	}
	return false
}

// A is an alias to String for brevity
func (p *Object) A() string {
	return p.String()
}

// M is an alias to ToStringMap
func (p *Object) M() *StringMap {
	return p.ToStringMap()
}

// MG is an alias to ToStringMapG
func (p *Object) MG() map[string]interface{} {
	return p.ToStringMapG()
}

// S is an alias to ToStringSlice for brevity
func (p *Object) S() *StringSlice {
	return p.ToStringSlice()
}

// String returns a string representation of the Object, implements Stringer interface.
func (p *Object) String() string {
	if p == nil {
		return ""
	}
	return ToString(p.o)
}

// Bool
//--------------------------------------------------------------------------------------------------

// ToBool converts an interface to a bool type.
func (p *Object) ToBool() bool {
	if p == nil {
		return false
	}
	v, _ := ToBoolE(p.o)
	return v
}

// ToBoolE converts an interface to a bool type.
func (p *Object) ToBoolE() (bool, error) {
	if p == nil {
		return false, nil
	}
	return ToBoolE(p.o)
}

// Char
//--------------------------------------------------------------------------------------------------

// C is an alias to ToChar for brevity
func (p *Object) C() *Char {
	return p.ToChar()
}

// ToChar converts an interface to a *Char type.
func (p *Object) ToChar() *Char {
	if p == nil {
		val := Char(0)
		return &val
	}
	return ToChar(p.o)
}

// R is an alias to ToRune for brevity
func (p *Object) R() rune {
	return p.ToRune()
}

// ToRune converts an interface to a rune type.
func (p *Object) ToRune() rune {
	if p == nil {
		return rune(0)
	}
	return rune(*ToChar(p.o))
}

// Map
//--------------------------------------------------------------------------------------------------

// Query an object if it is a StringMap type
func (p *Object) Query(key string) *Object {
	obj, _ := p.QueryE(key)
	return obj
}

// QueryE an object if it is a StringMap type
func (p *Object) QueryE(key string) (obj *Object, err error) {
	obj = &Object{}
	if p == nil {
		return
	}
	var m *StringMap
	if m, err = ToStringMapE(p.o); err != nil {
		return
	}
	if obj, err = m.QueryE(key); err != nil {
		return
	}
	if obj.o == nil {
		err = errors.Errorf("invalid key")
	}
	return
}

// Time related
//--------------------------------------------------------------------------------------------------

// ToTime converts an interface to a time.Time type.
func (p *Object) ToTime() time.Time {
	v, _ := ToTimeE(p.o)
	return v
}

// ToTimeE converts an interface to a time.Time type.
func (p *Object) ToTimeE() (time.Time, error) {
	return ToTimeE(p.o)
}

// ToDuration converts an interface to a time.Duration type.
func (p *Object) ToDuration() time.Duration {
	v, _ := ToDurationE(p.o)
	return v
}

// ToDurationE converts an interface to a time.Duration type.
func (p *Object) ToDurationE() (time.Duration, error) {
	return ToDurationE(p.o)
}

// Float related
//--------------------------------------------------------------------------------------------------

// ToFloat32 converts an interface to a float32 type.
func (p *Object) ToFloat32() float32 {
	if p == nil {
		return float32(0)
	}
	return ToFloat32(p.o)
}

// ToFloat32E converts an interface to a float32 type.
func (p *Object) ToFloat32E() (float32, error) {
	if p == nil {
		return float32(0), nil
	}
	return ToFloat32E(p.o)
}

// ToFloat64 converts an interface to a float64 type.
func (p *Object) ToFloat64() float64 {
	if p == nil {
		return float64(0)
	}
	return ToFloat64(p.o)
}

// ToFloat64E converts an interface to a float64 type.
func (p *Object) ToFloat64E() (float64, error) {
	if p == nil {
		return float64(0), nil
	}
	return ToFloat64E(p.o)
}

// Int related
//--------------------------------------------------------------------------------------------------

// ToInt converts an interface to an int type.
func (p *Object) ToInt() int {
	if p == nil {
		return 0
	}
	v, _ := ToIntE(p.o)
	return v
}

// ToIntE converts an interface to an int type.
func (p *Object) ToIntE() (int, error) {
	if p == nil {
		return 0, nil
	}
	return ToIntE(p.o)
}

// ToInt8 converts an interface to an int8 type.
func (p *Object) ToInt8() int8 {
	if p == nil {
		return 0
	}
	v, _ := ToInt8E(p.o)
	return v
}

// ToInt8E converts an interface to an int8 type.
func (p *Object) ToInt8E() (int8, error) {
	if p == nil {
		return 0, nil
	}
	return ToInt8E(p.o)
}

// ToInt16 converts an interface to an int16 type.
func (p *Object) ToInt16() int16 {
	if p == nil {
		return 0
	}
	v, _ := ToInt16E(p.o)
	return v
}

// ToInt16E converts an interface to an int16 type.
func (p *Object) ToInt16E() (int16, error) {
	if p == nil {
		return 0, nil
	}
	return ToInt16E(p.o)
}

// ToInt32 converts an interface to an int32 type.
func (p *Object) ToInt32() int32 {
	if p == nil {
		return 0
	}
	v, _ := ToInt32E(p.o)
	return v
}

// ToInt32E converts an interface to an int32 type.
func (p *Object) ToInt32E() (int32, error) {
	if p == nil {
		return 0, nil
	}
	return ToInt32E(p.o)
}

// ToInt64 converts an interface to an int64 type.
func (p *Object) ToInt64() int64 {
	if p == nil {
		return 0
	}
	v, _ := ToInt64E(p.o)
	return v
}

// ToInt64E converts an interface to an int64 type.
func (p *Object) ToInt64E() (int64, error) {
	if p == nil {
		return 0, nil
	}
	return ToInt64E(p.o)
}

// ToUint converts an interface to a uint type.
func (p *Object) ToUint() uint {
	if p == nil {
		return 0
	}
	v, _ := ToUintE(p.o)
	return v
}

// ToUintE converts an interface to a uint type.
func (p *Object) ToUintE() (uint, error) {
	if p == nil {
		return 0, nil
	}
	return ToUintE(p.o)
}

// ToUint8 converts an interface to a uint8 type.
func (p *Object) ToUint8() uint8 {
	if p == nil {
		return 0
	}
	v, _ := ToUint8E(p.o)
	return v
}

// ToUint8E converts an interface to a uint8 type.
func (p *Object) ToUint8E() (uint8, error) {
	if p == nil {
		return 0, nil
	}
	return ToUint8E(p.o)
}

// ToUint16 converts an interface to a uint16 type.
func (p *Object) ToUint16() uint16 {
	if p == nil {
		return 0
	}
	v, _ := ToUint16E(p.o)
	return v
}

// ToUint16E converts an interface to a uint16 type.
func (p *Object) ToUint16E() (uint16, error) {
	if p == nil {
		return 0, nil
	}
	return ToUint16E(p.o)
}

// ToUint32 converts an interface to a uint32 type.
func (p *Object) ToUint32() uint32 {
	if p == nil {
		return 0
	}
	v, _ := ToUint32E(p.o)
	return v
}

// ToUint32E converts an interface to a uint32 type.
func (p *Object) ToUint32E() (uint32, error) {
	if p == nil {
		return 0, nil
	}
	return ToUint32E(p.o)
}

// ToUint64 converts an interface to a uint64 type.
func (p *Object) ToUint64() uint64 {
	if p == nil {
		return 0
	}
	v, _ := ToUint64E(p.o)
	return v
}

// ToUint64E converts an interface to a uint64 type.
func (p *Object) ToUint64E() (uint64, error) {
	if p == nil {
		return 0, nil
	}
	return ToUint64E(p.o)
}

// String related
//--------------------------------------------------------------------------------------------------

// ToStr converts object into a *Str
func (p *Object) ToStr() *Str {
	if p == nil {
		return ToStr(p)
	}
	return ToStr(p.o)
}

// ToString converts an interface to a string type.
func (p *Object) ToString() string {
	if p == nil {
		return ""
	}
	return p.String()
}

// ToStringE converts an interface to a string type.
func (p *Object) ToStringE() (string, error) {
	if p == nil {
		return "", nil
	}
	return p.ToString(), nil
}

// Map related
//--------------------------------------------------------------------------------------------------

// ToStringMap converts an interface to a *StringMap type.
func (p *Object) ToStringMap() *StringMap {
	if p == nil {
		return NewStringMapV()
	}
	return ToStringMap(p.o)
}

// ToStringMapE converts an interface to a *StringMap type.
func (p *Object) ToStringMapE() (*StringMap, error) {
	if p == nil {
		return NewStringMapV(), nil
	}
	return ToStringMapE(p.o)
}

// ToStringMapG converts an interface to a map[string]interface{} type.
func (p *Object) ToStringMapG() map[string]interface{} {
	if p == nil {
		return map[string]interface{}{}
	}
	v, _ := ToStringMapE(p.o)
	return v.G()
}

// ToStringMapGE converts an interface to a map[string]interface{} type.
func (p *Object) ToStringMapGE() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if p == nil {
		return result, nil
	}
	v, err := ToStringMapE(p.o)
	if err != nil {
		return result, err
	}
	return v.G(), err
}

// Slice related
//--------------------------------------------------------------------------------------------------

// ToMapSlice converts an interface to a *MapSlice type.
func (p *Object) ToMapSlice() *MapSlice {
	if p == nil {
		return NewMapSliceV()
	}
	return ToMapSlice(p.o)
}

// ToMapSliceE converts an interface to a *MapSlice type.
func (p *Object) ToMapSliceE() (*MapSlice, error) {
	if p == nil {
		return NewMapSliceV(), nil
	}
	return ToMapSliceE(p.o)
}

// ToMapSliceG converts an interface to a []map[string]interface{} type.
func (p *Object) ToMapSliceG() []map[string]interface{} {
	if p == nil {
		return NewMapSliceV().G()
	}
	return ToMapSlice(p.o).G()
}

// ToMapSliceGE converts an interface to a []map[string]interface{} type.
func (p *Object) ToMapSliceGE() ([]map[string]interface{}, error) {
	if p == nil {
		return NewMapSliceV().G(), nil
	}
	m, err := ToMapSliceE(p.o)
	if err != nil {
		return nil, err
	}
	return m.G(), nil
}

// ToStringSlice converts an interface to a *StringSlice type.
func (p *Object) ToStringSlice() *StringSlice {
	if p == nil {
		return NewStringSliceV()
	}
	return ToStringSlice(p.o)
}

// ToStringSliceE converts an interface to a *StringSlice type.
func (p *Object) ToStringSliceE() (*StringSlice, error) {
	if p == nil {
		return NewStringSliceV(), nil
	}
	return ToStringSliceE(p.o)
}

// ToStrs converts an interface to a []string type.
func (p *Object) ToStrs() []string {
	return p.ToStringSlice().G()
}

// ToStrsE converts an interface to a []string type.
func (p *Object) ToStrsE() ([]string, error) {
	result, err := p.ToStringSliceE()
	if err != nil {
		return []string{}, err
	}
	return result.G(), nil
}

// ToIntSlice converts an interface to a *IntSlice type.
func (p *Object) ToIntSlice() *IntSlice {
	if p == nil {
		return NewIntSliceV()
	}
	v, _ := ToIntSliceE(p.o)
	return v
}

// ToIntSliceE converts an interface to a *IntSlice type.
func (p *Object) ToIntSliceE() (*IntSlice, error) {
	if p == nil {
		return NewIntSliceV(), nil
	}
	return ToIntSliceE(p.o)
}

// ToIntSliceG converts an interface to a []int type.
func (p *Object) ToIntSliceG() []int {
	result := []int{}
	if p == nil {
		return result
	}
	v, err := ToIntSliceE(p.o)
	if err != nil {
		return result
	}
	return v.G()
}

// ToIntSliceGE converts an interface to a []int type.
func (p *Object) ToIntSliceGE() ([]int, error) {
	result := []int{}
	if p == nil {
		return result, nil
	}
	v, err := ToIntSliceE(p.o)
	if err != nil {
		return result, err
	}
	return v.G(), err
}
