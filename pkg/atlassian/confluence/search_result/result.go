package search_result

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"

type Result struct {
	Name string

	Raw *response.Result
}
