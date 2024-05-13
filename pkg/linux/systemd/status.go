package systemd

import (
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"regexp"
	"strings"
)

func (c *Client) Status(name string) *ssh.Result {
	result := c.ssh.Run(join.Space(Command, Status, Output, Notation, name))

	return result
}

func GrabUntilFirstBlankLine(s string) string {
	return strings.Split(s, "\n\n")[0]
}

type StatusResult struct {
	Active      string
	Loaded      string
	MainProcess string
	Memory      string
	Processor   string
}

func ParseStatus(s string) *StatusResult {
	result := &StatusResult{}

	if m := regexp.MustCompile(
		`Loaded: (.+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Loaded = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`Active: (.+) since (.+);`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Active = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`Main PID: (\d+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.MainProcess = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`Memory: (.+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Memory = strings.TrimSpace(m[1])
	}

	if m := regexp.MustCompile(
		`CPU: (.+)`,
	).FindStringSubmatch(s); len(m) > 1 {
		result.Processor = strings.TrimSpace(m[1])
	}

	return result
}
