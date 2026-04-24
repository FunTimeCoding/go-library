package argument

type GetPipelineJob struct {
	Project string `json:"project"`
	Job     int64  `json:"job"`
}
