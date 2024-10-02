package commit

import "github.com/funtimecoding/go-library/pkg/console/description"

func (c *Commit) Meta() *description.Description {
	return description.New("Commit", "Commit")
}
