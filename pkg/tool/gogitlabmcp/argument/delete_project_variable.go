package argument

type DeleteProjectVariable struct {
	Project string `json:"project"`
	Key     string `json:"key"`
}
