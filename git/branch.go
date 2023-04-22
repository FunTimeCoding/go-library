package git

import (
	"github.com/funtimecoding/go-library/errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"strings"
)

func Branch(path string) string {
	fileSystem := osfs.New(path)
	dot, dotFail := fileSystem.Chroot(".git")
	errors.PanicOnError(dotFail)

	repository, openFail := git.Open(
		filesystem.NewStorage(dot, cache.NewObjectLRUDefault()),
		fileSystem,
	)
	errors.PanicOnError(openFail)

	branch, branchFail := repository.Head()
	errors.PanicOnError(branchFail)

	return strings.TrimPrefix(branch.Name().String(), "refs/heads/")
}
