package systemd

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/show"
)

func (c *Client) IsFailed(name string) bool {
	s := show.Parse(show.OutputToMap(c.ShowOutput(name)))

	if s.ActiveState == constant.FailedState &&
		s.SubState == constant.FailedSubState {
		return true
	}

	return false
}
