package environment

import (
	"log"
	"os"
)

func Get(s string) string {
	result := os.Getenv(s)

	if result == "" {
		log.Panicf("%s not set", s)
	}

	return result
}
