package argument

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func ParseBind() {
	pflag.Parse()
	errors.PanicOnError(viper.BindPFlags(pflag.CommandLine))
}
