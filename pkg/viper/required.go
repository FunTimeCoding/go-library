package viper

import (
	"github.com/spf13/viper"
	"log"
)

func Required(name string) {
	if viper.GetString(name) == "" {
		log.Panicf("argument empty: %s", name)
	}
}
