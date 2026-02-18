package formula

type Formula struct {
	MonitorIdentifier string
	Name              string
	InstalledVersions []string
	CurrentVersion    string
	Link              string
	Pinned            bool
}
