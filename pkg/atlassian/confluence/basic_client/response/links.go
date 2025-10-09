package response

type Links struct {
	Base     string `json:"base"`
	WebUI    string `json:"webui"`
	EditUI   string `json:"editui"`
	EditUIV2 string `json:"edituiv2"`
	TinyUI   string `json:"tinyui"`
	Context  string `json:"context"`
	Next     string `json:"next"`
	Self     string `json:"self"`
}
