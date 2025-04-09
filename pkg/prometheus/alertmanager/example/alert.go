package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
)

func Alert() {
	alerts, _ := internal.Alertmanager().Alerts(advanced_option.New())

	for _, a := range alerts {
		fmt.Printf("Alert: %+v\n", a.Format(option.ExtendedColor))
	}
}
