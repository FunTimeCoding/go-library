package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"golang.org/x/crypto/ssh"
	"net"
)

func Listen(
	c *ssh.Client,
	address string,
) net.Listener {
	result, e := c.Listen(constant.Transmission, address)
	errors.PanicOnError(e)

	return result
}
