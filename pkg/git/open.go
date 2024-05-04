package git

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
)

func Open(path string) *git.Repository {
	f := osfs.New(path)
	dot, dotFail := f.Chroot(Directory)
	errors.PanicOnError(dotFail)

	r, openFail := git.Open(
		filesystem.NewStorage(dot, cache.NewObjectLRUDefault()),
		f,
	)
	errors.PanicOnError(openFail)

	return r
}
