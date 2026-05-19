package result

type ContextEntry struct {
	Collection  string `json:"collection"`
	PathPrefix  string `json:"path_prefix"`
	Description string `json:"description"`
}
