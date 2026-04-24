package argument

type GetPipeline struct {
	Project  string `json:"project"`
	Pipeline int64  `json:"pipeline"`
}
