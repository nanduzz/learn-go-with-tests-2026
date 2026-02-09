package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResults := val.Call(nil)
		for _, result := range valFnResults {
			walkValue(result)
		}
	}

}

// getValue is a helper function that takes an interface{} and returns a reflect.Value.
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// If the value is a pointer, we need to get the value it points to before we can work with it.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
