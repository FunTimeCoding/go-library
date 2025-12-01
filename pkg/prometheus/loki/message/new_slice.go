package message

import "github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"

func NewSlice(v response.Data) ([]*Message, *Meta) {
	var result []*Message

	for _, e := range v.Result {
		for _, a := range e.Values {
			result = append(result, New(a, &e.Stream))
		}
	}

	return result, &Meta{Type: v.ResultType, Statistic: v.Stats}
}
