package systemd

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) NotFoundOutput() string {
	return c.ssh.Run(
		join.Space(
			constant.Command,
			constant.State,
			constant.NotFound,
			constant.NoLegend,
		),
	).OutputString
}
