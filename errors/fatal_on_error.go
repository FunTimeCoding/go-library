package errors

import "log"

func FatalOnError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
