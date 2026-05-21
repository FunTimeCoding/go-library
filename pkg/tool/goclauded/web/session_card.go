package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func sessionCard(
	s *session.Session,
	lines int,
) gomponents.Node {
	var details []gomponents.Node

	if s.Alias != nil {
		details = append(
			details,
			html.P(html.Strong(gomponents.Text(*s.Alias))),
		)
	}

	if s.Topic != "" {
		details = append(
			details,
			html.P(gomponents.Text(s.Topic)),
		)
	}

	if s.Description != "" {
		details = append(
			details,
			html.P(
				html.Small(gomponents.Text(s.Description)),
			),
		)
	}

	if s.Files != "" {
		details = append(
			details,
			html.P(
				html.Small(gomponents.Text(shortenPaths(s.Files))),
			),
		)
	}

	var metadata []string

	if s.TurnCount > 0 {
		metadata = append(metadata, fmt.Sprintf("%d turns", s.TurnCount))
	}

	if lines > 0 {
		metadata = append(metadata, fmt.Sprintf("%d lines", lines))
	}

	metadata = append(metadata, relativeTime(s.LastSeen))
	details = append(
		details,
		html.P(
			html.Small(gomponents.Text(strings.Join(metadata, " · "))),
		),
	)

	return html.A(
		gomponents.Attr(
			"href",
			fmt.Sprintf("/sessions/%s", s.Identifier),
		),
		html.Class("session-card"),
		html.H4(
			statusDot(s.LastSeen),
			gomponents.Text(s.Name),
		),
		gomponents.Group(details),
	)
}
