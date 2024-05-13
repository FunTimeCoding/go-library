package secure_shell

import (
	"errors"
	"golang.org/x/crypto/ssh"
)

func ExitCode(e error) int {
	var exit *ssh.ExitError

	if errors.As(e, &exit) {
		return exit.ExitStatus()
	}

	return 0
}
