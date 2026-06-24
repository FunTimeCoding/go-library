package scan

type Configuration struct {
	Exclude  []string                       `yaml:"exclude"`
	Suppress map[string]map[string][]string `yaml:"suppress"`
}
