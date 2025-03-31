package space

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"

func NewSlice(v []*response.Space) []*Space {
	var result []*Space

	for _, c := range v {
		result = append(result, New(c))
	}

	return result
}
