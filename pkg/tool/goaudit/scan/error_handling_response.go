package scan

type errorHandlingResponse struct {
	Reference   string                          `yaml:"$ref"`
	Description string                          `yaml:"description"`
	Content     map[string]errorHandlingContent `yaml:"content"`
}
