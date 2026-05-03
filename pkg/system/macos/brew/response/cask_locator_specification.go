package response

type CaskLocatorSpecification struct {
	Verified string `json:"verified,omitempty"`
	Branch   string `json:"branch,omitempty"`
	OnlyPath string `json:"only_path,omitempty"`
}
