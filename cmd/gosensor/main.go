package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func main() {
	p := monitor.BindGeneric()

	if p.Notation {
		r := report.New()
		r.AddItem(
			fmt.Sprintf("%s-%d", constant.ExamplePrefix, 1),
			constant.ErrorType,
			strings.Alfa,
			"https://example.org/1",
		)
		r.Print()
	}
}
