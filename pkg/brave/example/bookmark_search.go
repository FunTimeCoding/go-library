package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/brave"
	"github.com/funtimecoding/go-library/pkg/brave/bookmark"
	"github.com/funtimecoding/go-library/pkg/brave/bookmark/node"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BookmarkSearch() {
	pflag.String(argument.Type, bookmark.DirectoryType, "bookmark type")
	argument.ParseBind()
	search := argument.RequiredPositional(0, "NAME")
	bookmarkType := viper.GetString(argument.Type)
	b := brave.Bookmark(constant.DefaultProfile)
	f := constant.Format
	root := node.New(b.Root.Bar)
	results := node.FindAllByNameAndType(root, search, bookmarkType)
	node.SetParents(root)

	if len(results) == 0 {
		fmt.Println("No results")

		return
	}

	for _, n := range results {
		fmt.Println(n.Format(f))
		fmt.Printf("  %s\n", n.FormatPath())
	}
}
