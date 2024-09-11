package merge_request

import "github.com/xanzy/go-gitlab"

type Request struct {
	Project    int
	Identifier int
	Title      string
	State      string
	Link       string

	Raw *gitlab.MergeRequest
}
