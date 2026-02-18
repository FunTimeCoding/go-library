package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	monitor.CopyableArgument()
	pflag.Bool(argument.Watched, false, "Watched")
	pflag.Bool(argument.Favorites, false, "Favorites")
	argument.ParseBind()
	c := confluence.NewEnvironment()
	f := constant.Format.Copy()

	if viper.GetBool(argument.Copyable) {
		f.Tag(tag.Copyable)
	}

	if viper.GetBool(argument.Watched) || viper.GetBool(argument.Favorites) {
		fmt.Println("Watch")

		for _, p := range c.Watched() {
			fmt.Println(p.Format(f))
			p.PrintConsole()
		}

		fmt.Println("Favorite")

		for _, p := range c.Favorites() {
			fmt.Println(p.Format(f))
		}

		return
	}

	for _, p := range c.Pages() {
		fmt.Println(p.Format(f))
		p.PrintConsole()
	}

	// TODO: Synchronize runbooks
}
