package response

type UninstallScript struct {
	Executable string   `json:"executable"`
	Sudo       bool     `json:"sudo"`
	Input      []string `json:"input,omitempty"`
	Arguments  []string `json:"args,omitempty"`
}
