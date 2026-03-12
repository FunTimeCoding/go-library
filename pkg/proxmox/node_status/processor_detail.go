package node_status

type ProcessorDetail struct {
	Cores     int    `json:"cores"`
	Count     int    `json:"cpus"`
	Flag      string `json:"flags"`
	Hardware  string `json:"hvm"`
	Megahertz string `json:"mhz"`
	Model     string `json:"model"`
	Sockets   int    `json:"sockets"`
	UserHertz int    `json:"user_hz"`
}
