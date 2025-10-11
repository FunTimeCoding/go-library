package search_result

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"

func New(r *response.Result) *Result {
	return &Result{
		Name: r.Title,
		Raw:  r,
	}
}
