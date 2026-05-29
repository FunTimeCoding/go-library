package argument

type Get struct {
	ResourceType  string `json:"resourceType"`
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	AllNamespaces bool   `json:"allNamespaces"`
	LabelSelector string `json:"labelSelector"`
	FieldSelector string `json:"fieldSelector"`
	Unfiltered    bool   `json:"unfiltered"`
}
