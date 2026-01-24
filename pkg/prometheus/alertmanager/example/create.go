package example

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func Create() {
	c := common.Alertmanager()
	c.Create(
		constant.HighMemoryUsage,
		"localhost:9090",
		"High memory usage detected",
		"Memory usage exceeded 90% for 5 minutes.",
		"memory_usage",
	)
}
