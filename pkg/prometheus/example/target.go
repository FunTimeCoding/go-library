package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/prometheus/common/model"
)

func Target() {
	c := prometheus.NewEnvironment()
	r := c.Targets()

	for _, t := range r.Active {
		fmt.Printf("Active: %+v\n", t.ScrapePool)
	}

	for _, t := range r.Dropped {
		address, _ := t.DiscoveredLabels[model.AddressLabel]
		fmt.Printf("Dropped: %+v\n", address)
	}
}
