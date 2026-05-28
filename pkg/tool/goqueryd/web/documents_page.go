package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) documentsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.PathValue("name")
	documents, e := s.service.ListDocuments(name)

	if e != nil {
		s.view.RenderPage(
			w,
			name,
			constant.CollectionsPath,
			html.P(
				gomponents.Textf(
					"Failed to load documents: %s",
					e.Error(),
				),
			),
		)

		return
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(name)),
		html.P(gomponents.Textf("%d documents", len(documents))),
	)

	if len(documents) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No documents.")),
		)
	} else {
		var rows []gomponents.Node

		for _, d := range documents {
			rows = append(
				rows,
				html.Tr(
					html.Td(
						html.A(
							gomponents.Attr(
								"href",
								fmt.Sprintf(
									"/documents/%s",
									documentPath(d.VirtualPath),
								),
							),
							gomponents.Text(d.Title),
						),
					),
					html.Td(
						html.Small(
							gomponents.Text(d.VirtualPath),
						),
					),
				),
			)
		}

		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Title")),
						html.Th(gomponents.Text("Path")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	s.view.RenderPage(w, name, constant.CollectionsPath, content...)
}
