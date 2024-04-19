package errors

import "log"

func UnexpectedCount(
	expected int,
	actual int,
) {
	if actual != expected {
		log.Panicf("unexpected count: %d", expected)
	}
}
