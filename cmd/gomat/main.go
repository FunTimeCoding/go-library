package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/thread"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/text/template"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"time"
)

func main() {
	c := common.Mattermost()

	t := timeLibrary.Midnight(time.Now())
	fmt.Printf(
		"Unresolved threads since %s\n",
		t.Format(timeLibrary.DateMinute),
	)
	var relevant []*thread.Thread

	for _, p := range c.RecentPosts(c.DefaultChannel(), t.UnixMilli()) {
		h := c.LoadThread(p)

		if h.Resolved {
			continue
		}

		relevant = append(relevant, h)
	}

	f := constant.Format
	o := ollama.NewEnvironment()
	te := `
Summarize this tech support chat thread.
Identify:
- the MOST RECENT unclear question of the ongoing investigation (paraphrase it)
- the next step SOMEONE needs to do to get closer to a solution, concluded from the conversation people are having (paraphrase it)
Respond in the expected summary format, just the two lines and without extra styling.

## Expected summary format

Latest unclear question: [most recent problem]
Next step: [the concluded next step, or "probably resolved", or "unclear"]

## Example responses from previous threads

### Example 1 (concluded next step)

Latest unclear question: Why is the server not responding?
Next step: jdoe checks the servers logs

### Example 2 (probably resolved)

Latest unclear question: It looks like it works now
Next step: probably resolved

### Example 3 (unclear next step)

Latest unclear question: I don't understand why the database is not connecting
Next step: unclear

## Thread log

{{.}}
`
	tem := template.New("simple", te)

	for _, h := range relevant {
		formatted := h.Format(f)
		fmt.Println(formatted)
		r := o.GenerateSimple(template.Execute(tem, formatted))
		fmt.Printf("%s\n", console.Magenta("%s", r.Text))

		if false {
			r.Print()
		}
	}

	if len(relevant) == 0 {
		fmt.Println("No unresolved threads")
	}
}
