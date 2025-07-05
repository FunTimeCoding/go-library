package git

import "github.com/funtimecoding/go-library/pkg/git/constant"

func Push() {
	Run(constant.Push, constant.OriginRemote, constant.Tags)
}
