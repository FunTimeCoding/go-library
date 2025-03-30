package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Favorite() {
	c := confluence.NewEnvironment()
	f := constant.DenseFormat

	for _, o := range c.FavoritesKaos() {
		fmt.Println(o.Format(f))
	}
}
