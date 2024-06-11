package argument

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func RequiredStringFlag(
	name string,
	exitCode int,
) string {
	if s := viper.GetString(name); s != "" {
		return s
	}

	fmt.Printf("flag empty: %s\n", name)
	os.Exit(exitCode)

	return ""
}
