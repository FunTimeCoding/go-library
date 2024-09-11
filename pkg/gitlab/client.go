package gitlab

import "github.com/xanzy/go-gitlab"

type Client struct {
	client   *gitlab.Client
	user     *gitlab.User
	projects []int
}
