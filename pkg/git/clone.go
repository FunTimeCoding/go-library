package git

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/go-git/go-git/v5"
)

func Clone(locator string, path string) {
	_, e := git.PlainClone(
		path,
		false,
		&git.CloneOptions{URL: locator},
	)
	errors.FatalOnError(e)
}
