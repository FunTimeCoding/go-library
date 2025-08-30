package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readExtra(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.NotificationGroups() {
		fmt.Printf("NotificationGroup: %s\n", g.Format(f))
	}

	for _, t := range n.Tags() {
		fmt.Printf("Tag: %s\n", t.Format(f))
	}

	if false {
		// TODO: on load: panic: no value given for required property data_path
		for _, c := range n.ConfigContexts() {
			fmt.Printf("ConfigContext: %s\n", c.Format(f))
		}
	}

	// TODO: Must specify either local content or a data file
	//  How, what is this for?
	for _, t := range n.ConfigTemplates() {
		fmt.Printf("ConfigTemplate: %s\n", t.Format(f))
	}
}
