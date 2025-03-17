package alert

import "github.com/funtimecoding/go-library/pkg/markdown"

func (a *Alert) MarkdownLink() string {
	return markdown.Link(a.Name, a.Link)
}
