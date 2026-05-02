package argument

type GetFileContents struct {
	Project   string `json:"project"`
	Path      string `json:"path"`
	Reference string `json:"reference"`
}
