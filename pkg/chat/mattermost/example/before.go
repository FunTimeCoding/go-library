package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	library "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func Before() {
	argument.ParseBind()
	channel := argument.RequiredPositional(0, "CHANNEL")
	m := mattermost.NewEnvironment(mattermost.WithVerbose(false))
	c := m.Channel(channel)
	f := constant.Format
	fmt.Printf("Channel: %s\n", c.Name)
	t := time.Now().Add(-30 * 24 * time.Hour)
	reference := m.PostBefore(c, t)

	if reference == nil {
		fmt.Printf("No post before %s\n", t.Format(library.DateMinute))

		return
	}

	fmt.Printf("Reference: %s\n", reference.Format(f))
	fmt.Printf(
		"Date: %s\n",
		time.UnixMilli(reference.Raw.CreateAt).Format(library.DateMinute),
	)

	posts := m.PostsBefore(c, t)
	fmt.Printf(
		"Posts before %s (%d)\n",
		t.Format(library.DateMinute),
		len(posts),
	)

	for _, p := range posts {
		fmt.Println(p.Format(f))
		fmt.Printf("  Time: %s\n", p.Create.Format(library.DateMinute))

		if false {
			m.DeletePost(p.Raw)
		}
	}
}
