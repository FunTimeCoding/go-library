//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"testing"
)

func TestCreateAndGetMemory(t *testing.T) {
	s := store_tester.New(t)
	identifier := s.CreateMemory(
		&store.SaveOption{
			Content:     "Retry failed requests with exponential backoff.",
			Description: "retry with backoff; each caller gets its own policy",
			Type:        "feedback",
			Tags:        []string{"always", "http-client"},
		},
	)
	m := s.GetMemory(identifier)
	assert.String(
		t,
		"Retry failed requests with exponential backoff.",
		m.Content,
	)
	assert.Count(t, 2, m.Tags)
}
