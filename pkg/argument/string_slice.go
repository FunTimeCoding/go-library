package argument

import (
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/spf13/viper"
)

func StringSlice(name string) []string {
	if s := viper.GetString(name); s != "" {
		return split.Comma(s)
	}

	return []string{}
}
