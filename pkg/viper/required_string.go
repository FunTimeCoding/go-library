package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func RequiredString(name string) string {
	if s := viper.GetString(name); s != "" {
		return s
	}

	fmt.Printf("required argument empty: %s\n", name)
	os.Exit(1)

	return ""
}
