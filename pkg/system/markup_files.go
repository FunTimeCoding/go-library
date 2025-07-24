package system

import "github.com/funtimecoding/go-library/pkg/constant"

func MarkupFiles(root string) []string {
	return FilesByExtension(
		root,
		constant.MarkupExtension,
		constant.ShortMarkupExtension,
	)
}
