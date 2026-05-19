package result

type Document struct {
	VirtualPath string `json:"virtual_path"`
	FilePath    string `json:"file_path"`
	Collection  string `json:"collection"`
	Path        string `json:"path"`
	Title       string `json:"title"`
	Hash        string `json:"hash"`
	Context     string `json:"context,omitempty"`
	Body        string `json:"body"`
}
