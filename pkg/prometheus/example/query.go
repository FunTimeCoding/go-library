package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"time"
)

func Query() {
	c := prometheus.NewEnvironment()
	now := time.Now()
	load1 := c.QueryFloat(constant.Load1, now)
	load5 := c.QueryFloat(constant.Load5, now)
	load15 := c.QueryFloat(constant.Load15, now)
	fmt.Printf("Load: %.1f %.1f %.1f\n", load1, load5, load15)
}
