package helper

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"log"
	"slices"
)

func ValidateContains(
	valid []string,
	s string,
) {
	if !slices.Contains(valid, s) {
		log.Panicf("%s not in %s", s, join.Comma(valid))
	}
}
