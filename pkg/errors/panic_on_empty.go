package errors

import "log"

func PanicOnEmpty(
	s string,
	text string,
) {
	if s == "" {
		log.Panicf("empty: %s", text)
	}
}
