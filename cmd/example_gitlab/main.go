package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/web"
)

func main() {
	c := gitlab.New(
		web.TrimScheme(environment.Get(gitlab.Host, 1)),
		environment.Get(gitlab.Token, 1),
		[]int{},
	)

	f := format.New().Color().Extended()
	runners := c.Runners(true)
	fmt.Printf("Runners (%d):\n", len(runners))

	for _, element := range runners {
		fmt.Printf("  %s\n", element.Format(f))
	}
}
