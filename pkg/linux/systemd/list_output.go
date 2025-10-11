package systemd

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) ListOutput() string {
	return c.ssh.Run(
		join.Space(
			constant.Command,
			constant.ListUnits,
			constant.Type,
			constant.Service,
			constant.All,
			constant.Full,
			constant.Plain,
			constant.NoLegend,
		),
	).OutputString
}
