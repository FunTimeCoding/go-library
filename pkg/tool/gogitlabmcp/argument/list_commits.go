package argument

type ListCommits struct {
	Project   string `json:"project"`
	Reference string `json:"reference"`
}
