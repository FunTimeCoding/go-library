package merge_request

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.BasicMergeRequest) []*Request {
	var result []*Request

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
