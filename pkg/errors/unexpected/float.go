package unexpected

import "log"

func Float(v float64) {
	log.Panicf("unexpected: %f", v)
}
