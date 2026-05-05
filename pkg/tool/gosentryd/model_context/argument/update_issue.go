package argument

type UpdateIssue struct {
	Identifier string `json:"identifier"`
	Status     string `json:"status"`
	AssignedTo string `json:"assignedTo"`
}
