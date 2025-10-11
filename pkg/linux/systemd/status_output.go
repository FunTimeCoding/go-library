package systemd

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) StatusOutput(name string) string {
	return c.ssh.Run(
		join.Space(
			constant.Command,
			constant.Status,
			constant.Output,
			constant.Notation,
			name,
		),
	).OutputString
}
