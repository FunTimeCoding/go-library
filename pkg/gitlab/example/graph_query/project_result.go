package graph_query

type ProjectResult struct {
	Data struct {
		Project struct {
			ID string `json:"id"`
		} `json:"project"`
	} `json:"data"`
}
