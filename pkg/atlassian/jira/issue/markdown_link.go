package issue

import "github.com/funtimecoding/go-library/pkg/text/markdown"

func (i *Issue) MarkdownLink() string {
	return markdown.Link(i.Key, i.Link)
}
