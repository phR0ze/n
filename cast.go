package n

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
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
	case Char:
		return x
	case *Char:
		if x == nil {
			return *NewChar("")
		}
		return *x
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
			return *NewStrV()
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
	case template.CSS:
		return x
	case *template.CSS:
		if x == nil {
			return template.CSS("")
		}
		return *x
	case template.HTML:
		return x
	case *template.HTML:
		if x == nil {
			return template.HTML("")
		}
		return *x
	case template.HTMLAttr:
		return x
	case *template.HTMLAttr:
		if x == nil {
			return template.HTMLAttr("")
		}
		return *x
	case template.JS:
		return x
	case *template.JS:
		if x == nil {
			return template.JS("")
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

// ToBool casts an interface to a bool type.
func ToBool(obj interface{}) bool {
	val, _ := ToBoolE(obj)
	return val
}

// ToBoolE casts an interface to a bool type.
func ToBoolE(obj interface{}) (val bool, err error) {
	switch x := obj.(type) {
	case nil:
	case bool, *bool:
		val = Indirect(x).(bool)
	case int, *int:
		val = Indirect(x).(int) != 0
	case int8, *int8:
		val = Indirect(x).(int8) != 0
	case int16, *int16:
		val = Indirect(x).(int16) != 0
	case int32, *int32:
		val = Indirect(x).(int32) != 0
	case int64, *int64:
		val = Indirect(x).(int64) != 0
	case string, *string:
		val, err = strconv.ParseBool(Indirect(x).(string))
	case uint, *uint:
		val = Indirect(x).(uint) != 0
	case uint8, *uint8:
		val = Indirect(x).(uint8) != 0
	case uint16, *uint16:
		val = Indirect(x).(uint16) != 0
	case uint32, *uint32:
		val = Indirect(x).(uint32) != 0
	case uint64, *uint64:
		val = Indirect(x).(uint64) != 0
	default:
		err = errors.Errorf("unable to cast type %T to bool", x)
	}
	return
}

// ToString casts an interface to a string type.
func ToString(obj interface{}) string {
	switch x := obj.(type) {
	case nil:
		return ""
	case string, *string:
		return Indirect(x).(string)
	case bool, *bool:
		return strconv.FormatBool(Indirect(x).(bool))
	case []byte, *[]byte:
		return string(Indirect(x).([]byte))
	case Char, *Char:
		c := Indirect(x).(Char)
		return c.String()
	case error:
		return x.Error()
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
	case template.CSS, *template.CSS:
		return string(Indirect(x).(template.CSS))
	case template.HTML, *template.HTML:
		return string(Indirect(x).(template.HTML))
	case template.HTMLAttr, *template.HTMLAttr:
		return string(Indirect(x).(template.HTMLAttr))
	case template.JS, *template.JS:
		return string(Indirect(x).(template.JS))
	case template.URL, *template.URL:
		return string(Indirect(x).(template.URL))
	case []rune, *[]rune:
		return string(Indirect(x).([]rune))
	case uint, *uint:
		return strconv.FormatInt(int64(Indirect(x).(uint)), 10)
	case uint8, *uint8:
		return strconv.FormatInt(int64(Indirect(x).(uint8)), 10)
	case uint16, *uint16:
		return strconv.FormatInt(int64(Indirect(x).(uint16)), 10)
	case uint32, *uint32:
		return strconv.FormatInt(int64(Indirect(x).(uint32)), 10)
	case uint64, *uint64:
		return strconv.FormatInt(int64(Indirect(x).(uint64)), 10)
	case fmt.Stringer, *fmt.Stringer:
		if x, ok := x.(*fmt.Stringer); ok {
			if x == nil {
				return ""
			}
			return (*x).String()
		}
		return x.(fmt.Stringer).String()
	default:
		return fmt.Sprintf("%v", obj)
	}
}
