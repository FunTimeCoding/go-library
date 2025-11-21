package project

type Project struct {
	MonitorIdentifier string
	Name              string
	Path              string
	Version           string

	concern        []string
	runtimeVersion string
}
