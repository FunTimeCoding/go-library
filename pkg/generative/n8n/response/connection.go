package response

type Connection struct {
	Main [][]struct {
		Node  string `json:"node"`
		Type  string `json:"type"`
		Index int    `json:"index"`
	} `json:"main"`
}
