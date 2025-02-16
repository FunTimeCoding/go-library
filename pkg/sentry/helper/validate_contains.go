package helper

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"log"
	"slices"
)

func ValidateContains(
	valid []string,
	element string,
) {
	if !slices.Contains(valid, element) {
		log.Panicf("%s not in %s", element, join.Comma(valid))
	}
}
