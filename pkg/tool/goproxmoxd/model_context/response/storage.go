package response

type Storage struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Enabled bool   `json:"enabled"`
	Shared  bool   `json:"shared"`
	Active  bool   `json:"active"`
	Avail   uint64 `json:"avail"`
	Used    uint64 `json:"used"`
	Total   uint64 `json:"total"`
}
