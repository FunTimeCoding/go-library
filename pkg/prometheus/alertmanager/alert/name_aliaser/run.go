package name_aliaser

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

func (a *Aliaser) Run(v []*alert.Alert) []*alert.Alert {
	for _, e := range v {
		e.Name = a.Alias(e.Name)
	}

	return v
}
