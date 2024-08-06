package argument

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/viper"
)

func RequiredStringFlag(
	name string,
	exitCode int,
) string {
	if s := viper.GetString(name); s != "" {
		return s
	}

	system.Exitf(exitCode, "flag empty: %s\n", name)

	return ""
}
