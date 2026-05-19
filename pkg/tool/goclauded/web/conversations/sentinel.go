package conversations

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sentinel(offset int) gomponents.Node {
	return html.Div(
		gomponents.Attr(
			"data-load-more",
			fmt.Sprintf("/conversations/sidebar?offset=%d", offset),
		),
		html.Style("height:1px"),
	)
}
