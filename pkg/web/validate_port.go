package web

import "log"

func ValidatePort(p int) {
	if p < 1 || p > 65535 {
		log.Panicf("invalid port: %d", p)
	}
}
