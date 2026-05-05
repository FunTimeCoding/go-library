package argument

type CreateBranch struct {
	Project   string `json:"project"`
	Branch    string `json:"branch"`
	Reference string `json:"reference"`
}
