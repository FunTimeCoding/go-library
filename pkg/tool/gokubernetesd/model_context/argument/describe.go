package argument

type Describe struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Unfiltered   bool   `json:"unfiltered"`
}
