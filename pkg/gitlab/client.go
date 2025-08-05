package gitlab

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
)

type Client struct {
	context context.Context
	client  *gitlab.Client
	user    *gitlab.User

	verbose bool

	groups   []int
	projects []int

	projectCache map[int]*project.Project
}
