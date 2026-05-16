package query

import "github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/query_result"

type Query struct {
	Status string               `json:"status"`
	Result *query_result.Result `json:"data"`
}
