package errors

import "log"

func NotFound(name string) {
	log.Panicf("not found: %s", name)
}
