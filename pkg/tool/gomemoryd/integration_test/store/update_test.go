//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"testing"
)

func TestUpdateCreatesVersion(t *testing.T) {
	s := store_tester.New(t)
	identifier := s.CreateMemory(
		&store.SaveOption{
			Content:     "original",
			Description: "desc",
			Type:        "feedback",
		},
	)
	s.UpdateMemory(
		identifier,
		&store.SaveOption{
			Content:     "updated content",
			Description: "new desc",
			Tags:        []string{"always"},
		},
	)
	history := s.GetMemoryHistory(identifier)
	assert.Count(t, 2, history)
	assert.String(t, "created", history[0].ChangeType)
	assert.String(t, "updated", history[1].ChangeType)
	m := s.GetMemory(identifier)
	assert.String(t, "updated content", m.Content)
	assert.Count(t, 1, m.Tags)
	assert.String(t, "always", m.Tags[0])
}
