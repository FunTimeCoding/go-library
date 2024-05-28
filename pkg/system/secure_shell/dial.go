package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"golang.org/x/crypto/ssh"
)

func Dial(
	address string,
	c *ssh.ClientConfig,
) *ssh.Client {
	result, e := ssh.Dial(system.Transmission, address, c)
	errors.PanicOnError(e)

	return result
}
