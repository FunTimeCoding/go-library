package loki

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/message"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
)

func printLog(
	v []*message.Message,
	f *option.Format,
) {
	for _, m := range v {
		s := status.New(f).String(
			m.Time.Format(timeLibrary.DateMinute),
		).String(
			formatContent(m, f),
		).String(
			m.Stream,
		)
		fmt.Println(s.Format())
	}
}
