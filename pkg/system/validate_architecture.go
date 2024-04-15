package system

import (
	"log"
	"slices"
)

func ValidateArchitecture(s string) {
	if !slices.Contains(Architectures, s) {
		log.Panicf("unexpected architecture: %s", s)
	}
}
