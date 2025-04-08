package reflection

import "reflect"

func walk(x any, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := range val.NumField() {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
