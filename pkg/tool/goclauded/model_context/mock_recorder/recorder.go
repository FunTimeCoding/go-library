package mock_recorder

import "github.com/funtimecoding/go-library/pkg/telemetry/record"

type Recorder struct {
	Calls []*record.Record
}
