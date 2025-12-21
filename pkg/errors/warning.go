package errors

import (
	"fmt"
	"log"
)

func Warning(s string, a ...any) {
	if len(a) > 0 {
		s = fmt.Sprintf(s, a...)
	}

	log.Printf("warning: %s\n", s)
}
