package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"golang.org/x/crypto/ssh"
	"net"
)

func Listen(
	c *ssh.Client,
	address string,
) net.Listener {
	result, e := c.Listen("tcp", address)
	errors.PanicOnError(e)

	return result
}
