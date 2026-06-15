package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func searchResultCount(
	count int,
	collection string,
) gomponents.Node {
	text := fmt.Sprintf("%d results", count)

	if collection != "" {
		text = fmt.Sprintf("%d results in %s", count, collection)
	}

	return html.P(gomponents.Text(text))
}
