package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/viper"
)

func BindGeneric() Option {
	NotationArgument()
	argument.ParseAndBind()

	return Option{Notation: viper.GetBool(argument.Notation)}
}
