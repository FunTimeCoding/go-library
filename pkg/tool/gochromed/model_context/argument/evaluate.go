package argument

type Evaluate struct {
	Expression string `json:"expression"`
	TabID      string `json:"tab_id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
}
