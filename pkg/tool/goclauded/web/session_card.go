package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func sessionCard(s *session.Session) gomponents.Node {
	var details []gomponents.Node

	if s.Topic != "" {
		details = append(
			details,
			html.P(gomponents.Text(s.Topic)),
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

	metadata = append(metadata, relativeTime(s.LastSeen))
	details = append(
		details,
		html.P(
			html.Small(gomponents.Text(strings.Join(metadata, " · "))),
		),
	)

	if s.FirstMessage != "" {
		details = append(
			details,
			html.P(
				html.Em(
					html.Small(
						gomponents.Textf("\"%s\"", s.FirstMessage),
					),
				),
			),
		)
	}

	return html.Div(
		html.Class("session-card"),
		html.H4(
			statusDot(s.LastSeen),
			gomponents.Text(s.Name),
		),
		gomponents.Group(details),
	)
}
