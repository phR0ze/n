package n

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
)

// Indirect dereferences the interface if needed returning a non-pointer type
func Indirect(obj interface{}) interface{} {
	switch x := obj.(type) {
	case nil:
		return x
	case bool:
		return x
	case *bool:
		if x == nil {
			return false
		}
		return *x
	// []byte is just a uint8[] which is defined later
	// case []byte:
	// 	return x
	// case *[]byte:
	// 	if x == nil {
	// 		return []byte{}
	// 	}
	// 	return *x
	case float32:
		return x
	case *float32:
		if x == nil {
			return float32(0)
		}
		return *x
	case float64:
		return x
	case *float64:
		if x == nil {
			return float64(0)
		}
		return *x
	case int:
		return x
	case *int:
		if x == nil {
			return 0
		}
		return *x
	case int8:
		return x
	case *int8:
		if x == nil {
			return int8(0)
		}
		return *x
	case int16:
		return x
	case *int16:
		if x == nil {
			return int16(0)
		}
		return *x
	case int32:
		return x
	case *int32:
		if x == nil {
			return int32(0)
		}
		return *x
	case int64:
		return x
	case *int64:
		if x == nil {
			return int64(0)
		}
		return *x
	case []int:
		return x
	case *[]int:
		if x == nil {
			return []int{}
		}
		return *x
	case *[]int8:
		if x == nil {
			return []int8{}
		}
		return *x
	case []int16:
		return x
	case *[]int16:
		if x == nil {
			return []int16{}
		}
		return *x
	case []int32:
		return x
	case *[]int32:
		if x == nil {
			return []int32{}
		}
		return *x
	case []int64:
		return x
	case *[]int64:
		if x == nil {
			return []int64{}
		}
		return *x
	case IntSlice:
		return x
	case *IntSlice:
		if x == nil {
			return *NewIntSliceV()
		}
		return *x
	case Str:
		return x
	case *Str:
		if x == nil {
			return *NewStr("")
		}
		return *x
	// rune is a int32 which is already defined
	// case rune:
	// 	return x
	// case *rune:
	// 	if x == nil {
	// 		return rune(0)
	// 	}
	case StringSlice:
		return x
	case *StringSlice:
		if x == nil {
			return *NewStringSliceV()
		}
		return *x
	case string:
		return x
	case *string:
		if x == nil {
			return ""
		}
		return *x
	case []string:
		return x
	case *[]string:
		if x == nil {
			return []string{}
		}
		return *x
	case template.HTML:
		return x
	case *template.HTML:
		if x == nil {
			return template.HTML("")
		}
		return *x
	case template.URL:
		return x
	case *template.URL:
		if x == nil {
			return template.URL("")
		}
		return *x
	case uint:
		return x
	case *uint:
		if x == nil {
			return uint(0)
		}
		return *x
	case uint8:
		return x
	case *uint8:
		if x == nil {
			return uint8(0)
		}
		return *x
	case uint16:
		return x
	case *uint16:
		if x == nil {
			return uint16(0)
		}
		return *x
	case uint32:
		return x
	case *uint32:
		if x == nil {
			return uint32(0)
		}
		return *x
	case uint64:
		return x
	case *uint64:
		if x == nil {
			return uint64(0)
		}
		return *x
	case []uint:
		return x
	case *[]uint:
		if x == nil {
			return []uint{}
		}
		return *x
	case []uint8:
		return x
	case *[]uint8:
		if x == nil {
			return []uint8{}
		}
		return *x
	case []uint16:
		return x
	case *[]uint16:
		if x == nil {
			return []uint16{}
		}
		return *x
	case []uint32:
		return x
	case *[]uint32:
		if x == nil {
			return []uint32{}
		}
		return *x
	case []uint64:
		return x
	case *[]uint64:
		if x == nil {
			return []uint64{}
		}
		return *x
	default:
		return reflect.Indirect(reflect.ValueOf(obj)).Interface()
	}
}

// Cast functions
//--------------------------------------------------------------------------------------------------

// ToString casts an interface to a string type.
func ToString(obj interface{}) string {
	switch x := obj.(type) {
	case string, *string:
		return Indirect(x).(string)
	case bool, *bool:
		return strconv.FormatBool(Indirect(x).(bool))
	case []byte, *[]byte:
		return string(Indirect(x).([]byte))
	case float32, *float32:
		return strconv.FormatFloat(float64(Indirect(x).(float32)), 'f', -1, 32)
	case float64, *float64:
		return strconv.FormatFloat(Indirect(x).(float64), 'f', -1, 64)
	case int, *int:
		return strconv.FormatInt(int64(Indirect(x).(int)), 10)
	case int8, *int8:
		return strconv.FormatInt(int64(Indirect(x).(int8)), 10)
	case int16, *int16:
		return strconv.FormatInt(int64(Indirect(x).(int16)), 10)
	case int32, *int32:
		return strconv.FormatInt(int64(Indirect(x).(int32)), 10)
	case int64, *int64:
		return strconv.FormatInt(Indirect(x).(int64), 10)
	case template.HTML, *template.HTML:
		return string(Indirect(x).(template.HTML))
	case template.URL, *template.URL:
		return string(Indirect(x).(template.URL))
	// case template.JS:
	// 	return string(s)
	// case template.CSS:
	// 	return string(s)
	// case template.HTMLAttr:
	// 	return string(s)
	// case fmt.Stringer:
	// 	return s.String()
	// case error:
	// 	return s.Error()
	case nil:
		return ""
	case []rune:
		return string(x)
	case uint:
		return strconv.FormatInt(int64(x), 10)
	case uint8:
		return strconv.FormatInt(int64(x), 10)
	case uint16:
		return strconv.FormatInt(int64(x), 10)
	case uint32:
		return strconv.FormatInt(int64(x), 10)
	case uint64:
		return strconv.FormatInt(int64(x), 10)
	default:
		return fmt.Sprintf("%v", obj)
	}
}
