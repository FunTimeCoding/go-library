package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Parameter struct {
	Notation bool
}

func BindFlag() Parameter {
	pflag.BoolP(
		argument.Notation,
		"n",
		false,
		"JSON output",
	)
	argument.ParseAndBind()

	return Parameter{Notation: viper.GetBool(argument.Notation)}
}
