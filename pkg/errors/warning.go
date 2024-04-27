package errors

import "log"

func Warning(s string) {
	log.Printf("warning: %s\n", s)
}
