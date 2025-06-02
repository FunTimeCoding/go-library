package common

import "github.com/funtimecoding/go-library/pkg/argument"

func ValidateArguments() {
	argument.ParseBind()
	argument.RequiredStringFlag(argument.Host, 1)
	argument.RequiredStringFlag(argument.Token, 1)
	argument.RequiredStringFlag(argument.Owner, 1)
	argument.RequiredStringFlag(argument.Repository, 1)
}
