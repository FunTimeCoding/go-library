package merge_request

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.BasicMergeRequest) []*Request {
	var result []*Request

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
