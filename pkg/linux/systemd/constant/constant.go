package constant

const (
	Command   = "systemctl"
	ListUnits = "list-units"
	Status    = "status"
	Show      = "show"

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

	DateTime = "Mon 2006-01-02 15:04:05 MST"
)

const ActiveState = "active"

const RunningSubState = "running"

const (
	ActiveEnterTimestamp   = "ActiveEnterTimestamp"
	ExecMainStartTimestamp = "ExecMainStartTimestamp"
)
