package merge_request

import "gitlab.com/gitlab-org/api/client-go"

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
