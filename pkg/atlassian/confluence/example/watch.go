package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Watch() {
	c := confluence.NewEnvironment()
	f := constant.Dense

	if true {
		fmt.Println("Watch")

		for _, p := range c.Watched() {
			fmt.Println(p.Format(f))
		}

		fmt.Println("Favorite")

		for _, p := range c.Favorites() {
			fmt.Println(p.Format(f))
		}
	}

	if false {
		fmt.Println("FavoritesKaos")

		for _, o := range c.FavoritesKaos() {
			fmt.Println(o.Format(f))
		}
	}
}
