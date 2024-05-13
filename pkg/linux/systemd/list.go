package systemd

import (
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) List() *ssh.Result {
	result := c.ssh.Run(
		join.Space(
			Command,
			ListUnits,
			Type,
			Service,
			All,
			Full,
			Plain,
			NoLegend,
		),
	)

	return result
}
