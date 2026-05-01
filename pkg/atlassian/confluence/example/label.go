package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Label() {
	c := confluence.NewEnvironment()
	f := constant.Dense

	for _, l := range c.MustLabels() {
		fmt.Printf("Label: %+v\n", l)
	}

	c.SetLabels([]string{"favourite"})

	for _, o := range c.MustLabeled() {
		fmt.Println(o.Format(f))
	}
}
