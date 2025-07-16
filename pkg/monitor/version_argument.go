package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/pflag"
)

func VersionArgument() {
	pflag.Bool(
		argument.Version,
		false,
		"Show version information and exit",
	)
}
