package brave

import (
	"github.com/funtimecoding/go-library/pkg/brave/bookmark/file"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Bookmark(profile string) *file.Bookmark {
	return file.Parse(
		system.ReadFile(
			ProfileByNameStrict(profile).Path,
			constant.BookmarksFile,
		),
	)
}
