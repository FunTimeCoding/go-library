package issue

type CurrentStatus struct {
	Status         string   `json:"status"`
	StatusCategory string   `json:"statusCategory"`
	StatusDate     JiraDate `json:"statusDate"`
}
