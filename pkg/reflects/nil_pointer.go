package reflects

import "reflect"

func NilPointer(value any) bool {
	v := reflect.ValueOf(value)

	return v.Kind() == reflect.Pointer && v.IsNil()
}
