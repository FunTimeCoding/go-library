package reflects

import "reflect"

func Default(a any) bool {
	switch v := a.(type) {
	case bool:
		return !v
	case int:
		return v == 0
	case string:
		return v == ""
	case *string:
		return v == nil
	case *int:
		return v == nil
	default:
		l := reflect.ValueOf(a)

		if l.Kind() == reflect.Ptr {
			return l.IsNil()
		}

		return false
	}
}
