package merge_request

import "gitlab.com/gitlab-org/api/client-go"

type Request struct {
	Project    int
	Identifier int
	Title      string
	State      string
	Link       string

	Raw *gitlab.MergeRequest
}
