package git

import "github.com/funtimecoding/go-library/pkg/git/constant"

func Fetch() {
	Run(constant.Fetch, constant.Prune, constant.PruneTags)
}
