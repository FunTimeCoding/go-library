package strings

import (
	"log"
	"reflect"
)

func AnyToSlice(a any) []string {
	switch a.(type) {
	case string:
		return []string{a.(string)}
	case []string:
		return a.([]string)
	case []any:
		var result []string

		for _, v := range a.([]any) {
			result = append(result, v.(string))
		}

		return result
	case nil:
		return []string{}
	default:
		log.Panicf("unexpected type: %s", reflect.TypeOf(a))

		return []string{}
	}
}
