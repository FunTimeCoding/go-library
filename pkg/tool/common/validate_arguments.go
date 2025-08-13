package common

import "github.com/funtimecoding/go-library/pkg/argument"

func ValidateArguments() {
	argument.RequiredStringFlag(argument.Host)
	argument.RequiredStringFlag(argument.Token)
	argument.RequiredStringFlag(argument.Owner)
	argument.RequiredStringFlag(argument.Repository)
}
