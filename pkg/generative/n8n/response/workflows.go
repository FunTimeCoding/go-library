package response

type Workflows struct {
	Payload    []*Workflow `json:"data"`
	NextCursor any         `json:"nextCursor"`
}
