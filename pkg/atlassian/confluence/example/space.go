package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Space() {
	c := confluence.NewEnvironment()
	f := constant.DefaultFormat

	if true {
		for _, s := range c.Spaces() {
			fmt.Println(s.Format(f))
		}
	}

	if false {
		fmt.Println("SpacesKaos")

		// Requires a key, cannot list all spaces
		for _, s := range c.SpacesKaos([]string{""}) {
			fmt.Printf("  Space: %s\n", s.Name)
		}
	}

	if false {
		fmt.Println("SpacesVirtomize")

		for _, s := range c.SpacesVirtomize() {
			fmt.Printf("  Space: %s\n", s.Name)
		}
	}

	if false {
		fmt.Println("SpacesTreminio")

		for _, s := range c.SpacesTreminio() {
			fmt.Printf("  Space: %s\n", s.Name)
		}
	}

	if false {
		fmt.Println("SpacesTreminioV2")

		for _, s := range c.SpacesTreminioV2() {
			fmt.Printf("  Space: %s\n", s.Name)
		}
	}
}
