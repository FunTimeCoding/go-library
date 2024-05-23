package ssh

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ssh/command"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func environmentPrefix(o *command.Command) string {
	var parts []string

	for k, v := range o.Environment {
		parts = append(parts, fmt.Sprintf("%s=%s;", k, v))
	}

	return join.Space(parts...)
}
