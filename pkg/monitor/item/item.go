package item

import (
	"github.com/funtimecoding/go-library/pkg/monitor/collector"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"time"
)

type Item struct {
	collector *collector.Collector

	Identifier string
	Severity   constant.Severity
	Detail     string
	Status     constant.Status
	Assignee   string
	Link       string
	Create     *time.Time
	Update     *time.Time
	Score      float64
}
