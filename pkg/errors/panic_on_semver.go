package errors

import (
	"golang.org/x/mod/semver"
	"log"
)

func PanicOnSemver(v string) {
	if !semver.IsValid(v) {
		log.Panicf("bad semver 2.0: %s", v)
	}
}
