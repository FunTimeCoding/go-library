package errors

import "log"

func LogOnError(e error) {
	if e != nil {
		log.Println(e)
	}
}
