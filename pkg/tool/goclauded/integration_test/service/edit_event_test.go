//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestEditEventSummaryCascade(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Service.Summarize("session-1", "Cedar", "original summary"),
	)
	events := s.Store.EventsSince(
		time.Time{},
		time.Time{},
		constant.Summarize,
		10,
		0,
	)
	assert.Count(t, 1, events)
	_, e := s.Service.EditEvent(events[0].Identifier, "revised summary")
	assert.FatalOnError(t, e)
	body := s.Store.SummaryBySession("session-1")
	assert.String(t, "revised summary", body)
}

func TestEditEventCompletionCascade(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "fixing auth", "")
	topic := s.Store.CompleteTask(r.Callsign)
	s.Service.Complete("session-1", r.Callsign, topic, "original message")
	events := s.Store.EventsSince(
		time.Time{},
		time.Time{},
		constant.Complete,
		10,
		0,
	)
	assert.Count(t, 1, events)
	_, e := s.Service.EditEvent(events[0].Identifier, "corrected message")
	assert.FatalOnError(t, e)
	completions := s.Store.CompletionsBySession("session-1")
	assert.Count(t, 1, completions)
	assert.String(t, "corrected message", completions[0].Summary)
}
