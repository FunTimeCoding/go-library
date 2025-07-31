package collector

import (
	"github.com/funtimecoding/go-library/pkg/monitor/item/source"
	"time"
)

type Collector struct {
	Name     string
	Prefix   string
	Plural   string
	Interval time.Duration
	Source   *source.Source
}
