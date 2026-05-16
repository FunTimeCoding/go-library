package secure_shell

import (
	"errors"
	"golang.org/x/crypto/ssh"
)

func Exit(e error) int {
	if x, okay := errors.AsType[*ssh.ExitError](e); okay {
		return x.ExitStatus()
	}

	return 0
}
