package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func main() {
	c := confluence.NewEnvironment()
	f := constant.Format

	for _, p := range c.Pages() {
		fmt.Println(p.Format(f))
	}
}
