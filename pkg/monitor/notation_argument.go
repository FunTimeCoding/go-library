package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/pflag"
)

func NotationArgument() {
	pflag.BoolP(
		argument.Notation,
		"n",
		false,
		"JSON output",
	)
}
