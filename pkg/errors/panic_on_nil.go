package errors

import "log"

func PanicOnNil(
	a any,
	text string,
) {
	if a == nil {
		log.Panicf("nil: %s", text)
	}
}
