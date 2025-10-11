package argument

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/viper"
)

func RequiredExit(
	name string,
	exit int,
) string {
	if s := viper.GetString(name); s != "" {
		return s
	}

	system.Exitf(exit, "flag empty: %s\n", name)

	return ""
}
