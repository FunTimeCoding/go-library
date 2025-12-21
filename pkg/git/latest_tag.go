package git

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/go-git/go-git/v5/plumbing"
	"time"
)

func LatestTag(path string) string {
	r := Open(path)
	var result string
	var latest time.Time

	errors.PanicOnError(
		TagsIterator(r).ForEach(
			func(t *plumbing.Reference) error {
				// TODO: Under what conditions does CommitObject fail?
				c, _ := r.CommitObject(t.Hash())

				if c != nil && c.Committer.When.After(latest) {
					latest = c.Committer.When
					result = t.Name().Short()
				}

				return nil
			},
		),
	)

	return result
}
