package systemd

import (
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) NotFound() *ssh.Result {
	result := c.ssh.Run(join.Space(Command, State, NotFound, NoLegend))

	return result
}
