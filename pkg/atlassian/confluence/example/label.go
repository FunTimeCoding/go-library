package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Label() {
	c := confluence.NewEnvironment()
	f := constant.DenseFormat

	for _, o := range c.Labeled() {
		fmt.Println(o.Format(f))
	}
}
