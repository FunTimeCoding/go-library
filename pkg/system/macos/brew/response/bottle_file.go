package response

type BottleFile struct {
	Cellar  string `json:"cellar"`
	Locator string `json:"url"`
	Sha256  string `json:"sha256"`
}
