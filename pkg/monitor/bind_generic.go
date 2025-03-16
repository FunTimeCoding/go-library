package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/viper"
)

type Parameter struct {
	Notation bool
}

func BindGeneric() Parameter {
	NotationArgument()
	argument.ParseAndBind()

	return Parameter{Notation: viper.GetBool(argument.Notation)}
}
