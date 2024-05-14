package systemd

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/show"
)

func (c *Client) IsActiveRunning(name string) bool {
	s := show.Parse(show.OutputToMap(c.ShowOutput(name)))

	if s.ActiveState == constant.ActiveState &&
		s.SubState == constant.RunningSubState {
		return true
	}

	return false
}
