package repository

type Repository struct {
	MonitorIdentifier string
	Name              string
	Path              string
	IsClean           bool
	Status            string

	concern []string
}
