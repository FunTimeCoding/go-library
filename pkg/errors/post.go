package errors

import "log"

func Post(
	locator string,
	text string,
) {
	log.Printf("post error (%s): %s\n", locator, text)
}
