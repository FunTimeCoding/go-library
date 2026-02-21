package poller

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
	"testing"
	"time"
)

var testLogger = logger.New(context.Background())

func TestPoller(t *testing.T) {
	assert.NotNil(t, New(nil, nil, testLogger, 1*time.Minute, 30*24*time.Hour))
}

func TestPollSavesNewAlert(t *testing.T) {
	path := t.TempDir() + "/test.db"
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "abc123",
			Name:        "HighMemory",
			Severity:    "critical",
			Summary:     "Memory above 90%",
		},
	)
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	p.Poll()
	records := s.ByName("HighMemory")
	assert.Count(t, 1, records)
	assert.String(t, "abc123", records[0].Fingerprint)
	assert.String(t, "critical", records[0].Severity)
	assert.True(t, records[0].End == nil)
}

func TestPollResolvesRemovedAlert(t *testing.T) {
	path := t.TempDir() + "/test.db"
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "abc123",
			Name:        "HighMemory",
			Severity:    "critical",
			Summary:     "Memory above 90%",
		},
	)
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	p.Poll()
	c.Remove("abc123")
	p.Poll()
	records := s.ByName("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollIgnoresDuplicateFiring(t *testing.T) {
	path := t.TempDir() + "/test.db"
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "abc123",
			Name:        "HighMemory",
			Severity:    "critical",
		},
	)
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	p.Poll()
	p.Poll()
	p.Poll()
	assert.Count(t, 1, s.ByName("HighMemory"))
}

func TestRecoverStaleAdoptsFireing(t *testing.T) {
	path := t.TempDir() + "/test.db"
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "abc123",
			Name:        "HighMemory",
			Severity:    "critical",
		},
	)
	old := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	old.Poll()
	fresh := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	fresh.RecoverStale()
	assert.Count(t, 1, s.Unresolved())
	fresh.Poll()
	assert.Count(t, 1, s.ByName("HighMemory"))
	assert.True(t, s.ByName("HighMemory")[0].End == nil)
}

func TestRecoverStaleResolvesGone(t *testing.T) {
	path := t.TempDir() + "/test.db"
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "abc123",
			Name:        "HighMemory",
			Severity:    "critical",
		},
	)
	old := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	old.Poll()
	c.Remove("abc123")
	fresh := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	fresh.RecoverStale()
	assert.Count(t, 0, s.Unresolved())
	records := s.ByName("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollPrunesOldRecords(t *testing.T) {
	path := t.TempDir() + "/test.db"
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		&alert.Alert{
			Fingerprint: "abc123",
			Name:        "HighMemory",
			Severity:    "critical",
		},
	)
	p := New(c, s, testLogger, 1*time.Minute, 0)
	p.Poll()
	c.Remove("abc123")
	p.Poll()
	assert.Count(t, 0, s.ByName("HighMemory"))
}

func TestPollTracksLastPollTime(t *testing.T) {
	path := t.TempDir() + "/test.db"
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour)
	assert.True(t, p.LastPoll().IsZero())
	p.Poll()
	assert.False(t, p.LastPoll().IsZero())
}
