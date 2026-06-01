package convert

type SlimSilence struct {
	Identifier string `json:"identifier"`
	State      string `json:"state"`
	Rule       string `json:"rule"`
	Match      string `json:"match"`
	Author     string `json:"author"`
	Comment    string `json:"comment"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Remaining  string `json:"remaining,omitempty"`
	Link       string `json:"link,omitempty"`
}
