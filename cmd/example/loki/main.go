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
	if false {
		load()
	}

	if false {
		write()
	}

	if true {
		c := loki.NewEnvironment()
		end := time.Now()
		start := end.AddDate(0, 0, -7)

		if false {
			for _, l := range c.Labels(start, end) {
				fmt.Printf("Label: %s\n", l)
				fmt.Printf("  Values: %+v\n", c.LabelValues(start, end, l))
			}
		}

		fmt.Printf("Query: %s\n", c.Query("{namespace=\"i9n\"} |= ``"))

		if false {
			c.Statistic("")
			c.QueryRange("")
			c.Series("")
		}
	}
}

func write() {
	l := logrus.New()
	h, e := hook.NewHook(
		&hook.Config{
			URL: fmt.Sprintf(
				"https://%s:%s@%s/api/prom/push",
				environment.Get(constant.HostEnvironment, 1),
				environment.Get(constant.UserEnvironment, 1),
				environment.Get(constant.PasswordEnvironment, 1),
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
	// TODO: How to use https://github.com/grafana/loki/blob/main/pkg/loghttp/params.go
	//  Probably need a separate project to import this due to kubernetes client-go

	//lokiAddr := "http://localhost:3100"
	//
	//c, err := client.New(lokiAddr, client.Options{})
	//if err != nil {
	//	log.Fatalf("Failed to create Loki client: %v", err)
	//}
	//
	//query := `{job="my-app"}`
	//start := time.Now().Add(-1 * time.Hour)
	//end := time.Now()
	//
	//ctx := context.Background()
	//resp, err := c.QueryRange(
	//	ctx, query, loghttp.QueryRangeParams{
	//		Start:     start,
	//		End:       end,
	//		Step:      time.Minute, // Granularity
	//		Limit:     100,
	//		Direction: loghttp.FORWARD, // Or loghttp.BACKWARD
	//	},
	//)
	//if err != nil {
	//	log.Fatalf("Query failed: %v", err)
	//}
	//
	//if resp.Data.Result != nil {
	//	for _, stream := range resp.Data.Result.(logproto.Streams) {
	//		fmt.Printf("Stream: %+v\n", stream.Stream)
	//		for _, entry := range stream.Entries {
	//			fmt.Printf("[%s] %s\n", entry.Timestamp, entry.Line)
	//		}
	//	}
	//} else {
	//	fmt.Println("No logs found")
	//}
}
