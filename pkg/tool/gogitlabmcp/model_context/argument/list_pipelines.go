package argument

type ListPipelines struct {
	Project   string `json:"project"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
}
