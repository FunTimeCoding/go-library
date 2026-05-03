package response

type Requirement struct {
	Name     string   `json:"name"`
	Cask     any      `json:"cask"`
	Download any      `json:"download"`
	Version  *string  `json:"version"`
	Contexts []string `json:"contexts"`
	Specs    []string `json:"specs"`
}
