package systemd

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) ShowOutput(name string) string {
	return c.ssh.Run(
		join.Space(constant.Command, constant.Show, name),
	).OutputString
}
