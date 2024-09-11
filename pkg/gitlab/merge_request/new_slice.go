package merge_request

import "github.com/xanzy/go-gitlab"

func NewSlice(v []*gitlab.MergeRequest) []*Request {
	var result []*Request

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
