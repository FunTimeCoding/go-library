package argument

type RolloutRestart struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
}
