package mock_client

import "gitlab.com/gitlab-org/api/client-go/v2"

type Client struct {
	files   map[string]*gitlab.File
	commits []*Commit
}
