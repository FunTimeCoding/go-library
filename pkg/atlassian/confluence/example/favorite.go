package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Favorite() {
	c := confluence.NewEnvironment()
	f := constant.DenseFormat

	if true {
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
