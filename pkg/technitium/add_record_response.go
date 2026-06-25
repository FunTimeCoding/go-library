package technitium

import "github.com/funtimecoding/go-library/pkg/technitium/record"

type addRecordResponse struct {
	AddedRecord *record.Record `json:"addedRecord"`
}
