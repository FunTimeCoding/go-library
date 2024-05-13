package secure_shell

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func CloseSession(s *ssh.Session) {
	if e := s.Close(); e != nil && e.Error() != EndOfFile {
		errors.LogOnError(e)
	}
}
