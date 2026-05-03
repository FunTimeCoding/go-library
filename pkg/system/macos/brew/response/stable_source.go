package response

type StableSource struct {
	Locator  string  `json:"url"`
	Tag      *string `json:"tag"`
	Revision *string `json:"revision"`
	Using    *string `json:"using"`
	Checksum *string `json:"checksum"`
}
