package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/viper"
)

func BindGeneric() Option {
	NotationArgument()
	AllArgument()
	VerboseArgument()
	argument.ParseBind()

	return Option{
		Notation: viper.GetBool(argument.Notation),
		All:      viper.GetBool(argument.All),
		Verbose:  viper.GetBool(argument.Verbose),
	}
}
