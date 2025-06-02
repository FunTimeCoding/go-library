package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.Bool(argument.Watched, false, "Favorites")
	argument.ParseBind()
	c := confluence.NewEnvironment()
	f := constant.Format

	if viper.GetBool(argument.Watched) {
		for _, p := range c.Favorites() {
			fmt.Println(p.Format(f))
		}

		return
	}

	for _, p := range c.Pages() {
		fmt.Println(p.Format(f))
	}
}
