package common

import "github.com/funtimecoding/go-library/pkg/argument"

func ValidateArguments() {
	argument.Required(argument.Host)
	argument.Required(argument.Token)
	argument.Required(argument.Owner)
	argument.Required(argument.Repository)
}
