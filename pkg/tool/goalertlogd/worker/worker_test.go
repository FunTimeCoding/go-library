package worker

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"path/filepath"
	"testing"
	"time"
)

var testLogger = logger.New(context.Background())

func TestWorker(t *testing.T) {
	assert.NotNil(
		t,
		New(nil, nil, testLogger, 1*time.Minute, 30*24*time.Hour, nil),
	)
}

func TestPollSavesNewAlert(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			"critical",
			"Memory above 90%",
		),
	)
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	p.Poll()
	records := s.MustByName("HighMemory")
	assert.Count(t, 1, records)
	assert.String(t, "abc123", records[0].Fingerprint)
	assert.String(t, "critical", records[0].Severity)
	assert.True(t, records[0].End == nil)
}

func TestPollResolvesRemovedAlert(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			"critical",
			"Memory above 90%",
		),
	)
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	p.Poll()
	c.Remove("abc123")
	p.Poll()
	records := s.MustByName("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollIgnoresDuplicateFiring(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			"critical",
			"",
		),
	)
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	p.Poll()
	p.Poll()
	p.Poll()
	assert.Count(t, 1, s.MustByName("HighMemory"))
}

func TestRecoverStaleAdoptsFireing(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			"critical",
			"",
		),
	)
	old := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	old.Poll()
	fresh := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	fresh.RecoverStale()
	assert.Count(t, 1, s.MustUnresolved())
	fresh.Poll()
	assert.Count(t, 1, s.MustByName("HighMemory"))
	assert.True(t, s.MustByName("HighMemory")[0].End == nil)
}

func TestRecoverStaleResolvesGone(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			"critical",
			"",
		),
	)
	old := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	old.Poll()
	c.Remove("abc123")
	fresh := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	fresh.RecoverStale()
	assert.Count(t, 0, s.MustUnresolved())
	records := s.MustByName("HighMemory")
	assert.Count(t, 1, records)
	assert.True(t, records[0].End != nil)
}

func TestPollPrunesOldRecords(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			"critical",
			"",
		),
	)
	p := New(c, s, testLogger, 1*time.Minute, 0, nil)
	p.Poll()
	c.Remove("abc123")
	p.Poll()
	assert.Count(t, 0, s.MustByName("HighMemory"))
}

func TestPollTracksLastPollTime(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	defer system.Remove(path)
	s := store.New(path)
	defer s.Close()
	c := mock_client.New()
	p := New(c, s, testLogger, 1*time.Minute, 30*24*time.Hour, nil)
	assert.True(t, p.LastPoll().IsZero())
	p.Poll()
	assert.False(t, p.LastPoll().IsZero())
}
