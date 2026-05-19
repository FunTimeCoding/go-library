package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sessionRow(
	s *session.Session,
	activeNames map[string]bool,
	aliases map[string]string,
) gomponents.Node {
	name := []gomponents.Node{
		statusDot(s.LastSeen),
		gomponents.Text(s.Name),
	}

	if activeNames[s.Name] {
		name = append(
			name,
			gomponents.Text(" "),
			html.Small(gomponents.Text("(active)")),
		)
	}

	return html.Tr(
		html.Td(gomponents.Group(name)),
		html.Td(gomponents.Text(s.Topic)),
		html.Td(html.Small(gomponents.Text(aliases[s.Name]))),
		html.Td(gomponents.Textf("%d", s.TurnCount)),
		html.Td(html.Small(gomponents.Text(relativeTime(s.LastSeen)))),
		html.Td(html.Em(html.Small(gomponents.Text(s.FirstMessage)))),
	)
}
