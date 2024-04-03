package markup

import "reflect"

func RemoveEmpty(input any) any {
	switch result := input.(type) {
	case map[string]any:
		for k, v := range result {
			cleanV := RemoveEmpty(v)

			if cleanV == nil ||
				(reflect.ValueOf(cleanV).Kind() == reflect.Map &&
					len(cleanV.(map[string]any)) == 0) ||
				(reflect.ValueOf(cleanV).Kind() == reflect.String &&
					len(cleanV.(string)) == 0) {
				delete(result, k)
			} else {
				result[k] = cleanV
			}
		}

		return result
	case []any:
		for i, v := range result {
			result[i] = RemoveEmpty(v)
		}

		return result
	default:
		return input
	}
}
