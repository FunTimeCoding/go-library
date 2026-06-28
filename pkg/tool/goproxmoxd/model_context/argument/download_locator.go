package argument

type DownloadLocator struct {
	Node     string `json:"node"`
	Storage  string `json:"storage"`
	Content  string `json:"content"`
	Filename string `json:"filename"`
	Locator  string `json:"locator"`
}
