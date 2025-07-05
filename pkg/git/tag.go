package git

import "github.com/funtimecoding/go-library/pkg/git/constant"

func Tag(name string) {
	Run(constant.Tag, name)
}
