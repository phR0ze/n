package n

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

// DeReference dereferences the interface if needed returning a non-pointer type
func DeReference(obj interface{}) interface{} {
	switch x := obj.(type) {
	case nil:
		return x

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		return x
	case *bool:
		if x == nil {
			return false
		}
		return *x
	case []bool:
		return x
	case []*bool:
		return x
	case *[]bool:
		if x == nil {
			return []bool{}
		}
		return *x
	case *[]*bool:
		if x == nil {
			return []*bool{}
		}
		return *x

	// byte
	//----------------------------------------------------------------------------------------------
	// byte is just a uint8
	// case byte:
	// 	return x
	// case *byte:
	// 	if x == nil {
	// 		return byte(0)
	// 	}
	// 	return *x
	// []byte is just a uint8[] which is defined later
	// case []byte:
	// 	return x
	// case *[]byte:
	// 	if x == nil {
	// 		return []byte{}
	// 	}
	// 	return *x

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		return x
	case *Char:
		if x == nil {
			return *NewChar("")
		}
		return *x
	case []Char:
		return x
	case []*Char:
		return x
	case *[]Char:
		if x == nil {
			return []Char{}
		}
		return *x
	case *[]*Char:
		if x == nil {
			return []*Char{}
		}
		return *x

	// error
	//----------------------------------------------------------------------------------------------
	case error:
		return x

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		return x
	case *float32:
		if x == nil {
			return float32(0)
		}
		return *x
	case []float32:
		return x
	case []*float32:
		return x
	case *[]float32:
		if x == nil {
			return []float32{}
		}
		return *x
	case *[]*float32:
		if x == nil {
			return []*float32{}
		}
		return *x

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		return x
	case *float64:
		if x == nil {
			return float64(0)
		}
		return *x
	case []float64:
		return x
	case []*float64:
		return x
	case *[]float64:
		if x == nil {
			return []float64{}
		}
		return *x
	case *[]*float64:
		if x == nil {
			return []*float64{}
		}
		return *x

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		return x
	case []*interface{}:
		return x
	case *[]interface{}:
		if x == nil {
			return []interface{}{}
		}
		return *x
	case *[]*interface{}:
		if x == nil {
			return []*interface{}{}
		}
		return *x

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		return x
	case *int:
		if x == nil {
			return 0
		}
		return *x
	case []int:
		return x
	case []*int:
		return x
	case *[]int:
		if x == nil {
			return []int{}
		}
		return *x
	case *[]*int:
		if x == nil {
			return []*int{}
		}
		return *x

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		return x
	case *int8:
		if x == nil {
			return int8(0)
		}
		return *x
	case []int8:
		return x
	case []*int8:
		return x
	case *[]int8:
		if x == nil {
			return []int8{}
		}
		return *x
	case *[]*int8:
		if x == nil {
			return []*int8{}
		}
		return *x

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		return x
	case *int16:
		if x == nil {
			return int16(0)
		}
		return *x
	case []int16:
		return x
	case []*int16:
		return x
	case *[]int16:
		if x == nil {
			return []int16{}
		}
		return *x
	case *[]*int16:
		if x == nil {
			return []*int16{}
		}
		return *x

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		return x
	case *int32:
		if x == nil {
			return int32(0)
		}
		return *x
	case []int32:
		return x
	case []*int32:
		return x
	case *[]int32:
		if x == nil {
			return []int32{}
		}
		return *x
	case *[]*int32:
		if x == nil {
			return []*int32{}
		}
		return *x

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		return x
	case *int64:
		if x == nil {
			return int64(0)
		}
		return *x
	case []int64:
		return x
	case []*int64:
		return x
	case *[]int64:
		if x == nil {
			return []int64{}
		}
		return *x
	case *[]*int64:
		if x == nil {
			return []*int64{}
		}
		return *x

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		return x
	case *IntSlice:
		if x == nil {
			return *NewIntSliceV()
		}
		return *x
	case []IntSlice:
		return x
	case []*IntSlice:
		return x
	case *[]IntSlice:
		if x == nil {
			return []IntSlice{}
		}
		return *x
	case *[]*IntSlice:
		if x == nil {
			return []*IntSlice{}
		}
		return *x

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		return x
	case *Object:
		if x == nil {
			return Object{}
		}
		return *x
	case []Object:
		return x
	case []*Object:
		return x
	case *[]Object:
		if x == nil {
			return []Object{}
		}
		return *x
	case *[]*Object:
		if x == nil {
			return []*Object{}
		}
		return *x

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		return x
	case *Str:
		if x == nil {
			return *NewStrV()
		}
		return *x
	case []Str:
		return x
	case []*Str:
		return x
	case *[]Str:
		if x == nil {
			return []Str{}
		}
		return *x
	case *[]*Str:
		if x == nil {
			return []*Str{}
		}
		return *x

	// rune
	//----------------------------------------------------------------------------------------------
	// rune is a int32 which is already defined
	// case rune:
	// 	return x
	// case *rune:
	// 	if x == nil {
	// 		return rune(0)
	// 	}

	// StringSlice
	//----------------------------------------------------------------------------------------------
	case StringSlice:
		return x
	case *StringSlice:
		if x == nil {
			return *NewStringSliceV()
		}
		return *x
	case []StringSlice:
		return x
	case []*StringSlice:
		return x
	case *[]StringSlice:
		if x == nil {
			return []StringSlice{}
		}
		return *x
	case *[]*StringSlice:
		if x == nil {
			return []*StringSlice{}
		}
		return *x

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		return x
	case *string:
		if x == nil {
			return ""
		}
		return *x
	case []string:
		return x
	case []*string:
		return x
	case *[]string:
		if x == nil {
			return []string{}
		}
		return *x
	case *[]*string:
		if x == nil {
			return []*string{}
		}
		return *x

	// template.CSS
	//----------------------------------------------------------------------------------------------
	case template.CSS:
		return x
	case *template.CSS:
		if x == nil {
			return template.CSS("")
		}
		return *x
	case []template.CSS:
		return x
	case []*template.CSS:
		return x
	case *[]template.CSS:
		if x == nil {
			return []template.CSS{}
		}
		return *x
	case *[]*template.CSS:
		if x == nil {
			return []*template.CSS{}
		}
		return *x

	// template.HTML
	//----------------------------------------------------------------------------------------------
	case template.HTML:
		return x
	case *template.HTML:
		if x == nil {
			return template.HTML("")
		}
		return *x
	case []template.HTML:
		return x
	case []*template.HTML:
		return x
	case *[]template.HTML:
		if x == nil {
			return []template.HTML{}
		}
		return *x
	case *[]*template.HTML:
		if x == nil {
			return []*template.HTML{}
		}
		return *x

	// template.HTMLAttr
	//----------------------------------------------------------------------------------------------
	case template.HTMLAttr:
		return x
	case *template.HTMLAttr:
		if x == nil {
			return template.HTMLAttr("")
		}
		return *x
	case []template.HTMLAttr:
		return x
	case []*template.HTMLAttr:
		return x
	case *[]template.HTMLAttr:
		if x == nil {
			return []template.HTMLAttr{}
		}
		return *x
	case *[]*template.HTMLAttr:
		if x == nil {
			return []*template.HTMLAttr{}
		}
		return *x

	// template.JS
	//----------------------------------------------------------------------------------------------
	case template.JS:
		return x
	case *template.JS:
		if x == nil {
			return template.JS("")
		}
		return *x
	case []template.JS:
		return x
	case []*template.JS:
		return x
	case *[]template.JS:
		if x == nil {
			return []template.JS{}
		}
		return *x
	case *[]*template.JS:
		if x == nil {
			return []*template.JS{}
		}
		return *x

	// template.URL
	//----------------------------------------------------------------------------------------------
	case template.URL:
		return x
	case *template.URL:
		if x == nil {
			return template.URL("")
		}
		return *x
	case []template.URL:
		return x
	case []*template.URL:
		return x
	case *[]template.URL:
		if x == nil {
			return []template.URL{}
		}
		return *x
	case *[]*template.URL:
		if x == nil {
			return []*template.URL{}
		}
		return *x

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		return x
	case *uint:
		if x == nil {
			return uint(0)
		}
		return *x
	case []uint:
		return x
	case []*uint:
		return x
	case *[]uint:
		if x == nil {
			return []uint{}
		}
		return *x
	case *[]*uint:
		if x == nil {
			return []*uint{}
		}
		return *x

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		return x
	case *uint8:
		if x == nil {
			return uint8(0)
		}
		return *x
	case []uint8:
		return x
	case []*uint8:
		return x
	case *[]uint8:
		if x == nil {
			return []uint8{}
		}
		return *x
	case *[]*uint8:
		if x == nil {
			return []*uint8{}
		}
		return *x

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		return x
	case *uint16:
		if x == nil {
			return uint16(0)
		}
		return *x
	case []uint16:
		return x
	case []*uint16:
		return x
	case *[]uint16:
		if x == nil {
			return []uint16{}
		}
		return *x
	case *[]*uint16:
		if x == nil {
			return []*uint16{}
		}
		return *x

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		return x
	case *uint32:
		if x == nil {
			return uint32(0)
		}
		return *x
	case []uint32:
		return x
	case []*uint32:
		return x
	case *[]uint32:
		if x == nil {
			return []uint32{}
		}
		return *x
	case *[]*uint32:
		if x == nil {
			return []*uint32{}
		}
		return *x

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		return x
	case *uint64:
		if x == nil {
			return uint64(0)
		}
		return *x
	case []uint64:
		return x
	case []*uint64:
		return x
	case *[]uint64:
		if x == nil {
			return []uint64{}
		}
		return *x
	case *[]*uint64:
		if x == nil {
			return []*uint64{}
		}
		return *x

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		return reflect.Indirect(reflect.ValueOf(obj)).Interface()
	}
}

// Reference converts the interface to a pointer type. Unlike DeReference nil may be returned as
// all pointers are simply passed through. This allows for custom default handling for callers.
func Reference(obj interface{}) interface{} {
	switch x := obj.(type) {
	case nil:
		return x

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		return &x
	case *bool:
		return x
	case []bool:
		return &x
	case []*bool:
		return &x
	case *[]bool:
		return x
	case *[]*bool:
		return x

	// byte
	//----------------------------------------------------------------------------------------------
	// byte is just a uint8
	// []byte is just a uint8[] which is defined later

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		return &x
	case *Char:
		return x
	case []Char:
		return &x
	case []*Char:
		return &x
	case *[]Char:
		return x
	case *[]*Char:
		return x

	// error
	//----------------------------------------------------------------------------------------------
	case error:
		return x

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		return &x
	case *float32:
		return x
	case []float32:
		return &x
	case []*float32:
		return &x
	case *[]float32:
		return x
	case *[]*float32:
		return x

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		return &x
	case *float64:
		return x
	case []float64:
		return &x
	case []*float64:
		return &x
	case *[]float64:
		return x
	case *[]*float64:
		return x

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		return &x
	case []*interface{}:
		return &x
	case *[]interface{}:
		return x
	case *[]*interface{}:
		return x

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		return &x
	case *int:
		return x
	case []int:
		return &x
	case []*int:
		return &x
	case *[]int:
		return x
	case *[]*int:
		return x

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		return &x
	case *int8:
		return x
	case []int8:
		return &x
	case []*int8:
		return &x
	case *[]int8:
		return x
	case *[]*int8:
		return x

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		return &x
	case *int16:
		return x
	case []int16:
		return &x
	case []*int16:
		return &x
	case *[]int16:
		return x
	case *[]*int16:
		return x

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		return &x
	case *int32:
		return x
	case []int32:
		return &x
	case []*int32:
		return &x
	case *[]int32:
		return x
	case *[]*int32:
		return x

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		return &x
	case *int64:
		return x
	case []int64:
		return &x
	case []*int64:
		return &x
	case *[]int64:
		return x
	case *[]*int64:
		return x

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		return &x
	case *IntSlice:
		return x
	case []IntSlice:
		return &x
	case []*IntSlice:
		return &x
	case *[]IntSlice:
		return x
	case *[]*IntSlice:
		return x

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		return &x
	case *Object:
		return x
	case []Object:
		return &x
	case []*Object:
		return &x
	case *[]Object:
		return x
	case *[]*Object:
		return x

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		return &x
	case *Str:
		return x
	case []Str:
		return &x
	case []*Str:
		return &x
	case *[]Str:
		return x
	case *[]*Str:
		return x

	// rune
	//----------------------------------------------------------------------------------------------
	// rune is a int32 which is already defined

	// StringSlice
	//----------------------------------------------------------------------------------------------
	case StringSlice:
		return &x
	case *StringSlice:
		return x
	case []StringSlice:
		return &x
	case []*StringSlice:
		return &x
	case *[]StringSlice:
		return x
	case *[]*StringSlice:
		return x

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		return &x
	case *string:
		return x
	case []string:
		return &x
	case []*string:
		return &x
	case *[]string:
		return x
	case *[]*string:
		return x

	// template.CSS
	//----------------------------------------------------------------------------------------------
	case template.CSS:
		return &x
	case *template.CSS:
		return x
	case []template.CSS:
		return &x
	case []*template.CSS:
		return &x
	case *[]template.CSS:
		return x
	case *[]*template.CSS:
		return x

	// template.HTML
	//----------------------------------------------------------------------------------------------
	case template.HTML:
		return &x
	case *template.HTML:
		return x
	case []template.HTML:
		return &x
	case []*template.HTML:
		return &x
	case *[]template.HTML:
		return x
	case *[]*template.HTML:
		return x

	// template.HTMLAttr
	//----------------------------------------------------------------------------------------------
	case template.HTMLAttr:
		return &x
	case *template.HTMLAttr:
		return x
	case []template.HTMLAttr:
		return &x
	case []*template.HTMLAttr:
		return &x
	case *[]template.HTMLAttr:
		return x
	case *[]*template.HTMLAttr:
		return x

	// template.JS
	//----------------------------------------------------------------------------------------------
	case template.JS:
		return &x
	case *template.JS:
		return x
	case []template.JS:
		return &x
	case []*template.JS:
		return &x
	case *[]template.JS:
		return x
	case *[]*template.JS:
		return x

	// template.URL
	//----------------------------------------------------------------------------------------------
	case template.URL:
		return &x
	case *template.URL:
		return x
	case []template.URL:
		return &x
	case []*template.URL:
		return &x
	case *[]template.URL:
		return x
	case *[]*template.URL:
		return x

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		return &x
	case *uint:
		return x
	case []uint:
		return &x
	case []*uint:
		return &x
	case *[]uint:
		return x
	case *[]*uint:
		return x

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		return &x
	case *uint8:
		return x
	case []uint8:
		return &x
	case []*uint8:
		return &x
	case *[]uint8:
		return x
	case *[]*uint8:
		return x

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		return &x
	case *uint16:
		return x
	case []uint16:
		return &x
	case []*uint16:
		return &x
	case *[]uint16:
		return x
	case *[]*uint16:
		return x

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		return &x
	case *uint32:
		return x
	case []uint32:
		return &x
	case []*uint32:
		return &x
	case *[]uint32:
		return x
	case *[]*uint32:
		return x

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		return &x
	case *uint64:
		return x
	case []uint64:
		return &x
	case []*uint64:
		return &x
	case *[]uint64:
		return x
	case *[]*uint64:
		return x

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(x)
		if v.CanAddr() {
			if v.Kind() == reflect.Ptr {
				return x
			}
			return &x
		}
		return x
	}
}

// Convert functions
//--------------------------------------------------------------------------------------------------

// ToBool converts an interface to a bool type.
func ToBool(obj interface{}) bool {
	val, _ := ToBoolE(obj)
	return val
}

// ToBoolE converts an interface to a bool type.
func ToBoolE(obj interface{}) (val bool, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		val = x
	case int:
		val = x != 0
	case int8:
		val = x != 0
	case int16:
		val = x != 0
	case int32:
		val = x != 0
	case int64:
		val = x != 0
	case Object:
		val, err = x.ToBoolE()
	case Str:
		return ToBoolE(string(x))
	case string:
		if val, err = strconv.ParseBool(x); err != nil {
			err = errors.Wrapf(err, "failed to convert string to bool")
		}
	case uint:
		val = x != 0
	case uint8:
		val = x != 0
	case uint16:
		val = x != 0
	case uint32:
		val = x != 0
	case uint64:
		val = x != 0
	default:
		err = errors.Errorf("unable to convert type %T to bool", x)
	}
	return
}

// ToChar convert an interface to a Char type.
func ToChar(obj interface{}) *Char {
	val := Char(0)
	o := Reference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case *bool:
		if x != nil {
			v := 0
			if *x {
				v = 1
			}
			val = Char(([]rune(strconv.FormatInt(int64(v), 10)))[0])
		}

	// byte
	//----------------------------------------------------------------------------------------------
	case *[]byte:
		if x != nil {
			if len(*x) != 0 {
				v, _ := utf8.DecodeRune(*x)
				val = Char(v)
			}
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case *Char:
		if x != nil {
			return x
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case *float32:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case *float64:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// int
	//----------------------------------------------------------------------------------------------
	case *int:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case *int8:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case *int16:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// int32 i.e. rune
	//----------------------------------------------------------------------------------------------
	case *int32:
		if x != nil {
			val = Char(*x)
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case *int64:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(*x, 10)))[0])
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case *Object:
		if x != nil {
			val = *x.ToChar()
		}

	// *[]rune
	//----------------------------------------------------------------------------------------------
	case *[]rune:
		if x != nil && len(*x) != 0 {
			val = Char((*x)[0])
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case *Str:
		if x != nil && len(*x) != 0 {
			val = Char((*x)[0])
		}

	// string
	//----------------------------------------------------------------------------------------------
	case *string:
		if x != nil && len(*x) != 0 {
			v, _ := utf8.DecodeRuneInString(*x)
			val = Char(v)
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case *uint:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// uint8 i.e. byte handle differently
	//----------------------------------------------------------------------------------------------
	case *uint8:
		if x != nil && len(string(*x)) != 0 {
			val = Char(string(*x)[0])
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case *uint16:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case *uint32:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case *uint64:
		if x != nil {
			val = Char(([]rune(strconv.FormatInt(int64(*x), 10)))[0])
		}
	}
	return &val
}

// ToInt convert an interface to an int type.
func ToInt(obj interface{}) int {
	x, _ := ToIntE(obj)
	return x
}

// ToIntE convert an interface to an int type.
func ToIntE(obj interface{}) (val int, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToIntE(x.A())
	case float32:
		val = int(x)
	case float64:
		val = int(x)
	case int:
		val = x
	case int8:
		val = int(x)
	case int16:
		val = int(x)
	case int32:
		val = int(x)
	case int64:
		val = int(x)
	case Object:
		val, err = x.ToIntE()
	case uint:
		val = int(x)
	case uint8:
		val = int(x)
	case uint16:
		val = int(x)
	case uint32:
		val = int(x)
	case uint64:
		val = int(x)
	case Str:
		return ToIntE(string(x))
	case string:
		var v int64
		if v, err = strconv.ParseInt(x, 0, 0); err != nil {
			err = errors.Wrapf(err, "failed to convert string to int")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = int(v)
		}
	default:
		err = fmt.Errorf("unable to convert type %T to int", x)
	}
	return
}

// ToIntSlice convert an interface to a IntSlice type which will never be nil
func ToIntSlice(obj interface{}) *IntSlice {
	x, _ := ToIntSliceE(obj)
	if x == nil {
		return &IntSlice{}
	}
	return x
}

// ToIntSliceE convert an interface to a IntSlice type.
func ToIntSliceE(obj interface{}) (val *IntSlice, err error) {
	val = &IntSlice{}
	o := DeReference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		*val = append(*val, ToInt(x))
	case []bool:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}
	case []*bool:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		*val = append(*val, ToInt(x))
	case []Char:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}
	case []*Char:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		*val = append(*val, int(x))
	case []float32:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*float32:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		*val = append(*val, int(x))
	case []float64:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*float64:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.WithMessagef(e, "unable to convert %T to []int", x)
			}
		}

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		*val = append(*val, x)
	case []int:
		v := IntSlice(x)
		val = &v
	case []*int:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		*val = append(*val, int(x))
	case []int8:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int8:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		*val = append(*val, int(x))
	case []int16:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int16:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		*val = append(*val, int(x))
	case []int32:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int32:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		*val = append(*val, int(x))
	case []int64:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*int64:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		val = &x
	case []IntSlice:
		for i := range x {
			*val = append(*val, x[i]...)
		}
	case []*IntSlice:
		for i := range x {
			if x[i] != nil {
				*val = append(*val, *x[i]...)
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		if v, e := x.ToIntE(); e != nil {
			err = fmt.Errorf("unable to convert type Object to []int")
		} else {
			*val = append(*val, v)
		}
	case []Object:
		for i := range x {
			if v, e := x[i].ToIntE(); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type %T to []int", x[i].O())
				return
			}
		}
	case []*Object:
		for i := range x {
			if v, e := x[i].ToIntE(); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type %T to []int", x[i].O())
				return
			}
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		if v, e := ToIntE(x); e != nil {
			err = fmt.Errorf("unable to convert type Str to []int")
		} else {
			*val = append(*val, v)
		}
	case []Str:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type Str to []int")
				return
			}
		}
	case []*Str:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type *Str to []int")
				return
			}
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		if v, e := ToIntE(x); e != nil {
			err = fmt.Errorf("unable to convert type string to []int")
		} else {
			*val = append(*val, v)
		}
	case []string:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type string to []int")
				return
			}
		}
	case []*string:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type *string to []int")
				return
			}
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		*val = append(*val, int(x))
	case []uint:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		*val = append(*val, int(x))
	case []uint8:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint8:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		*val = append(*val, int(x))
	case []uint16:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint16:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		*val = append(*val, int(x))
	case []uint32:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint32:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		*val = append(*val, int(x))

	case []uint64:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []*uint64:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(x)
		k := v.Kind()

		switch {

		// generically convert array and slice types
		case k == reflect.Array || k == reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				if v, e := ToIntE(v.Index(i).Interface()); e == nil {
					*val = append(*val, v)
				} else {
					err = errors.WithMessagef(e, "unable to convert %T to []int", x)
					return
				}
			}

		// not supporting this type yet
		default:
			err = fmt.Errorf("unable to convert type %T to []int", x)
			return
		}
	}
	return
}

// ToStr convert an interface to a *Str type.
func ToStr(obj interface{}) *Str {
	val := Str{}
	o := Reference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case *bool:
		if x != nil {
			val = Str(ToString(x))
		}
	case *[]bool:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*bool:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case *Char:
		if x != nil {
			val = Str(x.String())
		}
	case *[]Char:
		if x != nil {
			for i := range *x {
				val = append(val, Str((*x)[i].String())...)
			}
		}
	case *[]*Char:
		if x != nil {
			for i := range *x {
				val = append(val, Str((*x)[i].String())...)
			}
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case *float32:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]float32:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*float32:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case *float64:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]float64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*float64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case *[]interface{}:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// int
	//----------------------------------------------------------------------------------------------
	case *int:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]int:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*int:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case *int8:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]int8:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*int8:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case *int16:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]int16:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*int16:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case *int32:
		if x != nil && *x != rune(0) {
			val = append(val, *x)
		}
	case *[]int32:
		if x != nil {
			val = append(val, *x...)
		}
	case *[]*int32:
		if x != nil {
			for i := range *x {
				y := (*x)[i]
				if y != nil {
					val = append(val, *y)
				}
			}
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case *int64:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]int64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*int64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case *IntSlice:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]IntSlice:
		if x != nil {
			for i := range *x {
				for j := range (*x)[i] {
					val = append(val, Str(ToString((*x)[i][j]))...)
				}
			}
		}
	case *[]*IntSlice:
		if x != nil {
			for i := range *x {
				y := (*x)[i]
				if y != nil {
					for j := range *y {
						val = append(val, Str(ToString((*y)[j]))...)
					}
				}
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case *Object:
		val = Str(x.String())
	case *[]Object:
		if x != nil {
			for i := range *x {
				val = append(val, Str((*x)[i].String())...)
			}
		}
	case *[]*Object:
		if x != nil {
			for i := range *x {
				val = append(val, Str((*x)[i].String())...)
			}
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case *Str:
		if x != nil {
			val = *x
		}
	case *[]Str:
		if x != nil {
			for i := range *x {
				val = append(val, (*x)[i]...)
			}
		}
	case *[]*Str:
		if x != nil {
			for i := range *x {
				y := (*x)[i]
				if y != nil {
					val = append(val, *y...)
				}
			}
		}

	// string
	//----------------------------------------------------------------------------------------------
	case *string:
		if x != nil {
			val = Str(*x)
		}
	case *[]string:
		if x != nil {
			for i := range *x {
				val = append(val, Str((*x)[i])...)
			}
		}
	case *[]*string:
		if x != nil {
			for i := range *x {
				y := (*x)[i]
				if y != nil {
					val = append(val, Str(*y)...)
				}
			}
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case *uint:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]uint:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*uint:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// uint8 i.e. byte handle differently
	//----------------------------------------------------------------------------------------------
	case *uint8:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]uint8:
		if x != nil && len(*x) != 0 {
			val = Str(string(*x))
		}
	case *[]*uint8:
		if x != nil {
			bytes := []byte{}
			for i := range *x {
				y := (*x)[i]
				if y != nil {
					bytes = append(bytes, *y)
				}
			}
			val = Str(string(bytes))
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case *uint16:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]uint16:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*uint16:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case *uint32:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]uint32:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*uint32:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case *uint64:
		if x != nil {
			val = Str(ToString(*x))
		}
	case *[]uint64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}
	case *[]*uint64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(ToString((*x)[i]))...)
			}
		}

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(x)
		k := v.Kind()

		switch {

		// generically convert array and slice types
		case k == reflect.Array || k == reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				val = append(val, *ToStr(v.Index(i).Interface())...)
			}

		default:
			val = Str(ToString(x))
		}
	}
	return &val
}

// ToString convert an interface to a string type.
func ToString(obj interface{}) (val string) {
	o := Reference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case *bool:
		if x != nil {
			val = strconv.FormatBool(*x)
		}

	// []byte
	//----------------------------------------------------------------------------------------------
	case *[]byte:
		if x != nil && len(*x) != 0 {
			val = string(*x)
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case *Char:
		val = x.String()

	// error
	//----------------------------------------------------------------------------------------------
	case error:
		val = x.Error()

	// float32
	//----------------------------------------------------------------------------------------------
	case *float32:
		if x != nil {
			val = strconv.FormatFloat(float64(*x), 'f', -1, 32)
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case *float64:
		if x != nil {
			val = strconv.FormatFloat(float64(*x), 'f', -1, 64)
		}

	// int
	//----------------------------------------------------------------------------------------------
	case *int:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case *int8:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case *int16:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case *int32:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case *int64:
		if x != nil {
			val = strconv.FormatInt(*x, 10)
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case *Object:
		if x != nil {
			val = x.ToString()
		}

	// []rune
	//----------------------------------------------------------------------------------------------
	case *[]rune:
		if x != nil && len(*x) != 0 {
			val = string(*x)
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case *Str:
		val = x.String()

	// string
	//----------------------------------------------------------------------------------------------
	case *string:
		if x != nil {
			val = *x
		}

	// template.CSS
	//----------------------------------------------------------------------------------------------
	case *template.CSS:
		if x != nil {
			val = string(*x)
		}

	// template.HTML
	//----------------------------------------------------------------------------------------------
	case *template.HTML:
		if x != nil {
			val = string(*x)
		}

	// template.HTMLAttr
	//----------------------------------------------------------------------------------------------
	case *template.HTMLAttr:
		if x != nil {
			val = string(*x)
		}

	// template.JS
	//----------------------------------------------------------------------------------------------
	case *template.JS:
		if x != nil {
			val = string(*x)
		}

	// template.URL
	//----------------------------------------------------------------------------------------------
	case *template.URL:
		if x != nil {
			val = string(*x)
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case *uint:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case *uint8:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case *uint16:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case *uint32:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case *uint64:
		if x != nil {
			val = strconv.FormatInt(int64(*x), 10)
		}

	// fmt.Stringer
	//----------------------------------------------------------------------------------------------
	case fmt.Stringer:
		if x != nil {
			val = x.String()
		}

	default:
		val = fmt.Sprintf("%v", obj)
	}
	return
}

// ToStringSlice convert an interface to a []int type.
func ToStringSlice(obj interface{}) *StringSlice {
	x, _ := ToStringSliceE(obj)
	if x == nil {
		return &StringSlice{}
	}
	return x
}

// ToStringSliceE convert an interface to a []string type.
func ToStringSliceE(obj interface{}) (val *StringSlice, err error) {
	val = &StringSlice{}
	o := DeReference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		*val = append(*val, ToString(x))
	case []bool:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*bool:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		*val = append(*val, ToString(x))
	case []Char:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*Char:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		*val = append(*val, ToString(x))
	case []float32:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*float32:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		*val = append(*val, ToString(x))
	case []float64:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*float64:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		*val = append(*val, ToString(x))
	case []int:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*int:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		*val = append(*val, ToString(x))
	case []int8:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*int8:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		*val = append(*val, ToString(x))
	case []int16:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*int16:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		*val = append(*val, ToString(x))
	case []int32:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*int32:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		*val = append(*val, ToString(x))
	case []int64:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*int64:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []IntSlice:
		for i := range x {
			for j := range x[i] {
				*val = append(*val, ToString(x[i][j]))
			}
		}
	case []*IntSlice:
		for i := range x {
			for j := range *x[i] {
				*val = append(*val, ToString((*x[i])[j]))
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		val, err = ToStringSliceE(ToString(x))
	case []Object:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*Object:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		val, err = ToStringSliceE(ToString(x))
	case []Str:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*Str:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		*val = strings.Fields(x)
	case []string:
		*val = x
	case []*string:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		*val = append(*val, ToString(x))
	case []uint:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*uint:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		*val = append(*val, ToString(x))
	case []uint8:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*uint8:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		*val = append(*val, ToString(x))
	case []uint16:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*uint16:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		*val = append(*val, ToString(x))
	case []uint32:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*uint32:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		*val = append(*val, ToString(x))

	case []uint64:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*uint64:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(x)
		k := v.Kind()

		switch {

		// generically convert array and slice types
		case k == reflect.Array || k == reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				*val = append(*val, ToString(v.Index(i).Interface()))
			}

		default:
			err = fmt.Errorf("unable to convert type %T to []int", x)
			return
		}
	}
	return
}

// ToStringSliceGE convert an interface to a []string type.
func ToStringSliceGE(obj interface{}) (val []string, err error) {
	val = []string{}
	o := DeReference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		val = append(val, ToString(x))
	case []bool:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*bool:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		val = append(val, ToString(x))
	case []Char:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*Char:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		val = append(val, ToString(x))
	case []float32:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*float32:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		val = append(val, ToString(x))
	case []float64:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*float64:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		val = append(val, ToString(x))
	case []int:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*int:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		val = append(val, ToString(x))
	case []int8:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*int8:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		val = append(val, ToString(x))
	case []int16:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*int16:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		val = append(val, ToString(x))
	case []int32:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*int32:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		val = append(val, ToString(x))
	case []int64:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*int64:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []IntSlice:
		for i := range x {
			for j := range x[i] {
				val = append(val, ToString(x[i][j]))
			}
		}
	case []*IntSlice:
		for i := range x {
			for j := range *x[i] {
				val = append(val, ToString((*x[i])[j]))
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		val, err = ToStringSliceGE(ToString(x))
	case []Object:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*Object:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		val, err = ToStringSliceGE(ToString(x))
	case []Str:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*Str:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		val = strings.Fields(x)
	case []string:
		val = x
	case []*string:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		val = append(val, ToString(x))
	case []uint:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*uint:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		val = append(val, ToString(x))
	case []uint8:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*uint8:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		val = append(val, ToString(x))
	case []uint16:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*uint16:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		val = append(val, ToString(x))
	case []uint32:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*uint32:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		val = append(val, ToString(x))

	case []uint64:
		for i := range x {
			val = append(val, ToString(x[i]))
		}
	case []*uint64:
		for i := range x {
			val = append(val, ToString(x[i]))
		}

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(x)
		k := v.Kind()

		switch {

		// generically convert array and slice types
		case k == reflect.Array || k == reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				val = append(val, ToString(v.Index(i).Interface()))
			}

		default:
			err = fmt.Errorf("unable to convert type %T to []int", x)
			return
		}
	}
	return
}
