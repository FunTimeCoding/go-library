package argument

type CancelPipeline struct {
	Project  string `json:"project"`
	Pipeline int64  `json:"pipeline"`
}
