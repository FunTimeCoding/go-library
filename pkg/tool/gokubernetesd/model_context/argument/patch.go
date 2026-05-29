package argument

type Patch struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Patch        string `json:"patch"`
	Type         string `json:"type"`
	DryRun       bool   `json:"dryRun"`
}
