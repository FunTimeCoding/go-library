package argument

type GetRepositoryTree struct {
	Project   string `json:"project"`
	Path      string `json:"path"`
	Reference string `json:"reference"`
	Recursive bool   `json:"recursive"`
}
