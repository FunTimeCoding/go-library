package parse

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/result/generic"
	"github.com/prometheus/common/model"
	"log"
)

func Generic(v model.Value) []*generic.Result {
	var result []*generic.Result

	switch t := v.Type(); t {
	case model.ValMatrix:
		for _, element := range v.(model.Matrix) {
			for _, inner := range element.Values {
				result = append(
					result,
					&generic.Result{
						Type:   constant.Matrix,
						Metric: element.Metric.String(),
						Time:   inner.Timestamp.Time(),
						Value:  inner.Value.String(),
					},
				)
			}
		}
	case model.ValVector:
		for _, element := range v.(model.Vector) {
			result = append(
				result,
				&generic.Result{
					Type:   constant.Vector,
					Metric: element.Metric.String(),
					Time:   element.Timestamp.Time(),
					Value:  element.Value.String(),
				},
			)
		}
	case model.ValScalar:
		s := v.(*model.Scalar)
		result = append(
			result,
			&generic.Result{
				Type:   constant.Scalar,
				Metric: s.String(),
				Time:   s.Timestamp.Time(),
				Value:  s.Value.String(),
			},
		)
	case model.ValString:
		s := v.(*model.String)
		result = append(
			result,
			&generic.Result{
				Type:   constant.String,
				Metric: s.String(),
				Time:   s.Timestamp.Time(),
				Value:  s.Value,
			},
		)
	default:
		log.Panicf("unexpected type: %d", t)
	}

	return result
}
