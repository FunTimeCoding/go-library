package coordination

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestSummarizePushesIndexer(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "summary for indexing",
		},
	)
	assert.Count(t, 1, s.Indexer.Pushed)
	assert.String(t, a.Name(), s.Indexer.Pushed[0].Name)
	assert.String(t, "summary for indexing", s.Indexer.Pushed[0].Body)
}

func TestSummarizeAmendPushesUpdatedBody(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "first draft",
		},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "revised with more detail",
		},
	)
	assert.Count(t, 2, s.Indexer.Pushed)
	assert.String(t, "revised with more detail", s.Indexer.Pushed[1].Body)
}
