package gitlab

import "gitlab.com/gitlab-org/api/client-go"

type Client struct {
	client   *gitlab.Client
	user     *gitlab.User
	projects []int
}
