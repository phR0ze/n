package n

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/ghodss/yaml"
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

	// FloatSlice
	//----------------------------------------------------------------------------------------------
	case FloatSlice:
		return x
	case *FloatSlice:
		if x == nil {
			return *NewFloatSliceV()
		}
		return *x
	case []FloatSlice:
		return x
	case []*FloatSlice:
		return x
	case *[]FloatSlice:
		if x == nil {
			return []FloatSlice{}
		}
		return *x
	case *[]*FloatSlice:
		if x == nil {
			return []*FloatSlice{}
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
	case [][]int32:
		return x
	case [][]*int32:
		return x
	case *[][]int32:
		if x == nil {
			return [][]int32{}
		}
		return *x
	case *[][]*int32:
		if x == nil {
			return [][]*int32{}
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
	case [][]uint8:
		return x
	case [][]*uint8:
		return x
	case *[][]uint8:
		if x == nil {
			return [][]uint8{}
		}
		return *x
	case *[][]*uint8:
		if x == nil {
			return [][]*uint8{}
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

	// FloatSlice
	//----------------------------------------------------------------------------------------------
	case FloatSlice:
		return &x
	case *FloatSlice:
		return x
	case []FloatSlice:
		return &x
	case []*FloatSlice:
		return &x
	case *[]FloatSlice:
		return x
	case *[]*FloatSlice:
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
	case [][]int32:
		return &x
	case [][]*int32:
		return &x
	case *[][]int32:
		return x
	case *[][]*int32:
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

	// StringMap
	//----------------------------------------------------------------------------------------------
	case StringMap:
		return &x
	case *StringMap:
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
	case []map[interface{}]interface{}:
		return &x
	case *[]map[interface{}]interface{}:
		return x
	case map[interface{}]interface{}:
		return &x
	case []map[interface{}]*interface{}:
		return &x
	case *[]map[interface{}]*interface{}:
		return x
	case map[interface{}]*interface{}:
		return &x
	case []*map[interface{}]interface{}:
		return &x
	case *[]*map[interface{}]interface{}:
		return x
	case *map[interface{}]interface{}:
		return x
	case []*map[interface{}]*interface{}:
		return &x
	case *[]*map[interface{}]*interface{}:
		return x
	case *map[interface{}]*interface{}:
		return x
	case []map[string]interface{}:
		return &x
	case *[]map[string]interface{}:
		return x
	case map[string]interface{}:
		return &x
	case []map[string]*interface{}:
		return &x
	case *[]map[string]*interface{}:
		return x
	case map[string]*interface{}:
		return &x
	case []*map[string]interface{}:
		return &x
	case *[]*map[string]interface{}:
		return x
	case *map[string]interface{}:
		return x
	case []*map[string]*interface{}:
		return &x
	case *[]*map[string]*interface{}:
		return x
	case *map[string]*interface{}:
		return x
	case []map[string]string:
		return &x
	case *[]map[string]string:
		return x
	case map[string]string:
		return &x
	case map[string]*string:
		return &x
	case *map[string]string:
		return x
	case *map[string]*string:
		return x
	case map[string]float32:
		return &x
	case map[string]*float32:
		return &x
	case *map[string]float32:
		return x
	case *map[string]*float32:
		return x
	case map[string]float64:
		return &x
	case map[string]*float64:
		return &x
	case *map[string]float64:
		return x
	case *map[string]*float64:
		return x
	case map[string]int:
		return &x
	case map[string]*int:
		return &x
	case *map[string]int:
		return x
	case *map[string]*int:
		return x
	case map[string]int64:
		return &x
	case map[string]*int64:
		return &x
	case *map[string]int64:
		return x
	case *map[string]*int64:
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
	case [][]uint8:
		return &x
	case [][]*uint8:
		return &x
	case *[][]uint8:
		return x
	case *[][]*uint8:
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

// C is an alias to ToChar for brevity
func C(obj interface{}) *Char {
	return ToChar(obj)
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

// ToFloat32 convert an interface to a float32 type.
func ToFloat32(obj interface{}) float32 {
	x, _ := ToFloat32E(obj)
	return x
}

// ToFloat32E convert an interface to a float32 type.
func ToFloat32E(obj interface{}) (val float32, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToFloat32E(x.A())
	case float32:
		val = x
	case float64:
		val = float32(x)
	case int:
		val = float32(x)
	case int8:
		val = float32(x)
	case int16:
		val = float32(x)
	case int32:
		val = float32(x)
	case int64:
		val = float32(x)
	case Object:
		val, err = x.ToFloat32E()
	case uint:
		val = float32(x)
	case uint8:
		val = float32(x)
	case uint16:
		val = float32(x)
	case uint32:
		val = float32(x)
	case uint64:
		val = float32(x)
	case Str:
		return ToFloat32E(string(x))
	case string:
		var v float64
		if v, err = strconv.ParseFloat(x, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to float32")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = float32(v)
		}
	default:
		err = fmt.Errorf("unable to convert type %T to int", x)
	}
	return
}

// ToFloat64 convert an interface to a float64 type.
func ToFloat64(obj interface{}) float64 {
	x, _ := ToFloat64E(obj)
	return x
}

// ToFloat64E convert an interface to a float64 type.
func ToFloat64E(obj interface{}) (val float64, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToFloat64E(x.A())
	case float32:
		val = float64(x)
	case float64:
		val = x
	case int:
		val = float64(x)
	case int8:
		val = float64(x)
	case int16:
		val = float64(x)
	case int32:
		val = float64(x)
	case int64:
		val = float64(x)
	case Object:
		val, err = x.ToFloat64E()
	case uint:
		val = float64(x)
	case uint8:
		val = float64(x)
	case uint16:
		val = float64(x)
	case uint32:
		val = float64(x)
	case uint64:
		val = float64(x)
	case Str:
		return ToFloat64E(string(x))
	case string:
		if val, err = strconv.ParseFloat(x, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to float64")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		}
	default:
		err = fmt.Errorf("unable to convert type %T to int", x)
	}
	return
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

// ToFloatSlice convert an interface to a FloatSlice type which will never be nil
func ToFloatSlice(obj interface{}) *FloatSlice {
	x, _ := ToFloatSliceE(obj)
	if x == nil {
		return &FloatSlice{}
	}
	return x
}

// ToFloatSliceE convert an interface to a FloatSlice type.
func ToFloatSliceE(obj interface{}) (val *FloatSlice, err error) {
	val = &FloatSlice{}
	o := DeReference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case bool:
		*val = append(*val, ToFloat64(x))
	case []bool:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*bool:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case Char:
		*val = append(*val, ToFloat64(x))
	case []Char:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*Char:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case float32:
		*val = append(*val, ToFloat64(x))
	case []float32:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*float32:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case float64:
		*val = append(*val, x)
	case []float64:
		*val = x
	case []*float64:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// FloatSlice
	//----------------------------------------------------------------------------------------------
	case FloatSlice:
		val = &x
	case []FloatSlice:
		for i := range x {
			*val = append(*val, x[i]...)
		}
	case []*FloatSlice:
		for i := range x {
			if x[i] != nil {
				*val = append(*val, *x[i]...)
			}
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.WithMessagef(e, "unable to convert %T to []float64", x)
			}
		}

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		*val = append(*val, ToFloat64(x))
	case []int:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*int:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case int8:
		*val = append(*val, ToFloat64(x))
	case []int8:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*int8:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case int16:
		*val = append(*val, ToFloat64(x))
	case []int16:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*int16:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case int32:
		*val = append(*val, ToFloat64(x))
	case []int32:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*int32:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case int64:
		*val = append(*val, ToFloat64(x))
	case []int64:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*int64:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []IntSlice:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*IntSlice:
		for i := range x {
			if x[i] != nil {
				for j := range *x[i] {
					*val = append(*val, ToFloat64((*x[i])[j]))
				}
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		if v, e := x.ToFloat64E(); e != nil {
			err = fmt.Errorf("unable to convert type Object to []float64")
		} else {
			*val = append(*val, v)
		}
	case []Object:
		for i := range x {
			if v, e := x[i].ToFloat64E(); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type %T to []float64", x[i].O())
				return
			}
		}
	case []*Object:
		for i := range x {
			if v, e := x[i].ToFloat64E(); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type %T to []float64", x[i].O())
				return
			}
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		if v, e := ToFloat64E(x); e != nil {
			err = fmt.Errorf("unable to convert type Str to []float64")
		} else {
			*val = append(*val, v)
		}
	case []Str:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type Str to []float64")
				return
			}
		}
	case []*Str:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type *Str to []float64")
				return
			}
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		if v, e := ToFloat64E(x); e != nil {
			err = fmt.Errorf("unable to convert type string to []float64")
		} else {
			*val = append(*val, v)
		}
	case []string:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type string to []float64")
				return
			}
		}
	case []*string:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = fmt.Errorf("unable to convert type *string to []float64")
				return
			}
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		*val = append(*val, ToFloat64(x))
	case []uint:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*uint:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8:
		*val = append(*val, ToFloat64(x))
	case []uint8:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*uint8:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16:
		*val = append(*val, ToFloat64(x))
	case []uint16:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*uint16:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32:
		*val = append(*val, ToFloat64(x))
	case []uint32:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*uint32:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64:
		*val = append(*val, ToFloat64(x))

	case []uint64:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*uint64:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
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
				if v, e := ToFloat64E(v.Index(i).Interface()); e == nil {
					*val = append(*val, v)
				} else {
					err = errors.WithMessagef(e, "unable to convert %T to []float64", x)
					return
				}
			}

		// not supporting this type yet
		default:
			err = fmt.Errorf("unable to convert type %T to []float64", x)
			return
		}
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

// ToMapSlice converts an interface to a MapSlice type.
func ToMapSlice(obj interface{}) *MapSlice {
	x, _ := ToMapSliceE(obj)
	if x == nil {
		return &MapSlice{}
	}
	return x
}

// ToMapSliceE converts an interface to a MapSlice type.
func ToMapSliceE(obj interface{}) (val *MapSlice, err error) {
	val = &MapSlice{}
	o := Reference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// []interface{}
	//----------------------------------------------------------------------------------------------
	case *[]interface{}:
		if x != nil {
			for _, raw := range *x {
				var m *MapSlice
				if m, err = ToMapSliceE(raw); err != nil {
					return
				}
				for _, v := range *m {
					*val = append(*val, v)
				}
			}
		}

	// MapSlice
	//----------------------------------------------------------------------------------------------
	case *MapSlice:
		if x != nil {
			val = x
		}
	case *map[string]interface{}:
		if x != nil {
			*val = append(*val, *x)
		}
	case *map[string]string:
		if x != nil {
			new := map[string]interface{}{}
			for k, v := range *x {
				new[k] = v
			}
			*val = append(*val, new)
		}

	// []map[string]interface{}
	//----------------------------------------------------------------------------------------------
	case *[]map[string]interface{}:
		if x != nil {
			v := MapSlice(*x)
			val = &v
		}
	case *[]map[string]string:
		if x != nil {
			for _, raw := range *x {
				var m *MapSlice
				if m, err = ToMapSliceE(raw); err != nil {
					return
				}
				for _, v := range *m {
					*val = append(*val, v)
				}
			}
		}

	default:
		err = errors.Errorf("failed to convert type %T to a MapString", x)
	}

	return
}

// R is an alias to ToRune for brevity
func R(obj interface{}) rune {
	return rune(*ToChar(obj))
}

// ToRune convert an interface to a rune type.
func ToRune(obj interface{}) rune {
	return rune(*ToChar(obj))
}

// A is an alias to ToStr for brevity
func A(obj interface{}) *Str {
	return ToStr(obj)
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
			val = Str(strconv.FormatBool(*x))
		}
	case *[]bool:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatBool((*x)[i]))...)
			}
		}
	case *[]*bool:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatBool(*y))...)
				}
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
				if y := (*x)[i]; y.G() != rune(0) {
					val = append(val, Str(y.String())...)
				}
			}
		}
	case *[]*Char:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil && y.G() != rune(0) {
					val = append(val, Str(y.String())...)
				}
			}
		}

	// error
	//----------------------------------------------------------------------------------------------
	case error:
		val = Str(x.Error())

	// float32
	//----------------------------------------------------------------------------------------------
	case *float32:
		if x != nil {
			val = Str(strconv.FormatFloat(float64(*x), 'f', -1, 32))
		}
	case *[]float32:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatFloat(float64((*x)[i]), 'f', -1, 32))...)
			}
		}
	case *[]*float32:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatFloat(float64(*y), 'f', -1, 32))...)
				}
			}
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case *float64:
		if x != nil {
			val = Str(strconv.FormatFloat(float64(*x), 'f', -1, 64))
		}
	case *[]float64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatFloat(float64((*x)[i]), 'f', -1, 64))...)
			}
		}
	case *[]*float64:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatFloat(float64(*y), 'f', -1, 64))...)
				}
			}
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case *[]interface{}:
		if x != nil {
			for i := range *x {
				val = append(val, *ToStr((*x)[i])...)
			}
		}

	// int
	//----------------------------------------------------------------------------------------------
	case *int:
		if x != nil {
			val = Str(strconv.FormatInt(int64(*x), 10))
		}
	case *[]int:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatInt(int64((*x)[i]), 10))...)
			}
		}
	case *[]*int:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatInt(int64(*y), 10))...)
				}
			}
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case *int8:
		if x != nil {
			val = Str(strconv.FormatInt(int64(*x), 10))
		}
	case *[]int8:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatInt(int64((*x)[i]), 10))...)
			}
		}
	case *[]*int8:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatInt(int64(*y), 10))...)
				}
			}
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case *int16:
		if x != nil {
			val = Str(strconv.FormatInt(int64(*x), 10))
		}
	case *[]int16:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatInt(int64((*x)[i]), 10))...)
			}
		}
	case *[]*int16:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatInt(int64(*y), 10))...)
				}
			}
		}

	// int32 a.k.a rune handle differently as Golang only has a single type
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
				if y := (*x)[i]; y != nil {
					val = append(val, *y)
				}
			}
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case *int64:
		if x != nil {
			val = Str(strconv.FormatInt(*x, 10))
		}
	case *[]int64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatInt((*x)[i], 10))...)
			}
		}
	case *[]*int64:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatInt(*y, 10))...)
				}
			}
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case *IntSlice:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatInt(int64((*x)[i]), 10))...)
			}
		}
	case *[]IntSlice:
		if x != nil {
			for i := range *x {
				for j := range (*x)[i] {
					val = append(val, Str(strconv.FormatInt(int64((*x)[i][j]), 10))...)
				}
			}
		}
	case *[]*IntSlice:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					for j := range *y {
						val = append(val, Str(strconv.FormatInt(int64((*y)[j]), 10))...)
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

	// template.CSS
	//----------------------------------------------------------------------------------------------
	case *template.CSS:
		if x != nil {
			val = Str(string(*x))
		}

	// template.HTML
	//----------------------------------------------------------------------------------------------
	case *template.HTML:
		if x != nil {
			val = Str(string(*x))
		}

	// template.HTMLAttr
	//----------------------------------------------------------------------------------------------
	case *template.HTMLAttr:
		if x != nil {
			val = Str(string(*x))
		}

	// template.JS
	//----------------------------------------------------------------------------------------------
	case *template.JS:
		if x != nil {
			val = Str(string(*x))
		}

	// template.URL
	//----------------------------------------------------------------------------------------------
	case *template.URL:
		if x != nil {
			val = Str(string(*x))
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case *uint:
		if x != nil {
			val = Str(strconv.FormatUint(uint64(*x), 10))
		}
	case *[]uint:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatUint(uint64((*x)[i]), 10))...)
			}
		}
	case *[]*uint:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatUint(uint64(*y), 10))...)
				}
			}
		}

	// uint8 a.k.a byte handle differently as Golang only has a single type
	//----------------------------------------------------------------------------------------------
	case *uint8:
		if x != nil && *x != byte(0) {
			val = append(val, rune(*x))
		}
	case *[]uint8:
		if x != nil && len(*x) != 0 {
			val = Str(string(*x))
		}
	case *[]*uint8:
		if x != nil {
			for i := range *x {
				y := (*x)[i]
				if y != nil {
					val = append(val, rune(*y))
				}
			}
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case *uint16:
		if x != nil {
			val = Str(strconv.FormatUint(uint64(*x), 10))
		}
	case *[]uint16:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatUint(uint64((*x)[i]), 10))...)
			}
		}
	case *[]*uint16:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatUint(uint64(*y), 10))...)
				}
			}
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case *uint32:
		if x != nil {
			val = Str(strconv.FormatUint(uint64(*x), 10))
		}
	case *[]uint32:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatUint(uint64((*x)[i]), 10))...)
			}
		}
	case *[]*uint32:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatUint(uint64(*y), 10))...)
				}
			}
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case *uint64:
		if x != nil {
			val = Str(strconv.FormatUint(*x, 10))
		}
	case *[]uint64:
		if x != nil {
			for i := range *x {
				val = append(val, Str(strconv.FormatUint((*x)[i], 10))...)
			}
		}
	case *[]*uint64:
		if x != nil {
			for i := range *x {
				if y := (*x)[i]; y != nil {
					val = append(val, Str(strconv.FormatUint(*y, 10))...)
				}
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
			if y, ok := x.(fmt.Stringer); ok && y != nil {
				val = Str(y.String())
			} else {
				val = Str(fmt.Sprintf("%v", x))
			}
		}
	}
	return &val
}

// ToString convert an interface to a string type.
func ToString(obj interface{}) string {
	return ToStr(obj).A()
}

// ToYaml is an alias to ToStringMap
func ToYaml(obj interface{}) *StringMap {
	x, _ := ToStringMapE(obj)
	if x == nil {
		return &StringMap{}
	}
	return x
}

// ToStringMap converts an interface to a StringMap type. Supports converting yaml string as well.
func ToStringMap(obj interface{}) *StringMap {
	x, _ := ToStringMapE(obj)
	if x == nil {
		return &StringMap{}
	}
	return x
}

// ToStringMapE converts an interface to a StringMap type. Supports converting yaml string as well.
func ToStringMapE(obj interface{}) (val *StringMap, err error) {
	val = &StringMap{}
	o := Reference(obj)

	// Optimized types
	switch x := o.(type) {
	case nil:

	// byte
	//----------------------------------------------------------------------------------------------
	case *[]byte:
		if x != nil && len(*x) != 0 {
			m := map[string]interface{}{}
			if err = yaml.Unmarshal(*x, &m); err != nil {
				err = errors.Wrap(err, "failed to unmarshal bytes into StringMap")
				return
			}
			val = ToStringMap(m)
		}

	// maps
	//----------------------------------------------------------------------------------------------
	case *map[interface{}]interface{}:
		if x != nil {
			for k, v := range *x {
				val.Set(ToString(k), v)
			}
		}
	case *map[string]interface{}:
		if x != nil {
			y := StringMap(*x)
			val = &y
		}
	case *map[string]string:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case *Object:
		if x != nil {
			val, err = ToStringMapE(x.o)
		}

	// string
	//----------------------------------------------------------------------------------------------
	case *string:
		if x != nil {
			m := map[string]interface{}{}
			if err = yaml.Unmarshal([]byte(*x), &m); err != nil {
				err = errors.Wrap(err, "failed to unmarshal string into StringMap")
				return
			}
			val = ToStringMap(m)
		}

	// StringMap
	//----------------------------------------------------------------------------------------------
	case *StringMap:
		if x != nil {
			val = x
		}

	default:
		err = errors.Errorf("failed to convert type %T to a StringMap", obj)
	}

	return
}

// ToStrs is an alias to ToStringSliceE
func ToStrs(obj interface{}) *StringSlice {
	x, _ := ToStringSliceE(obj)
	if x == nil {
		return &StringSlice{}
	}
	return x
}

// ToStringSlice convert an interface to a StringSlice type.
func ToStringSlice(obj interface{}) *StringSlice {
	x, _ := ToStringSliceE(obj)
	if x == nil {
		return &StringSlice{}
	}
	return x
}

// ToStringSliceE convert an interface to a StringSlice type.
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

	// StringSlice
	//----------------------------------------------------------------------------------------------
	case StringSlice:
		val = &x
	case []StringSlice:
		for i := range x {
			*val = append(*val, x[i]...)
		}
	case []*StringSlice:
		for i := range x {
			if x[i] != nil {
				*val = append(*val, *x[i]...)
			}
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

// ToStringSliceG convert an interface to a []string type.
func ToStringSliceG(obj interface{}) (val []string) {
	val, _ = ToStringSliceGE(obj)
	return val
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

	// StringSlice
	//----------------------------------------------------------------------------------------------
	case StringSlice:
		val = x.G()
	case []StringSlice:
		for i := range x {
			val = append(val, x[i].G()...)
		}
	case []*StringSlice:
		for i := range x {
			if x[i] != nil {
				val = append(val, x[i].G()...)
			}
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
