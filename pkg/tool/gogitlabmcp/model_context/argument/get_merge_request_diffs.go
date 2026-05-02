package argument

type GetMergeRequestDiffs struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
}
