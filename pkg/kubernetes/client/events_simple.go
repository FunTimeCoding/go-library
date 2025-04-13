package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/kubernetes/filter"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/event"
)

func (c *Client) EventsSimple(
	normal bool,
	warning bool,
) []*event.Event {
	f := filter.New()

	if warning && !normal {
		return c.Events(f, 0, constant.TypeWarning)
	}

	if normal && !warning {
		return c.Events(f, 0, constant.TypeNormal)
	}

	return c.Events(f, 0, "")
}
