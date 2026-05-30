package record

import "github.com/funtimecoding/go-library/pkg/telemetry/constant"

func NewBaseline(
	tool string,
	surface string,
	actor string,
	outcome string,
) *Record {
	return newRecord(tool, surface, actor, outcome, constant.Baseline)
}
