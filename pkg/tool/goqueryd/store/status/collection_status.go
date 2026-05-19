package status

type CollectionStatus struct {
	Name          string `json:"name"`
	Path          string `json:"path"`
	Pattern       string `json:"pattern"`
	DocumentCount int    `json:"document_count"`
}
