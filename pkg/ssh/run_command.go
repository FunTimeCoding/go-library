package ssh

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ssh/command"
	"github.com/funtimecoding/go-library/pkg/ssh/result"
	"github.com/funtimecoding/go-library/pkg/strings/trim"
	"github.com/funtimecoding/go-library/pkg/system/secure_shell"
	"golang.org/x/crypto/ssh"
)

func (c *Client) RunCommand(o *command.Command) *result.Result {
	s := secure_shell.Session(c.client)
	defer secure_shell.CloseSession(s)
	stdout, stderr := secure_shell.SessionBuffers(s)

	if o.RequestTeletype {
		errors.PanicOnError(
			s.RequestPty(
				"xterm",
				25,
				80,
				ssh.TerminalModes{
					//ssh.ECHO:          0,     // disable echo
					ssh.TTY_OP_ISPEED: 14400, // input speed kilo-baud
					ssh.TTY_OP_OSPEED: 14400, // output speed in kilo-baud
				},
			),
		)
	}

	setEnvironment(s, o)
	var text string

	if prefix := environmentPrefix(o); prefix != "" {
		text = fmt.Sprintf("%s %s", prefix, o.Command)
	} else {
		text = o.Command
	}

	e := s.Run(text)

	return &result.Result{
		OutputString: trim.NewLine(stdout.String()),
		ErrorString:  trim.NewLine(stderr.String()),
		ExitCode:     secure_shell.ExitCode(e),
		Error:        e,
	}
}
