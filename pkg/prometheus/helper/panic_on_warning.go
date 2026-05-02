package helper

import (
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
)

func PanicOnWarning(w v1.Warnings) {
	if len(w) > 0 {
		log.Panicf("warnings: %v", w)
	}
}
