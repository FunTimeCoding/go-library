package scan

type errorHandlingSpec struct {
	Paths      map[string]map[string]errorHandlingOperation `yaml:"paths"`
	Components errorHandlingComponents                      `yaml:"components"`
}
