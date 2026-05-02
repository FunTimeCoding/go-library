package argument

type MergeRequestDiscussions struct {
	Project      string `json:"project"`
	MergeRequest int64  `json:"merge_request"`
}
