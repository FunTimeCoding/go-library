package errors

import "log"

func UnexpectedCount(expected int, a []any) {
	if len(a) != expected {
		log.Panicf("unexpected count: %d", expected)
	}
}
