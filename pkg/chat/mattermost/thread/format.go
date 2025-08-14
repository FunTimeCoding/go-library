package thread

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/text/multi_line"
)

func (t *Thread) Format(f *option.Format) string {
	l := multi_line.New()
	l.Add(t.Root.Format(f))

	for _, p := range t.Posts {
		l.Add(strings.PrefixMultiline(p.Format(f), "> "))
	}

	return l.Render()
}
