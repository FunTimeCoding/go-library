package tag

type ListResponse struct {
	Count   int        `json:"count"`
	Results []response `json:"results"`
}
