package convert

type SlimAlertResult struct {
	Alerts   []*SlimAlert       `json:"alerts"`
	Total    int                `json:"total"`
	Severity SlimSeverityCount  `json:"severity"`
	State    SlimStateCount     `json:"state"`
}
