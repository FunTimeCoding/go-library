package repository

type Repository struct {
	MonitorIdentifier string

	Path    string
	IsClean bool
	Status  string
}
