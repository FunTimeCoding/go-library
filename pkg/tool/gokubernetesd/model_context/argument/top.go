package argument

type Top struct {
	ResourceType  string `json:"resourceType"`
	Namespace     string `json:"namespace"`
	AllNamespaces bool   `json:"allNamespaces"`
	Containers    bool   `json:"containers"`
}
