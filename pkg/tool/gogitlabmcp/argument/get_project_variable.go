package argument

type GetProjectVariable struct {
	Project string `json:"project"`
	Key     string `json:"key"`
}
