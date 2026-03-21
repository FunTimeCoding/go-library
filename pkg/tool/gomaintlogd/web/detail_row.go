package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func detailRow(e *store.Entry) g.Node {
	target := fmt.Sprintf("#detail-%d", e.ID)

	return h.Tr(
		h.ID(fmt.Sprintf("detail-%d", e.ID)),
		h.Class("detail-row"),
		h.Td(
			g.Attr("colspan", "6"),
			h.Div(
				h.Class("detail-content"),
				h.Div(
					h.Class("grid"),
					h.Div(
						h.Strong(g.Text("Timestamp: ")),
						g.Text(formatTime(e.Timestamp)),
					),
					h.Div(
						h.Strong(g.Text("Action: ")),
						g.Text(e.Action),
					),
					h.Div(
						h.Strong(g.Text("User: ")),
						g.Text(e.User),
					),
				),
				h.Div(
					h.Class("grid"),
					h.Div(
						h.Strong(g.Text("System: ")),
						g.Text(e.System),
					),
					h.Div(
						h.Strong(g.Text("Service: ")),
						g.Text(e.Service),
					),
					h.Div(),
				),
				h.Div(
					h.Strong(g.Text("Description:")),
					h.Pre(g.Text(e.Description)),
				),
				h.Div(
					h.Class("detail-actions"),
					h.Button(
						h.Class("outline"),
						hx.Get(fmt.Sprintf("/entry/edit?id=%d", e.ID)),
						hx.Target(target),
						hx.Swap("outerHTML"),
						g.Text("Edit"),
					),
					h.Button(
						h.Class("outline contrast"),
						hx.Post(fmt.Sprintf("/entry/delete?id=%d", e.ID)),
						hx.Confirm("Delete this entry?"),
						g.Attr("hx-on::after-request",
							fmt.Sprintf(
								"document.getElementById('row-%d')?.remove();document.getElementById('detail-%d')?.remove()",
								e.ID, e.ID,
							),
						),
						g.Text("Delete"),
					),
					h.Button(
						h.Type("button"),
						h.Class("outline secondary"),
						g.Attr("onclick",
							fmt.Sprintf(
								"var r=document.getElementById('detail-%d');r.style.display='none';r.innerHTML='';r.className=''",
								e.ID,
							),
						),
						g.Text("Close"),
					),
				),
			),
		),
	)
}
