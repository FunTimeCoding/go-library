package technitium

import "github.com/funtimecoding/go-library/pkg/technitium/record"

type recordsResponse struct {
	Records []*record.Record `json:"records"`
}
