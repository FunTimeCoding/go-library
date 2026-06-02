package convert

import "github.com/prometheus/common/model"

func labelSetToMap(v model.LabelSet) map[string]string {
	if len(v) == 0 {
		return nil
	}

	result := make(map[string]string, len(v))

	for k, a := range v {
		result[string(k)] = string(a)
	}

	return result
}
