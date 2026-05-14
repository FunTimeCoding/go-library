package face

import "github.com/funtimecoding/go-library/pkg/telemetry/record"

type Recorder interface {
	Record(r *record.Record)
}
