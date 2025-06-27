package github

import "github.com/funtimecoding/go-library/pkg/github/repository"

func (c *Client) ActionRepository() []*repository.Repository {
	var result []*repository.Repository

	for _, o := range c.SearchCode(
		"actions/checkout user:%s in:file language:yaml",
		c.User().Name,
	) {
		result = append(result, repository.New(o.Raw.Repository))
	}

	return result
}
