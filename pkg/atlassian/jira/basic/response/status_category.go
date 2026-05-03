package response

type StatusCategory struct {
	Self      string `json:"self"`
	Id        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}
