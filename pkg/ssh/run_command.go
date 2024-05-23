package ssh

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ssh/command"
	"github.com/funtimecoding/go-library/pkg/strings/trim"
	"github.com/funtimecoding/go-library/pkg/system/secure_shell"
)

func (c *Client) RunCommand(o *command.Command) *Result {
	s := secure_shell.Session(c.client)
	defer secure_shell.CloseSession(s)
	stdout, stderr := secure_shell.SessionBuffers(s)

	for k, v := range o.Environment {
		errors.PanicOnError(s.Setenv(k, v))
	}

	e := s.Run(o.Command)

	return &Result{
		OutputString: trim.NewLine(stdout.String()),
		ErrorString:  trim.NewLine(stderr.String()),
		ExitCode:     secure_shell.ExitCode(e),
		Error:        e,
	}
}
