package mock_client

import "gitlab.com/gitlab-org/api/client-go/v2"

type Commit struct {
	Branch  string
	Message string
	Actions []*gitlab.CommitActionOptions
}
