package argument

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/viper"
)

func RequiredIntegerFlag(
	name string,
	exitCode int,
) int64 {
	if s := viper.GetInt64(name); s != 0 {
		return s
	}

	system.Exitf(exitCode, "flag empty: %s\n", name)

	return 0
}
