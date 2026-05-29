package argument

type PortForward struct {
	ResourceType string `json:"resourceType"`
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	LocalPort    int    `json:"localPort"`
	TargetPort   int    `json:"targetPort"`
}
