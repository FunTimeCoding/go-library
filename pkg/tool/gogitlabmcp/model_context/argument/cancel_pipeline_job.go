package argument

type CancelPipelineJob struct {
	Project string `json:"project"`
	Job     int64  `json:"job"`
}
