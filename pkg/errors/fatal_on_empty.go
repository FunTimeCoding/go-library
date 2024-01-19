package errors

import "log"

func FatalOnEmpty(
	s string,
	text string,
) {
	if s == "" {
		log.Fatalf("empty: %s", text)
	}
}
