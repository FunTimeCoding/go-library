package merge_request

import "github.com/xanzy/go-gitlab"

func New(
	input *gitlab.MergeRequest,
) *Request {
	return &Request{
		Project:    input.ProjectID,
		Identifier: input.IID,
		Title:      input.Title,
		State:      input.State,
		Link:       input.WebURL,
		Raw:        input,
	}
}
