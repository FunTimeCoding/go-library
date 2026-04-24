package argument

type GetCommit struct {
	Project string `json:"project"`
	Sha     string `json:"sha"`
}
