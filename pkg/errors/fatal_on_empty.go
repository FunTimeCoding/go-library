package errors

import "log"

func FatalOnEmpty(
	s string,
	text string,
) {
	if s == "" {
		log.Fatalf("Empty: %s", text)
	}
}
