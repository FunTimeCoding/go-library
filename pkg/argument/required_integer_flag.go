package argument

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func RequiredIntegerFlag(
	name string,
	exitCode int,
) int64 {
	if s := viper.GetInt64(name); s != 0 {
		return s
	}

	fmt.Printf("flag empty: %s\n", name)
	os.Exit(exitCode)

	return 0
}
