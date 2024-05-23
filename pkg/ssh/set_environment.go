package ssh

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ssh/command"
	"golang.org/x/crypto/ssh"
)

func setEnvironment(
	s *ssh.Session,
	o *command.Command,
) {
	for k, v := range o.Environment {
		// TODO: Fails if the variable is not allowed in sshd via AcceptEnv
		if false {
			errors.PanicOnError(s.Setenv(k, v))
		}
	}
}
