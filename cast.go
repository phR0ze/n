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
	o := Indirect(obj)

	switch x := o.(type) {
	case nil:
		return ""
	case string:
		return x
	case bool:
		return strconv.FormatBool(x)
	case []byte:
		return string(x)
	case Char:
		return x.String()
	case error:
		return x.Error()
	case float32:
		return strconv.FormatFloat(float64(x), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(x), 10)
	case int8:
		return strconv.FormatInt(int64(x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case int64:
		return strconv.FormatInt(x, 10)
	case template.CSS:
		return string(x)
	case template.HTML:
		return string(x)
	case template.HTMLAttr:
		return string(x)
	case template.JS:
		return string(x)
	case template.URL:
		return string(x)
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

// ToInt casts an interface to an int type.
func ToInt(obj interface{}) int {
	x, _ := ToIntE(obj)
	return x
}

// ToIntE casts an interface to an int type.
func ToIntE(obj interface{}) (int, error) {
	o := Indirect(obj)

	switch x := o.(type) {
	case int:
		return x, nil
	case int8:
		return int(x), nil
	case int16:
		return int(x), nil
	case int32:
		return int(x), nil
	case int64:
		return int(x), nil
	case uint:
		return int(x), nil
	case uint64:
		return int(x), nil
	case uint32:
		return int(x), nil
	case uint16:
		return int(x), nil
	case uint8:
		return int(x), nil
	case float64:
		return int(x), nil
	case float32:
		return int(x), nil
	case string:
		v, err := strconv.ParseInt(x, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to cast type %T to int", obj)
	case bool:
		if x {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast type %T to int", obj)
	}
}
