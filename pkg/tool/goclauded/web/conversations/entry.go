package conversations

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/session"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func entry(
	s *session.Session,
	aliases map[string]string,
) gomponents.Node {
	name := s.Slug

	if a, okay := aliases[s.Identifier]; okay {
		name = a
	}

	if name == "" {
		name = "unnamed"
	}

	return html.Div(
		html.ID(fmt.Sprintf("entry-%s", s.Identifier)),
		html.Class("sidebar-entry"),
		html.Div(
			html.Class("entry-name"),
			gomponents.Attr(
				"hx-get",
				fmt.Sprintf("/conversations/%s", s.Identifier),
			),
			gomponents.Attr("hx-target", "#panel"),
			html.Span(gomponents.Text(name)),
			html.Span(
				html.Class("rename-icon"),
				gomponents.Attr(
					"hx-get",
					fmt.Sprintf("/conversations/%s/rename", s.Identifier),
				),
				gomponents.Attr(
					"hx-target",
					fmt.Sprintf("#entry-%s", s.Identifier),
				),
				gomponents.Attr("hx-swap", "innerHTML"),
				gomponents.Attr("onclick", "event.stopPropagation()"),
				gomponents.Text("✎"),
			),
		),
		html.Small(gomponents.Text(relativeTimestamp(s.Timestamp))),
	)
}
