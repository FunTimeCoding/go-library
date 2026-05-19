package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestSendDirectMessage(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "receiver")
	a.MustCallTool(
		constant.Send,
		map[string]any{
			constant.To:   b.Name(),
			constant.Body: "heads up: running lint",
		},
	)
	r := b.CheckLive()
	assert.Count(t, 1, r.Messages)
	assert.String(t, a.Name(), r.Messages[0].From)
	assert.String(t, "heads up: running lint", r.Messages[0].Body)
}

func TestSendBroadcast(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "sender")
	b.Announce(b.Name(), "listener")
	a.MustCallTool(
		constant.Send,
		map[string]any{
			constant.Body: "deploying service",
		},
	)
	r := b.CheckLive()
	assert.Count(t, 1, r.Messages)
	assert.String(t, "deploying service", r.Messages[0].Body)
}
