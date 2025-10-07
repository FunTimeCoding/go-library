package provisioner

type Provisioner struct {
	Type   string   `json:"type"`
	Inline []string `json:"inline"`
}
