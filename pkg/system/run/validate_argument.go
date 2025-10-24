package run

import (
	"log"
	"regexp"
)

func ValidateArgument(s string) {
	if s == "" || regexp.MustCompile(`[=/\\\n\r]`).MatchString(s) {
		log.Panicf("invalid: %q", s)
	}
}
