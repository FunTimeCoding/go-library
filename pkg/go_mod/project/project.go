package project

type Project struct {
	MonitorIdentifier string
	Path              string
	Version           string

	concern []string

	runtimeVersion string
}
