package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
)

func Alert() {
	for _, a := range alertmanager.NewEnvironment().Alerts() {
		fmt.Printf("Alert: %+v\n", a.Format(option.ExtendedColor))
	}
}
