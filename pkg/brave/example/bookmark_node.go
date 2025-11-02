package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/brave"
	"github.com/funtimecoding/go-library/pkg/brave/bookmark/node"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BookmarkNode() {
	pflag.Int(argument.Depth, 0, "")
	argument.ParseBind()
	depth := viper.GetInt(argument.Depth)
	directory := argument.RequiredPositional(0, "DIRECTORY")
	b := brave.Bookmark(constant.DefaultProfile)
	f := constant.Format
	d := node.DirectoryByNameStrict(node.New(b.Root.Bar), directory)
	fmt.Printf("Root: %s\n", d.Format(f))

	if depth > 0 {
		node.WalkLevels(
			d,
			depth,
			func(o *node.Node) {
				fmt.Println(o.Format(f))
			},
		)
	} else {
		for _, l := range node.Collect(d) {
			fmt.Println(l.Format(f))
		}
	}

	for _, g := range node.GroupByDirectory(d) {
		fmt.Printf(
			"Group %s (%d)\n",
			g.Directory.Name,
			len(g.Links),
		)
	}
}
