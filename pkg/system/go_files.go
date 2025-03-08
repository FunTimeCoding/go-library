package system

import "github.com/funtimecoding/go-library/pkg/constant"

func GoFiles(root string) []string {
	return FilesByExtension(root, constant.GoExtension)
}
