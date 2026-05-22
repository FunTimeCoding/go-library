package status

type CollectionStatus struct {
	Name          string `json:"name"`
	Path          string `json:"path"`
	Pattern       string `json:"pattern"`
	Type          string `json:"type"`
	DocumentCount int    `json:"document_count"`
}
