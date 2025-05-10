package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go"
)

type Client struct {
	context  context.Context
	client   *gitlab.Client
	user     *gitlab.User
	projects []int
}
