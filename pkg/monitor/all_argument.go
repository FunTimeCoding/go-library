package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/pflag"
)

func AllArgument() {
	pflag.Bool(argument.All, false, "Include filtered in output")
}
