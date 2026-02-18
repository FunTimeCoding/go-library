package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/pflag"
)

func CopyableArgument() {
	pflag.Bool(
		argument.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
}
