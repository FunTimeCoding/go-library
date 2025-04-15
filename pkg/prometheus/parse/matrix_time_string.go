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

	for _, e := range matrix {
		for _, v := range e.Values {
			result = append(
				result,
				&time_string.Result{
					Metric: e.Metric.String(),
					Time:   v.Timestamp.Time(),
					Value:  v.Value.String(),
				},
			)
		}
	}

	return result
}
