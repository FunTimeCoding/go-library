package status

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/text/markdown"
)

func (s *Status) DetailLink(
	url string,
	label string,
	prefix string,
) *Status {
	if url == "" {
		return s
	}

	if label == "" {
		label = "Link"
	}

	if s.format.HasTag(tag.Copyable) {
		return s.TagLine(tag.Copyable, "  %s", prefixed(prefix, url))
	}

	if s.format.HasTag(tag.Markdown) {
		return s.TagLine(
			tag.Markdown,
			"  %s",
			prefixed(prefix, markdown.Link(label, url)),
		)
	}

	s.bubbles = append(s.bubbles, console.Link(url, label, true))

	return s
}
