package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
)

func eventsLink(o *store.QueryOption, page int) string {
	link := fmt.Sprintf("/events?page=%d", page)

	if o.Tool != "" {
		link = fmt.Sprintf("%s&tool=%s", link, o.Tool)
	}

	if o.Surface != "" {
		link = fmt.Sprintf("%s&surface=%s", link, o.Surface)
	}

	if o.Actor != "" {
		link = fmt.Sprintf("%s&actor=%s", link, o.Actor)
	}

	return link
}
