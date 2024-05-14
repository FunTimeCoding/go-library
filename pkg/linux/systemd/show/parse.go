package show

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/helper"
)

func Parse(m map[string]string) *Result {
	activeState := m["ActiveState"]
	subState := m["SubState"]
	activeEnter := helper.ParseTimestamp(m["ActiveEnterTimestamp"])
	execMainStart := helper.ParseTimestamp(m["ExecMainStartTimestamp"])

	return &Result{
		ActiveState:   activeState,
		SubState:      subState,
		ActiveEnter:   activeEnter,
		ExecMainStart: execMainStart,
	}
}
