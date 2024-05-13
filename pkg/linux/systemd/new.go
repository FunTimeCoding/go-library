package systemd

import "github.com/funtimecoding/go-library/pkg/ssh"

func New(c *ssh.Client) *Client {
	return &Client{ssh: c}
}
