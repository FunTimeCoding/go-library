package coordination

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/base"
	"github.com/google/uuid"
	"testing"
)

func TestCheckCreatesSession(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	assert.True(t, a.Name() != "")
}

func TestCheckPreviewNoSession(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	restClient := s.RESTClient(t)
	response, e := restClient.GetCheckWithResponse(
		context.Background(),
		&client.GetCheckParams{
			Session: uuid.New().String(),
			Preview: new(true),
		},
	)
	assert.FatalOnError(t, e)
	assert.False(t, response.JSON200.Changed)
	assert.Count(t, 0, response.JSON200.Sessions)
}

func TestCheckRecentActivity(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "build search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "indexing done",
		},
	)
	check := b.Check()
	assert.True(t, len(check.Completions) > 0)
	found := false

	for _, c := range check.Completions {
		if c.Name == a.Name() && c.Topic == "build search index" && c.Kind == constant.Complete {
			found = true
		}
	}

	assert.True(t, found)
}

func TestCheckRecentActivityLive(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "indexing done",
		},
	)
	check := b.CheckLive()
	assert.True(t, check.Changed)
	found := false

	for _, c := range check.Completions {
		if c.Name == a.Name() && c.Topic == "search index" && c.Kind == constant.Complete {
			found = true
		}
	}

	assert.True(t, found)
}

func TestCheckRecentActivityUpdate(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "initial")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "milestone reached",
		},
	)
	check := b.Check()
	found := false

	for _, c := range check.Completions {
		if c.Name == a.Name() && c.Kind == constant.Update {
			found = true
		}
	}

	assert.True(t, found)
}
