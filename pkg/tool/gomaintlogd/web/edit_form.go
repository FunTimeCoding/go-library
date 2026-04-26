package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func editForm(e *entry.Entry) g.Node {
	target := fmt.Sprintf("#detail-%d", e.ID)

	return h.Tr(
		h.ID(fmt.Sprintf("detail-%d", e.ID)),
		h.Class("detail-row"),
		h.Td(
			g.Attr("colspan", "6"),
			h.Form(
				h.Class("edit-form"),
				hx.Post(fmt.Sprintf("/entry/edit?id=%d", e.ID)),
				hx.Target(target),
				hx.Swap("outerHTML"),
				h.Div(
					h.Class("grid"),
					h.Label(
						g.Text("Action"),
						h.Input(
							h.Type("text"),
							h.Name(constant.Action),
							h.Value(e.Action),
							h.Required(),
						),
					),
					h.Label(
						g.Text("User"),
						h.Input(
							h.Type("text"),
							h.Name(constant.User),
							h.Value(e.User),
							h.Required(),
						),
					),
					h.Label(
						g.Text("Timestamp"),
						h.Input(
							h.Type("datetime-local"),
							h.Name(constant.Timestamp),
							h.Value(
								e.Timestamp.Format("2006-01-02T15:04"),
							),
						),
					),
				),
				h.Div(
					h.Class("grid"),
					h.Label(
						g.Text("System"),
						h.Input(
							h.Type("text"),
							h.Name(constant.System),
							h.Value(e.System),
						),
					),
					h.Label(
						g.Text("Service"),
						h.Input(
							h.Type("text"),
							h.Name(constant.Service),
							h.Value(e.Service),
						),
					),
				),
				h.Label(
					g.Text("Description"),
					h.Textarea(
						h.Name(constant.Description),
						h.Rows("4"),
						g.Text(e.Description),
					),
				),
				h.Div(
					h.Class("detail-actions"),
					h.Button(h.Type("submit"), g.Text("Save")),
					h.Button(
						h.Type("button"),
						h.Class("outline secondary"),
						hx.Get(fmt.Sprintf("/entry/detail?id=%d", e.ID)),
						hx.Target(target),
						hx.Swap("outerHTML"),
						g.Text("Cancel"),
					),
				),
			),
		),
	)
}
