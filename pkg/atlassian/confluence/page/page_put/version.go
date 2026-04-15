package page_put

type Version struct {
	Number  int    `json:"number"`
	Message string `json:"message,omitempty"`
}
