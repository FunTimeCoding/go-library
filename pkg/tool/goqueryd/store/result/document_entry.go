package result

type DocumentEntry struct {
	VirtualPath string `json:"virtual_path"`
	FilePath    string `json:"file_path"`
	Title       string `json:"title"`
}
