package common

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/spf13/pflag"
)

func Arguments() {
	pflag.String(argument.Host, "", "Host")
	pflag.String(argument.Token, "", "Token")
	pflag.String(argument.Owner, "", "Owner")
	pflag.String(argument.Repository, "", "Repository")
}
