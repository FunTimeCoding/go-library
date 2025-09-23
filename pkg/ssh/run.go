package ssh

import (
	"github.com/funtimecoding/go-library/pkg/ssh/result"
	"github.com/funtimecoding/go-library/pkg/strings/trim"
	"github.com/funtimecoding/go-library/pkg/system/secure_shell"
)

func (c *Client) Run(command string) *result.Result {
	s := secure_shell.Session(c.client)
	defer secure_shell.CloseSession(s)
	stdout, stderr := secure_shell.SessionBuffers(s)
	e := s.Run(command)
	r := &result.Result{
		OutputString: trim.NewLine(stdout.String()),
		ErrorString:  trim.NewLine(stderr.String()),
		Exit:         secure_shell.Exit(e),
		Error:        e,
	}

	if c.Panic {
		result.PanicOnExit(r)
		result.PanicOnError(r)
	}

	return r
}
