package response

type Workflows struct {
	Data       []*Workflow `json:"data"`
	NextCursor any         `json:"nextCursor"`
}
