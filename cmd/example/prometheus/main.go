package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"slices"
)

func main() {
	if false {
		prometheusExample()
	}

	if true {
		alertmanagerExample()
	}
}

func prometheusExample() {
	c := prometheus.NewEnvironment()
	var severities []string

	for _, r := range c.Rules().Get() {
		fmt.Printf("Rule: %+v\n", r)

		if r.RawAlert != nil {
			fmt.Printf("  Alert: %+v\n", r.RawAlert)

			for k, v := range r.RawAlert.Labels {
				if k == constant.SeverityField {
					asString := string(v)

					if !slices.Contains(severities, asString) {
						severities = append(severities, asString)
					}
				}
			}
		}

		if r.RawRecord != nil {
			fmt.Printf("  Record: %+v\n", r.RawRecord)

			for k, v := range r.RawRecord.Labels {
				if k == constant.SeverityField {
					asString := string(v)

					if !slices.Contains(severities, asString) {
						severities = append(severities, asString)
					}
				}
			}
		}
	}

	fmt.Printf("Severities: %+v\n", severities)
}

func alertmanagerExample() {
	c := alertmanager.NewEnvironment()

	for _, a := range c.Alerts() {
		fmt.Printf("Alert: %+v\n", a.Format(format.ExtendedColor))
	}
}
