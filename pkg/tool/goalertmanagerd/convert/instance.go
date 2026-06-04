package convert

import "github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/inventory"

func Instance(
	i *inventory.Instance,
	active bool,
) *SlimInstance {
	return &SlimInstance{
		Name:             i.Name,
		AlertmanagerHost: i.Alertmanager.Host,
		PrometheusHost:   i.Prometheus.Host,
		Active:           active,
	}
}
