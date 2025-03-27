package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Favourite() {
	c := confluence.NewEnvironment()
	f := constant.DenseFormat

	for _, o := range c.Favourites() {
		fmt.Println(o.Format(f))
	}
}
