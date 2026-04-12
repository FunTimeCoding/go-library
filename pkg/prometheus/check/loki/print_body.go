package loki

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/message"
	telemetry "github.com/funtimecoding/go-library/pkg/web/telemetry/constant"
)

func printBody(messages []*message.Message) {
	for _, v := range messages {
		body := v.Value(telemetry.Body)

		if body != "" {
			fmt.Println(body)
		}
	}
}
