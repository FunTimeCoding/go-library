package merge_request

import "github.com/xanzy/go-gitlab"

func New(v *gitlab.MergeRequest) *Request {
	return &Request{
		Project:    v.ProjectID,
		Identifier: v.IID,
		Title:      v.Title,
		State:      v.State,
		Link:       v.WebURL,
		Raw:        v,
	}
}
