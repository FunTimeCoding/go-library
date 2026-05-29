package argument

type Apply struct {
	Manifest  string `json:"manifest"`
	Namespace string `json:"namespace"`
	Override  bool   `json:"override"`
	DryRun    bool   `json:"dryRun"`
}
