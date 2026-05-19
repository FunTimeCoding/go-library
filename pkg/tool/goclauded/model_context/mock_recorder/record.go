package mock_recorder

import "github.com/funtimecoding/go-library/pkg/telemetry/record"

func (r *Recorder) Record(e *record.Record) {
	r.Calls = append(r.Calls, e)
}
