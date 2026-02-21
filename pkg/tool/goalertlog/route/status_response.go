package route

type StatusResponse struct {
	TotalRecords int    `json:"totalRecords"`
	LastPoll     string `json:"lastPoll,omitempty"`
}
