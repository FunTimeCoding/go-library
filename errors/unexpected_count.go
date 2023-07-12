package errors

import "log"

func UnexpectedCount(c int) {
	log.Panicf("unexpected count: %d", c)
}
