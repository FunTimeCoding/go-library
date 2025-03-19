package statistic

type Statistic struct {
	Total    int
	Relevant int

	Severity SeverityCount
	State    StateCount
	Group    GroupCount
}
