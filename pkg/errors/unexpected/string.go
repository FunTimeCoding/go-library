package unexpected

import "log"

func String(v string) {
	log.Panicf("unexpected: %s", v)
}
