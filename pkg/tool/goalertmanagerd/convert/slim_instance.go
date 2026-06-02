package convert

type SlimInstance struct {
	Name             string `json:"name"`
	AlertmanagerHost string `json:"alertmanager_host"`
	PrometheusHost   string `json:"prometheus_host"`
	Active           bool   `json:"active"`
}
