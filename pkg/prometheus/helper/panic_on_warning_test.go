package helper

import (
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"testing"
)

func TestPanicOnWarning(t *testing.T) {
	PanicOnWarning(prometheus.Warnings{})
}
