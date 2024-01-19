package errors

import "log"

func FatalOnZero(
	i int,
	text string,
) {
	if i == 0 {
		log.Fatalf("zero: %s", text)
	}
}
