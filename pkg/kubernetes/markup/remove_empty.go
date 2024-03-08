package markup

import "reflect"

func RemoveEmpty(input any) any {
	switch val := input.(type) {
	case map[string]any:
		for k, v := range val {
			cleanedV := RemoveEmpty(v)
			if cleanedV == nil ||
				(reflect.ValueOf(cleanedV).Kind() == reflect.Map &&
					len(cleanedV.(map[string]any)) == 0) {
				delete(val, k)
			} else {
				val[k] = cleanedV
			}
		}

		return val
	case []any:
		for i, v := range val {
			val[i] = RemoveEmpty(v)
		}

		return val
	default:
		return input
	}
}
