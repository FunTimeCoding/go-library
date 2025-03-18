package statistic

type Statistic struct {
	Severity SeverityCount
	State    StateCount
	Group    GroupCount
}

type SeverityCount struct {
	Critical int
	Warning  int
	Info     int
	None     int
	Unknown  int
}

type StateCount struct {
	Active     int
	Suppressed int
}

type GroupCount struct {
	All        int
	Kubernetes int
	Other      int
}
