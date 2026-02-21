package assistant

import (
	"github.com/funtimecoding/go-library/pkg/assistant/constant"
	"github.com/funtimecoding/go-library/pkg/assistant/message"
	"github.com/getsentry/sentry-go"
)

func (c *Client) Start() {
	c.connection.Connect()
	h := c.subscriber
	wrap := h

	if c.hub != nil {
		wrap = func(m *message.Message) {
			defer func() {
				if r := recover(); r != nil {
					c.hub.WithScope(
						func(s *sentry.Scope) {
							s.SetContext(
								constant.EventContext,
								map[string]any{
									constant.TypeKey: m.Event.Type,
									constant.RawKey:  string(m.Event.Raw),
								},
							)
							c.hub.RecoverWithContext(c.context, r)
						},
					)
				}
			}()
			h(m)
		}
	}

	c.connection.Subscribe("", wrap)
	go c.connection.Read()
	go c.poll()
}
