package argument

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/viper"
)

func RequiredIntegerFlagExit(
	name string,
	exit int,
) int {
	if s := viper.GetInt(name); s != 0 {
		return s
	}

	system.Exitf(exit, "flag empty: %s\n", name)

	return 0
}
