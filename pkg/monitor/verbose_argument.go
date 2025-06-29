package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/pflag"
)

func VerboseArgument() {
	pflag.Bool(argument.Verbose, false, "Verbose output")
}
