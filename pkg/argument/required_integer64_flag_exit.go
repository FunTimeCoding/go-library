package argument

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/viper"
)

func RequiredInteger64FlagExit(
	name string,
	exit int,
) int64 {
	if s := viper.GetInt64(name); s != 0 {
		return s
	}

	system.Exitf(exit, "flag empty: %s\n", name)

	return 0
}
