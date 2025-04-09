package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Label() {
	c := confluence.NewEnvironment()
	f := constant.Dense

	if true {
		for _, l := range c.Labels() {
			fmt.Printf("Label: %+v\n", l)
		}

		c.SetLabels([]string{"favourite"})

		for _, o := range c.Labeled() {
			fmt.Println(o.Format(f))
		}
	}

	if false {
		fmt.Println("LabeledKaos")

		for _, o := range c.LabeledKaos() {
			fmt.Println(o.Format(f))
		}
	}
}
