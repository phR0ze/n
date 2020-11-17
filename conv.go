package n

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

var gUseLocalTime bool

// TimeLayouts is just a simple wrapper around popular time layouts for time.Parse
var TimeLayouts = []string{
	time.RFC3339,  // "2006-01-02T15:04:05Z07:00" // ISO8601
	time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST"
	time.RFC822Z,  // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	time.RFC822,   // "02 Jan 06 15:04 MST"
	time.RFC850,   // "Monday, 02-Jan-06 15:04:05 MST"
	time.ANSIC,    // "Mon Jan _2 15:04:05 2006"
	time.UnixDate, // "Mon Jan _2 15:04:05 MST 2006"
	time.RubyDate, // "Mon Jan 02 15:04:05 -0700 2006"

	// Human formats based on Golang's magic value "Mon Jan 2 15:04:05 MST 2006" or 1136239445
	"January 2, 2006", // US: Month Day, Year
	"02 Jan 2006",     // Day Month Year
	"2006-01-02",      // ISO: Year-Month-Day
	time.Kitchen,      // "3:04PM"

	// Time stamps
	time.StampNano,  // "Jan _2 15:04:05.000000000"
	time.StampMicro, // "Jan _2 15:04:05.000000"
	time.StampMilli, // "Jan _2 15:04:05.000"
	time.Stamp,      // "Jan _2 15:04:05"
}

// UseLocalTime controls whether the ToTime functions will use UTC or Local for Unix functions
func UseLocalTime(useLocal bool) {
	gUseLocalTime = useLocal
}

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

	// InterSlice
	//----------------------------------------------------------------------------------------------
	case InterSlice:
		return x
	case *InterSlice:
		if x == nil {
			return *NewInterSliceV()
		}
		return *x
	case []InterSlice:
		return x
	case []*InterSlice:
		return x
	case *[]InterSlice:
		if x == nil {
			return []InterSlice{}
		}
		return *x
	case *[]*InterSlice:
		if x == nil {
			return []*InterSlice{}
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

	// StringMap
	//---------------------------------------------------------------------------------
	case StringMap:
		return x
	case *StringMap:
		if x == nil {
			return *NewStringMapV()
		}
		return *x
	case []StringMap:
		return x
	case []*StringMap:
		return x
	case *[]StringMap:
		if x == nil {
			return []StringMap{}
		}
		return *x
	case *[]*StringMap:
		if x == nil {
			return []*StringMap{}
		}
		return *x

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

	// map[interface{}]interface{}
	//----------------------------------------------------------------------------------------------
	case []map[interface{}]interface{}:
		return x
	case *[]map[interface{}]interface{}:
		if x == nil {
			return []map[interface{}]interface{}{}
		}
		return *x
	case map[interface{}]interface{}:
		return x
	case []map[interface{}]*interface{}:
		return x
	case *[]map[interface{}]*interface{}:
		if x == nil {
			return []map[interface{}]*interface{}{}
		}
		return *x
	case map[interface{}]*interface{}:
		return x
	case []*map[interface{}]interface{}:
		return x
	case *[]*map[interface{}]interface{}:
		if x == nil {
			return []*map[interface{}]interface{}{}
		}
		return *x
	case *map[interface{}]interface{}:
		if x == nil {
			return map[interface{}]interface{}{}
		}
		return *x
	case []*map[interface{}]*interface{}:
		return x
	case *[]*map[interface{}]*interface{}:
		if x == nil {
			return []*map[interface{}]*interface{}{}
		}
		return *x
	case *map[interface{}]*interface{}:
		if x == nil {
			return map[interface{}]*interface{}{}
		}
		return *x

	// map[string]interface{}
	//----------------------------------------------------------------------------------------------
	case []map[string]interface{}:
		return x
	case *[]map[string]interface{}:
		if x == nil {
			return []map[string]interface{}{}
		}
		return *x
	case map[string]interface{}:
		return x
	case []map[string]*interface{}:
		return x
	case *[]map[string]*interface{}:
		if x == nil {
			return []map[string]*interface{}{}
		}
		return *x
	case map[string]*interface{}:
		return x
	case []*map[string]interface{}:
		return x
	case *[]*map[string]interface{}:
		if x == nil {
			return []*map[string]interface{}{}
		}
		return *x
	case *map[string]interface{}:
		if x == nil {
			return map[string]interface{}{}
		}
		return *x
	case []*map[string]*interface{}:
		return x
	case *[]*map[string]*interface{}:
		if x == nil {
			return []*map[string]*interface{}{}
		}
		return *x
	case *map[string]*interface{}:
		if x == nil {
			return map[string]*interface{}{}
		}
		return *x

	// map[string]string
	//----------------------------------------------------------------------------------------------
	case []map[string]string:
		return x
	case *[]map[string]string:
		if x == nil {
			return []map[string]string{}
		}
		return *x
	case map[string]string:
		return x
	case map[string]*string:
		return x
	case *map[string]string:
		if x == nil {
			return map[string]string{}
		}
		return *x
	case *map[string]*string:
		if x == nil {
			return map[string]*string{}
		}
		return *x

	// map[string]bool
	//----------------------------------------------------------------------------------------------
	case map[string]bool:
		return x
	case map[string]*bool:
		return x
	case *map[string]bool:
		if x == nil {
			return map[string]bool{}
		}
		return *x
	case *map[string]*bool:
		if x == nil {
			return map[string]*bool{}
		}
		return *x

	// map[string]float32
	//----------------------------------------------------------------------------------------------
	case map[string]float32:
		return x
	case map[string]*float32:
		return x
	case *map[string]float32:
		if x == nil {
			return map[string]float32{}
		}
		return *x
	case *map[string]*float32:
		if x == nil {
			return map[string]*float32{}
		}
		return *x

	// map[string]float64
	//----------------------------------------------------------------------------------------------
	case map[string]float64:
		return x
	case map[string]*float64:
		return x
	case *map[string]float64:
		if x == nil {
			return map[string]float64{}
		}
		return *x
	case *map[string]*float64:
		if x == nil {
			return map[string]*float64{}
		}
		return *x

	// map[string]int
	//----------------------------------------------------------------------------------------------
	case map[string]int:
		return x
	case map[string]*int:
		return x
	case *map[string]int:
		if x == nil {
			return map[string]int{}
		}
		return *x
	case *map[string]*int:
		if x == nil {
			return map[string]*int{}
		}
		return *x

	// map[string]int8
	//----------------------------------------------------------------------------------------------
	case map[string]int8:
		return x
	case map[string]*int8:
		return x
	case *map[string]int8:
		if x == nil {
			return map[string]int8{}
		}
		return *x
	case *map[string]*int8:
		if x == nil {
			return map[string]*int8{}
		}
		return *x

	// map[string]int16
	//----------------------------------------------------------------------------------------------
	case map[string]int16:
		return x
	case map[string]*int16:
		return x
	case *map[string]int16:
		if x == nil {
			return map[string]int16{}
		}
		return *x
	case *map[string]*int16:
		if x == nil {
			return map[string]*int16{}
		}
		return *x

	// map[string]int32
	//----------------------------------------------------------------------------------------------
	case map[string]int32:
		return x
	case map[string]*int32:
		return x
	case *map[string]int32:
		if x == nil {
			return map[string]int32{}
		}
		return *x
	case *map[string]*int32:
		if x == nil {
			return map[string]*int32{}
		}
		return *x

	// map[string]int64
	//----------------------------------------------------------------------------------------------
	case map[string]int64:
		return x
	case map[string]*int64:
		return x
	case *map[string]int64:
		if x == nil {
			return map[string]int64{}
		}
		return *x
	case *map[string]*int64:
		if x == nil {
			return map[string]*int64{}
		}
		return *x

	// map[string]uint
	//----------------------------------------------------------------------------------------------
	case map[string]uint:
		return x
	case map[string]*uint:
		return x
	case *map[string]uint:
		if x == nil {
			return map[string]uint{}
		}
		return *x
	case *map[string]*uint:
		if x == nil {
			return map[string]*uint{}
		}
		return *x

	// map[string]uint8
	//----------------------------------------------------------------------------------------------
	case map[string]uint8:
		return x
	case map[string]*uint8:
		return x
	case *map[string]uint8:
		if x == nil {
			return map[string]uint8{}
		}
		return *x
	case *map[string]*uint8:
		if x == nil {
			return map[string]*uint8{}
		}
		return *x

	// map[string]uint8
	//----------------------------------------------------------------------------------------------
	case map[string]uint16:
		return x
	case map[string]*uint16:
		return x
	case *map[string]uint16:
		if x == nil {
			return map[string]uint16{}
		}
		return *x
	case *map[string]*uint16:
		if x == nil {
			return map[string]*uint16{}
		}
		return *x

	// map[string]uint32
	//----------------------------------------------------------------------------------------------
	case map[string]uint32:
		return x
	case map[string]*uint32:
		return x
	case *map[string]uint32:
		if x == nil {
			return map[string]uint32{}
		}
		return *x
	case *map[string]*uint32:
		if x == nil {
			return map[string]*uint32{}
		}
		return *x

	// map[string]uint64
	//----------------------------------------------------------------------------------------------
	case map[string]uint64:
		return x
	case map[string]*uint64:
		return x
	case *map[string]uint64:
		if x == nil {
			return map[string]uint64{}
		}
		return *x
	case *map[string]*uint64:
		if x == nil {
			return map[string]*uint64{}
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

	// InterSlice
	//----------------------------------------------------------------------------------------------
	case InterSlice:
		return &x
	case *InterSlice:
		return x
	case []InterSlice:
		return &x
	case []*InterSlice:
		return &x
	case *[]InterSlice:
		return x
	case *[]*InterSlice:
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

	// StringMap
	//----------------------------------------------------------------------------------------------
	case StringMap:
		return &x
	case *StringMap:
		return x

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

	// map[interface{}]interface{}
	//----------------------------------------------------------------------------------------------
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

	// map[string]interface{}
	//----------------------------------------------------------------------------------------------
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

	// map[string]string
	//----------------------------------------------------------------------------------------------
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

	// map[string]bool
	//----------------------------------------------------------------------------------------------
	case map[string]bool:
		return &x
	case map[string]*bool:
		return &x
	case *map[string]bool:
		return x
	case *map[string]*bool:
		return x

	// map[string]float32
	//----------------------------------------------------------------------------------------------
	case map[string]float32:
		return &x
	case map[string]*float32:
		return &x
	case *map[string]float32:
		return x
	case *map[string]*float32:
		return x

	// map[string]float64
	//----------------------------------------------------------------------------------------------
	case map[string]float64:
		return &x
	case map[string]*float64:
		return &x
	case *map[string]float64:
		return x
	case *map[string]*float64:
		return x

	// map[string]int
	//----------------------------------------------------------------------------------------------
	case map[string]int:
		return &x
	case map[string]*int:
		return &x
	case *map[string]int:
		return x
	case *map[string]*int:
		return x

	// map[string]int8
	//----------------------------------------------------------------------------------------------
	case map[string]int8:
		return &x
	case map[string]*int8:
		return &x
	case *map[string]int8:
		return x
	case *map[string]*int8:
		return x

	// map[string]int16
	//----------------------------------------------------------------------------------------------
	case map[string]int16:
		return &x
	case map[string]*int16:
		return &x
	case *map[string]int16:
		return x
	case *map[string]*int16:
		return x

	// map[string]int32
	//----------------------------------------------------------------------------------------------
	case map[string]int32:
		return &x
	case map[string]*int32:
		return &x
	case *map[string]int32:
		return x
	case *map[string]*int32:
		return x

	// map[string]int64
	//----------------------------------------------------------------------------------------------
	case map[string]int64:
		return &x
	case map[string]*int64:
		return &x
	case *map[string]int64:
		return x
	case *map[string]*int64:
		return x

	// map[string]uint
	//----------------------------------------------------------------------------------------------
	case map[string]uint:
		return &x
	case map[string]*uint:
		return &x
	case *map[string]uint:
		return x
	case *map[string]*uint:
		return x

	// map[string]uint8
	//----------------------------------------------------------------------------------------------
	case map[string]uint8:
		return &x
	case map[string]*uint8:
		return &x
	case *map[string]uint8:
		return x
	case *map[string]*uint8:
		return x

	// map[string]uint8
	//----------------------------------------------------------------------------------------------
	case map[string]uint16:
		return &x
	case map[string]*uint16:
		return &x
	case *map[string]uint16:
		return x
	case *map[string]*uint16:
		return x

	// map[string]uint32
	//----------------------------------------------------------------------------------------------
	case map[string]uint32:
		return &x
	case map[string]*uint32:
		return &x
	case *map[string]uint32:
		return x
	case *map[string]*uint32:
		return x

	// map[string]uint64
	//----------------------------------------------------------------------------------------------
	case map[string]uint64:
		return &x
	case map[string]*uint64:
		return &x
	case *map[string]uint64:
		return x
	case *map[string]*uint64:
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

// B converts an interface to a bool type.
func B(obj interface{}) bool {
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

// ToDuration converts an interface to a time.Duration type.
func ToDuration(obj interface{}) (val time.Duration) {
	x, _ := ToDurationE(obj)
	return x
}

// ToDurationE converts an interface to a time.Duration type.
func ToDurationE(obj interface{}) (val time.Duration, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case time.Duration:
		val = x
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		val = time.Duration(ToInt64(x))
	case float32, float64:
		val = time.Duration(ToFloat64(x))
	default:
		err = errors.Errorf("failed to convert type %T to time.Duration", obj)
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
		err = errors.Errorf("unable to convert type %T to int", x)
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
		err = errors.Errorf("unable to convert type %T to int", x)
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

	// InterSlice
	//----------------------------------------------------------------------------------------------
	case InterSlice:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []InterSlice:
		for i := range x {
			*val = append(*val, ToFloat64(x[i]))
		}
	case []*InterSlice:
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
			err = errors.Errorf("unable to convert type Object to []float64")
		} else {
			*val = append(*val, v)
		}
	case []Object:
		for i := range x {
			if v, e := x[i].ToFloat64E(); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type %T to []float64", x[i].O())
				return
			}
		}
	case []*Object:
		for i := range x {
			if v, e := x[i].ToFloat64E(); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type %T to []float64", x[i].O())
				return
			}
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		if v, e := ToFloat64E(x); e != nil {
			err = errors.Errorf("unable to convert type Str to []float64")
		} else {
			*val = append(*val, v)
		}
	case []Str:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type Str to []float64")
				return
			}
		}
	case []*Str:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type *Str to []float64")
				return
			}
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		if v, e := ToFloat64E(x); e != nil {
			err = errors.Errorf("unable to convert type string to []float64")
		} else {
			*val = append(*val, v)
		}
	case []string:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type string to []float64")
				return
			}
		}
	case []*string:
		for i := range x {
			if v, e := ToFloat64E(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type *string to []float64")
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
			err = errors.Errorf("unable to convert type %T to []float64", x)
			return
		}
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
		if v, err = strconv.ParseInt(x, 10, 64); err != nil {
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
		err = errors.Errorf("unable to convert type %T to int", x)
	}
	return
}

// ToInt8 convert an interface to an int8 type.
func ToInt8(obj interface{}) int8 {
	x, _ := ToInt8E(obj)
	return x
}

// ToInt8E convert an interface to an int8 type.
func ToInt8E(obj interface{}) (val int8, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToInt8E(x.A())
	case float32:
		val = int8(x)
	case float64:
		val = int8(x)
	case int:
		val = int8(x)
	case int8:
		val = x
	case int16:
		val = int8(x)
	case int32:
		val = int8(x)
	case int64:
		val = int8(x)
	case Object:
		val, err = x.ToInt8E()
	case uint:
		val = int8(x)
	case uint8:
		val = int8(x)
	case uint16:
		val = int8(x)
	case uint32:
		val = int8(x)
	case uint64:
		val = int8(x)
	case Str:
		return ToInt8E(string(x))
	case string:
		var v int64
		if v, err = strconv.ParseInt(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to int8")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = int8(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to int8", x)
	}
	return
}

// ToInt16 convert an interface to an int16 type.
func ToInt16(obj interface{}) int16 {
	x, _ := ToInt16E(obj)
	return x
}

// ToInt16E convert an interface to an int16 type.
func ToInt16E(obj interface{}) (val int16, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToInt16E(x.A())
	case float32:
		val = int16(x)
	case float64:
		val = int16(x)
	case int:
		val = int16(x)
	case int8:
		val = int16(x)
	case int16:
		val = x
	case int32:
		val = int16(x)
	case int64:
		val = int16(x)
	case Object:
		val, err = x.ToInt16E()
	case uint:
		val = int16(x)
	case uint8:
		val = int16(x)
	case uint16:
		val = int16(x)
	case uint32:
		val = int16(x)
	case uint64:
		val = int16(x)
	case Str:
		return ToInt16E(string(x))
	case string:
		var v int64
		if v, err = strconv.ParseInt(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to int16")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = int16(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to int16", x)
	}
	return
}

// ToInt32 convert an interface to an int32 type.
func ToInt32(obj interface{}) int32 {
	x, _ := ToInt32E(obj)
	return x
}

// ToInt32E convert an interface to an int32 type.
func ToInt32E(obj interface{}) (val int32, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToInt32E(x.A())
	case float32:
		val = int32(x)
	case float64:
		val = int32(x)
	case int:
		val = int32(x)
	case int8:
		val = int32(x)
	case int16:
		val = int32(x)
	case int32:
		val = x
	case int64:
		val = int32(x)
	case Object:
		val, err = x.ToInt32E()
	case uint:
		val = int32(x)
	case uint8:
		val = int32(x)
	case uint16:
		val = int32(x)
	case uint32:
		val = int32(x)
	case uint64:
		val = int32(x)
	case Str:
		return ToInt32E(string(x))
	case string:
		var v int64
		if v, err = strconv.ParseInt(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to int32")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = int32(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to int32", x)
	}
	return
}

// ToInt64 convert an interface to an int64 type.
func ToInt64(obj interface{}) int64 {
	x, _ := ToInt64E(obj)
	return x
}

// ToInt64E convert an interface to an int64 type.
func ToInt64E(obj interface{}) (val int64, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToInt64E(x.A())
	case float32:
		val = int64(x)
	case float64:
		val = int64(x)
	case int:
		val = int64(x)
	case int8:
		val = int64(x)
	case int16:
		val = int64(x)
	case int32:
		val = int64(x)
	case int64:
		val = x
	case Object:
		val, err = x.ToInt64E()
	case uint:
		val = int64(x)
	case uint8:
		val = int64(x)
	case uint16:
		val = int64(x)
	case uint32:
		val = int64(x)
	case uint64:
		val = int64(x)
	case Str:
		return ToInt64E(string(x))
	case string:
		var v int64
		if v, err = strconv.ParseInt(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to int64")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = int64(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to int64", x)
	}
	return
}

// ToInts convert an interface to a []int type
func ToInts(obj interface{}) []int {
	x, _ := ToIntSliceE(obj)
	if x == nil {
		return []int{}
	}
	return x.G()
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

	// FloatSlice
	//----------------------------------------------------------------------------------------------
	case FloatSlice:
		for i := range x {
			*val = append(*val, int(x[i]))
		}
	case []FloatSlice:
		for i := range x {
			for j := range x[i] {
				*val = append(*val, int(x[i][j]))
			}
		}
	case []*FloatSlice:
		for i := range x {
			if x[i] != nil {
				for j := range *x[i] {
					*val = append(*val, int((*x[i])[j]))
				}
			}
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

	// InterSlice
	//----------------------------------------------------------------------------------------------
	case InterSlice:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}
	case []InterSlice:
		for i := range x {
			*val = append(*val, ToInt(x[i]))
		}
	case []*InterSlice:
		for i := range x {
			if x[i] != nil {
				for j := range *x[i] {
					*val = append(*val, ToInt((*x[i])[j]))
				}
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case Object:
		if v, e := x.ToIntE(); e != nil {
			err = errors.Errorf("unable to convert type Object to []int")
		} else {
			*val = append(*val, v)
		}
	case []Object:
		for i := range x {
			if v, e := x[i].ToIntE(); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type %T to []int", x[i].O())
				return
			}
		}
	case []*Object:
		for i := range x {
			if v, e := x[i].ToIntE(); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type %T to []int", x[i].O())
				return
			}
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		if v, e := ToIntE(x); e != nil {
			err = errors.Errorf("unable to convert type Str to []int")
		} else {
			*val = append(*val, v)
		}
	case []Str:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type Str to []int")
				return
			}
		}
	case []*Str:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type *Str to []int")
				return
			}
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string:
		if v, e := ToIntE(x); e != nil {
			err = errors.Errorf("unable to convert type string to []int")
		} else {
			*val = append(*val, v)
		}
	case []string:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type string to []int")
				return
			}
		}
	case []*string:
		for i := range x {
			if v, e := ToIntE(x[i]); e == nil {
				*val = append(*val, v)
			} else {
				err = errors.Errorf("unable to convert type *string to []int")
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
			err = errors.Errorf("unable to convert type %T to []int", x)
			return
		}
	}
	return
}

// ToSlice is an alias to ToInterSlice
func ToSlice(obj interface{}) (slice *InterSlice) {
	return ToInterSlice(obj)
}

// ToInterSlice converts the given slice to an *InterSlice
func ToInterSlice(obj interface{}) (slice *InterSlice) {
	slice = &InterSlice{}

	// Optimized types
	switch x := obj.(type) {
	case nil:

	// bool
	//----------------------------------------------------------------------------------------------
	case bool, *bool:
		*slice = append(*slice, x)
	case []bool:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]bool:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*bool:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*bool:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// Char
	//----------------------------------------------------------------------------------------------
	case Char, *Char:
		*slice = append(*slice, x)
	case []Char:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]Char:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*Char:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*Char:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// float32
	//----------------------------------------------------------------------------------------------
	case float32, *float32:
		*slice = append(*slice, x)
	case []float32:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]float32:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*float32:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*float32:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// float64
	//----------------------------------------------------------------------------------------------
	case float64, *float64:
		*slice = append(*slice, x)
	case []float64:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]float64:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*float64:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*float64:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// FloatSlice
	//----------------------------------------------------------------------------------------------
	case FloatSlice:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *FloatSlice:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []FloatSlice:
		for i := range x {
			for j := range x[i] {
				*slice = append(*slice, x[i][j])
			}
		}
	case *[]FloatSlice:
		if x != nil {
			for i := range *x {
				for j := range (*x)[i] {
					*slice = append(*slice, (*x)[i][j])
				}
			}
		}
	case []*FloatSlice:
		for i := range x {
			if x[i] != nil {
				for j := range *x[i] {
					*slice = append(*slice, (*x[i])[j])
				}
			}
		}
	case *[]*FloatSlice:
		if x != nil {
			for i := range *x {
				if (*x)[i] != nil {
					for j := range *(*x)[i] {
						*slice = append(*slice, (*(*x)[i])[j])
					}
				}
			}
		}

	// interface
	//----------------------------------------------------------------------------------------------
	case []interface{}:
		val := InterSlice(x)
		slice = &val

	case *[]interface{}:
		if x != nil {
			val := InterSlice(*x)
			slice = &val
		}

	// int
	//----------------------------------------------------------------------------------------------
	case int, *int:
		*slice = append(*slice, x)
	case []int:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]int:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*int:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*int:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// int8
	//----------------------------------------------------------------------------------------------
	case int8, *int8:
		*slice = append(*slice, x)
	case []int8:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]int8:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*int8:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*int8:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// int16
	//----------------------------------------------------------------------------------------------
	case int16, *int16:
		*slice = append(*slice, x)
	case []int16:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]int16:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*int16:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*int16:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// int32
	//----------------------------------------------------------------------------------------------
	case int32, *int32:
		*slice = append(*slice, x)
	case []int32:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]int32:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*int32:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*int32:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// int64
	//----------------------------------------------------------------------------------------------
	case int64, *int64:
		*slice = append(*slice, x)
	case []int64:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]int64:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*int64:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*int64:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// IntSlice
	//----------------------------------------------------------------------------------------------
	case IntSlice:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *IntSlice:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []IntSlice:
		for i := range x {
			for j := range x[i] {
				*slice = append(*slice, x[i][j])
			}
		}
	case []*IntSlice:
		for i := range x {
			if x[i] != nil {
				for j := range *x[i] {
					*slice = append(*slice, (*x[i])[j])
				}
			}
		}

	// InterSlice
	//----------------------------------------------------------------------------------------------
	case InterSlice:
		slice = &x
	case *InterSlice:
		if x != nil {
			slice = x
		}
	case []InterSlice:
		for i := range x {
			*slice = append(*slice, x[i]...)
		}
	case *[]InterSlice:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i]...)
			}
		}
	case []*InterSlice:
		for i := range x {
			if x[i] != nil {
				*slice = append(*slice, (*x[i])...)
			}
		}
	case *[]*InterSlice:
		if x != nil {
			for i := range *x {
				if (*x)[i] != nil {
					*slice = append(*slice, (*(*x)[i])...)
				}
			}
		}

	// Object
	//----------------------------------------------------------------------------------------------
	case Object, *Object:
		*slice = append(*slice, x)
	case []Object:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]Object:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*Object:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*Object:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// Str
	//----------------------------------------------------------------------------------------------
	case Str:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *Str:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []Str:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]Str:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*Str:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*Str:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// string
	//----------------------------------------------------------------------------------------------
	case string, *string:
		*slice = append(*slice, x)
	case []string:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]string:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*string:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*string:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// map[string]
	//----------------------------------------------------------------------------------------------
	case map[string]string, *map[string]string:
		*slice = append(*slice, x)
	case []map[string]string:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]map[string]string:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*map[string]string:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*map[string]string:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case map[string]interface{}, *map[string]interface{}:
		*slice = append(*slice, x)
	case []map[string]interface{}:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]map[string]interface{}:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*map[string]interface{}:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*map[string]interface{}:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// StringMap
	//----------------------------------------------------------------------------------------------
	case StringMap:
		slice = ToInterSlice(x.G())
	case *StringMap:
		slice = ToInterSlice(x.G())
	case []StringMap:
		for i := range x {
			*slice = append(*slice, *ToInterSlice(x[i].G())...)
		}
	case *[]StringMap:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, *ToInterSlice((*x)[i].G())...)
			}
		}
	case []*StringMap:
		for i := range x {
			*slice = append(*slice, *ToInterSlice(x[i].G())...)
		}
	case *[]*StringMap:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, *ToInterSlice((*x)[i].G())...)
			}
		}

	// StringSlice
	//----------------------------------------------------------------------------------------------
	case StringSlice:
		slice = ToInterSlice(x.G())
	case *StringSlice:
		slice = ToInterSlice(x.G())
	case []StringSlice:
		for i := range x {
			*slice = append(*slice, *ToInterSlice(x[i].G())...)
		}
	case *[]StringSlice:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, *ToInterSlice((*x)[i].G())...)
			}
		}
	case []*StringSlice:
		for i := range x {
			*slice = append(*slice, *ToInterSlice(x[i].G())...)
		}
	case *[]*StringSlice:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, *ToInterSlice((*x)[i].G())...)
			}
		}

	// uint
	//----------------------------------------------------------------------------------------------
	case uint, *uint:
		*slice = append(*slice, x)
	case []uint:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]uint:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*uint:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*uint:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// uint8
	//----------------------------------------------------------------------------------------------
	case uint8, *uint8:
		*slice = append(*slice, x)
	case []uint8:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]uint8:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*uint8:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*uint8:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// uint16
	//----------------------------------------------------------------------------------------------
	case uint16, *uint16:
		*slice = append(*slice, x)
	case []uint16:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]uint16:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*uint16:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*uint16:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// uint32
	//----------------------------------------------------------------------------------------------
	case uint32, *uint32:
		*slice = append(*slice, x)
	case []uint32:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]uint32:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*uint32:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*uint32:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// uint64
	//----------------------------------------------------------------------------------------------
	case uint64, *uint64:
		*slice = append(*slice, x)
	case []uint64:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]uint64:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}
	case []*uint64:
		for i := range x {
			*slice = append(*slice, x[i])
		}
	case *[]*uint64:
		if x != nil {
			for i := range *x {
				*slice = append(*slice, (*x)[i])
			}
		}

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(x)
		k := v.Kind()

		switch {

		// generically convert array and *slice types
		case k == reflect.Array || k == reflect.Slice:
			for i := 0; i < v.Len(); i++ {
				*slice = append(*slice, v.Index(i).Interface())
			}

		// everything else just drop in the slice directly
		default:
			*slice = append(*slice, x)
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

	// InterSlice
	//----------------------------------------------------------------------------------------------
	case *InterSlice:
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
	case *[]InterSlice:
		if x != nil {
			for i := range *x {
				for _, raw := range (*x)[i] {
					var m *MapSlice
					if m, err = ToMapSliceE(raw); err != nil {
						return
					}
					for _, v := range *m {
						*val = append(*val, v)
					}
				}
			}
		}
	case *[]*InterSlice:
		if x != nil {
			for i := range *x {
				if (*x)[i] != nil {
					for _, raw := range *((*x)[i]) {
						var m *MapSlice
						if m, err = ToMapSliceE(raw); err != nil {
							return
						}
						for _, v := range *m {
							*val = append(*val, v)
						}
					}
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

// ToStringMap converts an interface to a StringMap type. Supports converting yaml string as well.
func ToStringMap(obj interface{}) *StringMap {
	x, _ := ToStringMapE(obj)
	if x == nil {
		return &StringMap{}
	}
	return x
}

// ToStringMapE converts an interface to a StringMap type. Supports converting yaml string as well.
// Specifically restricting the number of conversions here to keep it in line with support YAML types.
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
	case *map[string]bool:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]float32:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]float64:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]int:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]int8:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]int16:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]int32:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]int64:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]uint:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]uint8:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]uint16:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]uint32:
		if x != nil {
			for k, v := range *x {
				val.Set(k, v)
			}
		}
	case *map[string]uint64:
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

	// fall back on reflection
	//----------------------------------------------------------------------------------------------
	default:
		v := reflect.ValueOf(DeReference(obj))
		k := v.Kind()

		switch {

		// generically convert map
		case k == reflect.Map:
			for _, k := range v.MapKeys() {
				val.Set(k.Interface(), v.MapIndex(k).Interface())
			}

		default:
			err = errors.Errorf("unable to convert type %T to a StringMap", x)
			return
		}
	}

	return
}

// ToStringMapG converts an interface to a map[string]interface{} type. Supports converting yaml string as well.
func ToStringMapG(obj interface{}) map[string]interface{} {
	x, _ := ToStringMapE(obj)
	if x == nil {
		return map[string]interface{}{}
	}
	return x.G()
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

	// FloatSlice
	//----------------------------------------------------------------------------------------------
	case FloatSlice:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []FloatSlice:
		for i := range x {
			for j := range x[i] {
				*val = append(*val, ToString(x[i][j]))
			}
		}
	case []*FloatSlice:
		for i := range x {
			for j := range *x[i] {
				*val = append(*val, ToString((*x[i])[j]))
			}
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

	// InterSlice
	//----------------------------------------------------------------------------------------------
	case InterSlice:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []InterSlice:
		for i := range x {
			*val = append(*val, ToString(x[i]))
		}
	case []*InterSlice:
		for i := range x {
			if x[i] != nil {
				for j := range *x[i] {
					*val = append(*val, ToString((*x[i])[j]))
				}
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
			err = errors.Errorf("unable to convert type %T to []int", x)
			return
		}
	}
	return
}

// ToStrs convert an interface to a []string type.
func ToStrs(obj interface{}) (val []string) {
	val, _ = ToStrsE(obj)
	return val
}

// ToStrsE convert an interface to a []string type.
func ToStrsE(obj interface{}) (val []string, err error) {
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
		val, err = ToStrsE(ToString(x))
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
		val, err = ToStrsE(ToString(x))
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
			err = errors.Errorf("unable to convert type %T to []int", x)
			return
		}
	}
	return
}

// ToTime converts an interface to a time.Time type, invalid types will simply return the default time.Time
func ToTime(obj interface{}) time.Time {
	x, _ := ToTimeE(obj)
	return x
}

// ToTimeE converts an interface to a time.Time type.
func ToTimeE(obj interface{}) (val time.Time, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:

	// time.Time
	//----------------------------------------------------------------------------------------------
	case time.Time:
		val = x

	// string
	// Parse the string trying popular layouts
	//----------------------------------------------------------------------------------------------
	case string:

		// Try int conversion first
		if v, e := strconv.ParseInt(x, 0, 0); e == nil {
			return ToTimeE(v)
		}

		// Parse string time formats
		for _, layout := range TimeLayouts {
			if val, err = time.Parse(layout, x); err == nil {
				return val, nil
			}
		}
		err = errors.Errorf("failed to parse time %s", x)

	// int
	//----------------------------------------------------------------------------------------------
	case int:
		val = time.Unix(int64(x), 0)
	case int32:
		val = time.Unix(int64(x), 0)
	case int64:
		val = time.Unix(x, 0)

	// uint
	//----------------------------------------------------------------------------------------------
	case uint:
		val = time.Unix(int64(x), 0)
	case uint32:
		val = time.Unix(int64(x), 0)
	case uint64:
		val = time.Unix(int64(x), 0)
	default:
		err = errors.Errorf("failed to convert type %T to time.Time", obj)
	}

	// Use UTC if set to false
	if !gUseLocalTime {
		val = val.UTC()
	}

	return
}

// ToUint convert an interface to an uint type.
func ToUint(obj interface{}) uint {
	x, _ := ToUintE(obj)
	return x
}

// ToUintE convert an interface to an uint type.
func ToUintE(obj interface{}) (val uint, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToUintE(x.A())
	case float32:
		val = uint(x)
	case float64:
		val = uint(x)
	case int:
		val = uint(x)
	case int8:
		val = uint(x)
	case int16:
		val = uint(x)
	case int32:
		val = uint(x)
	case int64:
		val = uint(x)
	case Object:
		val, err = x.ToUintE()
	case uint:
		val = x
	case uint8:
		val = uint(x)
	case uint16:
		val = uint(x)
	case uint32:
		val = uint(x)
	case uint64:
		val = uint(x)
	case Str:
		return ToUintE(string(x))
	case string:
		var v uint64
		if v, err = strconv.ParseUint(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to uint")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = uint(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to uint", x)
	}
	return
}

// ToUint8 convert an interface to an uint8 type.
func ToUint8(obj interface{}) uint8 {
	x, _ := ToUint8E(obj)
	return x
}

// ToUint8E convert an interface to an uint8 type.
func ToUint8E(obj interface{}) (val uint8, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToUint8E(x.A())
	case float32:
		val = uint8(x)
	case float64:
		val = uint8(x)
	case int:
		val = uint8(x)
	case int8:
		val = uint8(x)
	case int16:
		val = uint8(x)
	case int32:
		val = uint8(x)
	case int64:
		val = uint8(x)
	case Object:
		val, err = x.ToUint8E()
	case uint:
		val = uint8(x)
	case uint8:
		val = x
	case uint16:
		val = uint8(x)
	case uint32:
		val = uint8(x)
	case uint64:
		val = uint8(x)
	case Str:
		return ToUint8E(string(x))
	case string:
		var v uint64
		if v, err = strconv.ParseUint(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to uint8")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = uint8(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to uint8", x)
	}
	return
}

// ToUint16 convert an interface to an uint16 type.
func ToUint16(obj interface{}) uint16 {
	x, _ := ToUint16E(obj)
	return x
}

// ToUint16E convert an interface to an uint16 type.
func ToUint16E(obj interface{}) (val uint16, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToUint16E(x.A())
	case float32:
		val = uint16(x)
	case float64:
		val = uint16(x)
	case int:
		val = uint16(x)
	case int8:
		val = uint16(x)
	case int16:
		val = uint16(x)
	case int32:
		val = uint16(x)
	case int64:
		val = uint16(x)
	case Object:
		val, err = x.ToUint16E()
	case uint:
		val = uint16(x)
	case uint8:
		val = uint16(x)
	case uint16:
		val = x
	case uint32:
		val = uint16(x)
	case uint64:
		val = uint16(x)
	case Str:
		return ToUint16E(string(x))
	case string:
		var v uint64
		if v, err = strconv.ParseUint(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to uint16")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = uint16(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to uint16", x)
	}
	return
}

// ToUint32 convert an interface to an uint32 type.
func ToUint32(obj interface{}) uint32 {
	x, _ := ToUint32E(obj)
	return x
}

// ToUint32E convert an interface to an uint32 type.
func ToUint32E(obj interface{}) (val uint32, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToUint32E(x.A())
	case float32:
		val = uint32(x)
	case float64:
		val = uint32(x)
	case int:
		val = uint32(x)
	case int8:
		val = uint32(x)
	case int16:
		val = uint32(x)
	case int32:
		val = uint32(x)
	case int64:
		val = uint32(x)
	case Object:
		val, err = x.ToUint32E()
	case uint:
		val = uint32(x)
	case uint8:
		val = uint32(x)
	case uint16:
		val = uint32(x)
	case uint32:
		val = x
	case uint64:
		val = uint32(x)
	case Str:
		return ToUint32E(string(x))
	case string:
		var v uint64
		if v, err = strconv.ParseUint(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to uint32")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = uint32(v)
		}
	default:
		err = errors.Errorf("unable to convert type %T to uint32", x)
	}
	return
}

// ToUint64 convert an interface to an uint64 type.
func ToUint64(obj interface{}) uint64 {
	x, _ := ToUint64E(obj)
	return x
}

// ToUint64E convert an interface to an uint64 type.
func ToUint64E(obj interface{}) (val uint64, err error) {
	o := DeReference(obj)

	switch x := o.(type) {
	case nil:
	case bool:
		if x {
			val = 1
		}
	case Char:
		val, err = ToUint64E(x.A())
	case float32:
		val = uint64(x)
	case float64:
		val = uint64(x)
	case int:
		val = uint64(x)
	case int8:
		val = uint64(x)
	case int16:
		val = uint64(x)
	case int32:
		val = uint64(x)
	case int64:
		val = uint64(x)
	case Object:
		val, err = x.ToUint64E()
	case uint:
		val = uint64(x)
	case uint8:
		val = uint64(x)
	case uint16:
		val = uint64(x)
	case uint32:
		val = uint64(x)
	case uint64:
		val = x
	case Str:
		return ToUint64E(string(x))
	case string:
		var v uint64
		if v, err = strconv.ParseUint(x, 10, 64); err != nil {
			err = errors.Wrapf(err, "failed to convert string to uint64")

			// Also convert true|false|TRUE|FALSE
			if b, e := strconv.ParseBool(x); e == nil {
				err = nil
				if b {
					val = 1
				}
			}
		} else {
			val = v
		}
	default:
		err = errors.Errorf("unable to convert type %T to uint64", x)
	}
	return
}

// YAMLCont checks if the given value is a valid YAML container
func YAMLCont(obj interface{}) bool {
	switch obj.(type) {
	case map[string]interface{}, *StringMap:
		return true
	case []interface{}, []string, []int:
		return true
	default:
		return false
	}
}
