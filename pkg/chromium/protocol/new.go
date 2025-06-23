package protocol

import "github.com/funtimecoding/go-library/pkg/chromium"

func New(tab string) *Protocol {
	c := chromium.NewEnvironment()
	t := c.TabByHost(tab)

	return &Protocol{client: c, context: c.TargetContext(t.Id)}
}
