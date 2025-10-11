package response

type Result struct {
	Stream struct {
		App       string `json:"app"`
		Container string `json:"container"`
		Filename  string `json:"filename"`
		Job       string `json:"job"`
		Namespace string `json:"namespace"`
		NodeName  string `json:"node_name"`
		Pod       string `json:"pod"`
		Stream    string `json:"stream"`
	} `json:"stream"`
	Values [][]string `json:"values"`
}
