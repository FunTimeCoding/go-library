package main

import (
	"fmt"
	"github.com/akkuman/logrus-loki-hook"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	if true {
		load()
	}

	if false {
		write()
	}
}

func write() {
	l := logrus.New()
	h, e := hook.NewHook(
		&hook.Config{
			URL: fmt.Sprintf(
				"https://%s:%s@%s/api/prom/push",
				environment.Required(constant.HostEnvironment),
				environment.Required(constant.UserEnvironment),
				environment.Required(constant.PasswordEnvironment),
			),
			LevelName: "severity",
			Labels:    map[string]string{"application": "test"},
		},
	)

	if e != nil {
		l.Error(e)
	} else {
		l.AddHook(h)
	}

	l.Info("test message")
}

func load() {
	c := loki.NewEnvironment()
	end := time.Now()

	if false {
		start := end.AddDate(0, 0, -7)

		for _, l := range c.Labels(start, end) {
			fmt.Printf("Label: %s\n", l)
			fmt.Printf("  Values: %+v\n", c.LabelValues(start, end, l))
		}
	}

	start := end.Add(-1 * time.Hour)
	fmt.Printf(
		"QueryRange: %+v\n",
		c.QueryRange("{namespace=\"i9n\"} |= ``", start, end),
	)

	if false {
		fmt.Printf(
			"Query: %+v\n",
			c.Query("{namespace=\"i9n\"} |= ``"),
		)
		c.Statistic("")
		c.Series("")
	}
}
