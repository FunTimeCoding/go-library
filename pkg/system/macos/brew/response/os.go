package response

type OS struct {
	Cellar string `json:"cellar"`
	Url    string `json:"url"`
	Sha256 string `json:"sha256"`
}
