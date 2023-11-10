package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func Dial(
	address string,
	c *ssh.ClientConfig,
) *ssh.Client {
	result, e := ssh.Dial("tcp", address, c)
	errors.PanicOnError(e)

	return result
}
