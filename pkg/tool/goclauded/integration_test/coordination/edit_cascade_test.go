package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
	"time"
)

func TestEditCascadeSummary(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "original summary",
		},
	)
	events := s.Store().EventsSince(
		time.Time{},
		time.Time{},
		constant.Summarize,
		1,
		0,
	)
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.Edit,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "corrected summary via edit",
		},
	)
	summaries := s.Store().ListSummaries()
	assert.Count(t, 1, summaries)
	assert.String(t, "corrected summary via edit", summaries[0].Body)
}

func TestEditCascadeSummaryPushesIndexer(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "original",
		},
	)
	before := len(s.Indexer().Pushed)
	events := s.Store().EventsSince(
		time.Time{},
		time.Time{},
		constant.Summarize,
		1,
		0,
	)
	a.MustCallTool(
		constant.Edit,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "edited via cascade",
		},
	)
	assert.Integer(t, before+1, len(s.Indexer().Pushed))
	last := s.Indexer().Pushed[len(s.Indexer().Pushed)-1]
	assert.String(t, "edited via cascade", last.Body)
}

func TestEditCascadeCompletion(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "hasty completion",
		},
	)
	events := s.Store().EventsSince(
		time.Time{},
		time.Time{},
		constant.Complete,
		1,
		0,
	)
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.Edit,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "corrected completion with details",
		},
	)
	completions := s.Store().RecentCompletions()
	var found bool

	for _, c := range completions {
		if c.Topic == "search index" && c.Kind == constant.Complete {
			assert.String(t, "corrected completion with details", c.Summary)
			found = true
		}
	}

	assert.True(t, found)
}

func TestEditMomentNoCascade(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Moment,
		map[string]any{
			constant.Line: "something landed",
		},
	)
	before := len(s.Indexer().Pushed)
	events := s.Store().EventsSince(
		time.Time{},
		time.Time{},
		constant.Moment,
		1,
		0,
	)
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.Edit,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "revised moment",
		},
	)
	assert.Integer(t, before, len(s.Indexer().Pushed))
	edited := s.Store().EventsSince(
		time.Time{},
		time.Time{},
		constant.Moment,
		1,
		0,
	)
	assert.String(t, "revised moment", edited[0].Body)
}
