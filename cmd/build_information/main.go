package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/system"
	timeHelper "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func main() {
	r := git.Open(system.WorkingDirectory())
	h := git.Head(r).Hash()
	fmt.Printf("Short hash: %s\n", h.String()[:8])

	if false {
		c := git.CommitObject(r, h)
		fmt.Printf("Long hash: %s\n", h)
		fmt.Printf("Author: %s\n", c.Author.Name)
		fmt.Printf("Date: %s\n", c.Author.When)
		fmt.Printf("Message: %s", c.Message)
	}

	fmt.Printf("Latest: %s\n", git.LatestTag(system.WorkingDirectory()))

	fmt.Printf("Date: %s\n", time.Now().Format(timeHelper.DateMinute))
}
