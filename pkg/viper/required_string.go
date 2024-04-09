package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func RequiredString(name string) {
	if viper.GetString(name) == "" {
		fmt.Printf("required argument empty: %s\n", name)

		os.Exit(1)
	}
}
