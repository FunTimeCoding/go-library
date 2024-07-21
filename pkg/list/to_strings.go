package list

import (
	"fmt"
	"reflect"
)

func ToStrings(a any) []string {
	var result []string
	v := reflect.ValueOf(a)

	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			if s, okay := v.Index(i).Interface().(fmt.Stringer); okay {
				result = append(result, s.String())
			}
		}
	}

	return result
}
