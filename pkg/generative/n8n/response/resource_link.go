package response

type ResourceLink struct {
	Rl    bool   `json:"__rl"`
	Mode  string `json:"mode"`
	Value string `json:"value"`
}
