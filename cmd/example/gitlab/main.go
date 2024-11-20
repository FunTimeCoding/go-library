package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/gitlab"
)

func main() {
	f := format.New().Color().Extended()
	runners := gitlab.NewEnvironment().Runners(true)
	fmt.Printf("Runners (%d):\n", len(runners))

	for _, element := range runners {
		fmt.Printf("  %s\n", element.Format(f))
	}
}
