package unexpected

import "log"

func Integer(v int) {
	log.Panicf("unexpected: %d", v)
}
