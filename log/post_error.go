package log

import "log"

func PostError(
	locator string,
	text string,
) {
	log.Printf("Post error (%s): %s\n", locator, text)
}
