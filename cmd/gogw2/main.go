package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.StringP(argument.Tag, "g", "", "Guild tag")
	argument.ParseAndBind()
	c := gw2.New(environment.Get(gw2.TokenEnvironment, 1))
	tag := viper.GetString(argument.Tag)
	account := c.Account()

	if len(account.GuildLeader) == 0 {
		fmt.Println("No guilds with leader permissions")

		return
	}

	if tag == "" {
		fmt.Printf("No guild tag provided, not printing members\n")
	}

	var tagFound bool

	for _, element := range account.GuildLeader {
		g := c.Guild(element)
		fmt.Printf("Guild: %s\n", g.Name)
		fmt.Printf("  Tag: %s\n", g.Tag)

		if tag != "" && tag == g.Tag {
			tagFound = true
			members := c.Members(element)
			fmt.Printf("  Members: %d\n", len(members))

			for _, member := range members {
				fmt.Printf("    Member: %+v\n", member.Name)
			}
		}
	}

	if tag != "" && !tagFound {
		fmt.Printf("Guild tag %s not found\n", tag)
	}
}
