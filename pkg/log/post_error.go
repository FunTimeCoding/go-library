package log

import "log"

func PostError(
	locator string,
	text string,
) {
	log.Printf("post error (%s): %s\n", locator, text)
}
