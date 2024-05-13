package ssh

import (
	"github.com/funtimecoding/go-library/pkg/strings/trim"
	"github.com/funtimecoding/go-library/pkg/system/secure_shell"
)

func (c *Client) Run(command string) *Result {
	s := secure_shell.Session(c.client)
	defer secure_shell.CloseSession(s)
	stdout, stderr := secure_shell.SessionBuffers(s)
	e := s.Run(command)

	return &Result{
		OutputString: trim.NewLine(stdout.String()),
		ErrorString:  trim.NewLine(stderr.String()),
		ExitCode:     secure_shell.ExitCode(e),
		Error:        e,
	}
}
