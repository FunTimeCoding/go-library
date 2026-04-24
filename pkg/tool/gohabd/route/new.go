package route

import "github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"

func New(c *habitica.Client) *Router {
	return &Router{habitica: c}
}
