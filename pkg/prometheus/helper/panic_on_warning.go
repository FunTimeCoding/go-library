package helper

import (
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"log"
)

func PanicOnWarning(w prometheus.Warnings) {
	if len(w) > 0 {
		log.Panicf("warnings: %v", w)
	}
}
