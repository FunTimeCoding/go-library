package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func main() {
	c := internal.Mattermost()

	t := timeLibrary.Midnight(time.Now())
	fmt.Printf(
		"Unresolved posts since %s\n",
		t.Format(timeLibrary.DateMinute),
	)
	posts := c.RecentPosts(c.DefaultChannel(), t.UnixMilli())

	for _, p := range posts {
		h := c.LoadThread(p)

		if h.Resolved {
			continue
		}

		fmt.Println(h.Root.Format())

		for _, p2 := range h.Posts {
			fmt.Printf("  %s\n", p2.Format())
		}
	}

	if len(posts) == 0 {
		fmt.Println("No unresolved posts")
	}
}
