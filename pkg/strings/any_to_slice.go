package strings

import (
	"log"
	"reflect"
)

func AnyToSlice(a any) []string {
	switch t := a.(type) {
	case string:
		return []string{t}
	case []string:
		return t
	case []any:
		var result []string

		for _, v := range t {
			result = append(result, v.(string))
		}

		return result
	case nil:
		return []string{}
	default:
		log.Panicf("unexpected %s", reflect.TypeOf(a))

		return []string{}
	}
}
