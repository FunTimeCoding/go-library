package argument

type CreatePipeline struct {
	Project   string            `json:"project"`
	Reference string            `json:"reference"`
	Variables map[string]string `json:"variables"`
}
