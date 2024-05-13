package systemd

const (
	Command   = "systemctl"
	ListUnits = "list-units"
	Status    = "status"

	NoLegend = "--no-legend"

	All   = "--all"   // Units in memory, including dead and empty
	Full  = "--full"  // Do not shorten unit names
	Plain = "--plain" // Dependencies as a list instead of tree

	State    = "--state"
	NotFound = "not-found"

	Type    = "--type"
	Service = "service"

	Output   = "--output"
	Notation = "json"
)
