package example

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func Create() {
	c := alertmanager.NewEnvironment()
	c.Create(
		constant.HighMemoryUsage,
		"localhost:9090",
		"High memory usage detected",
		"Memory usage exceeded 90% for 5 minutes.",
		"memory_usage",
	)
}
