package argument

type GetCommitDiff struct {
	Project string `json:"project"`
	Sha     string `json:"sha"`
}
