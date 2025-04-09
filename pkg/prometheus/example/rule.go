package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"slices"
)

func Rule() {
	c := prometheus.NewEnvironment()
	f := option.ExtendedColor.Copy()
	var severities []string

	for _, r := range c.Rules().Alert() {
		if r.RawAlert != nil {
			fmt.Printf("Alert: %s\n", r.Format(f))

			for k, v := range r.RawAlert.Labels {
				if k == constant.SeverityLabel {
					asString := string(v)

					if !slices.Contains(severities, asString) {
						severities = append(severities, asString)
					}
				}
			}
		} else if r.RawRecord != nil {
			fmt.Printf("Record: %s\n", r.Format(f))

			for k, v := range r.RawRecord.Labels {
				if k == constant.SeverityLabel {
					asString := string(v)

					if !slices.Contains(severities, asString) {
						severities = append(severities, asString)
					}
				}
			}
		} else {
			fmt.Printf("Rule: %+v\n", r)
		}
	}

	fmt.Printf("Severities: %+v\n", severities)
}
