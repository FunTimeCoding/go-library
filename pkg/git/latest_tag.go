package git

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/go-git/go-git/v5/plumbing"
	"time"
)

func LatestTag(path string) string {
	r := Open(path)
	t := TagsIterator(r)
	var result string
	var latest time.Time

	errors.PanicOnError(
		t.ForEach(
			func(t *plumbing.Reference) error {
				commit := CommitObject(r, t.Hash())

				if commit.Committer.When.After(latest) {
					latest = commit.Committer.When
					result = t.Name().Short()
				}

				return nil
			},
		),
	)

	return result
}
