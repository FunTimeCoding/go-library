package check_result

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/memory_activity"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pulse"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
)

type Result struct {
	Callsign       string
	Changed        bool
	Sessions       []session.Session
	Messages       []message.Message
	Pulses         []pulse.Pulse
	Completions    []completion.Completion
	MemoryActivity []memory_activity.Activity
	TimeoutMessage string
	Reannounce     bool
}
