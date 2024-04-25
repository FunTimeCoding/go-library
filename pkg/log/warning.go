package log

import "log"

func Warning(s string) {
	log.Printf("warning: %s\n", s)
}
