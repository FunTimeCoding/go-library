package response

type Result struct {
	Stream Stream     `json:"stream"`
	Values [][]string `json:"values"`
}
