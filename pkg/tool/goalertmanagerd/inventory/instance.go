package inventory

type Instance struct {
	Name         string     `yaml:"name"`
	Alertmanager Connection `yaml:"alertmanager"`
	Prometheus   Prometheus `yaml:"prometheus"`
}
