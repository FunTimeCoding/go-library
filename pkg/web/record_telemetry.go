package web

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/telemetry/constant"
	"github.com/funtimecoding/go-library/pkg/telemetry/record"
)

func RecordTelemetry(
	c face.Recorder,
	operation string,
	e error,
) {
	outcome := constant.Success

	if e != nil {
		outcome = constant.Error
	}

	c.Record(
		record.NewBaseline(
			operation,
			constant.WebService,
			"user",
			outcome,
		),
	)
}
