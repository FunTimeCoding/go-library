package response

type StatusCategory struct {
	Self       string `json:"self"`
	Identifier int    `json:"id"`
	Key        string `json:"key"`
	ColorName  string `json:"colorName"`
	Name       string `json:"name"`
}
