package helper

import "github.com/prometheus/common/model"

func LabelValuesToStrings(v model.LabelValues) []string {
	var result []string

	for _, label := range v {
		result = append(result, string(label))
	}

	return result
}
