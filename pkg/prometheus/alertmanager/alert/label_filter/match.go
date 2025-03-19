package label_filter

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"slices"
)

func (f *Filter) match(a *alert.Alert) bool {
	for k, v := range a.Labels {
		if slices.Contains(f.keepLabel, k) {
			fmt.Printf("keepLabel: %s\n", k)
			return true
		}

		if slices.Contains(f.dropLabel, k) {
			fmt.Printf("dropLabel: %s\n", k)
			return false
		}

		for _, l := range f.keepValue {
			if l.Match(k, v) {
				fmt.Printf("keepValue: %s=%s\n", k, v)
				return true
			}
		}

		for _, l := range f.dropValue {
			if l.Match(k, v) {
				fmt.Printf("dropValue: %s=%s\n", k, v)
				return false
			}
		}
	}

	return f.KeepByDefault
}
