package parse

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/result/time_string"
	"github.com/prometheus/common/model"
	"log"
)

func MatrixTimeString(v model.Value) []*time_string.Result {
	var result []*time_string.Result
	matrix, okay := v.(model.Matrix)

	if !okay {
		log.Panicf("not a matrix")
	}

	for _, element := range matrix {
		for _, value := range element.Values {
			result = append(
				result,
				&time_string.Result{
					Metric: element.Metric.String(),
					Time:   value.Timestamp.Time(),
					Value:  value.Value.String(),
				},
			)
		}
	}

	return result
}
