//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"testing"
)

func TestSearchMemories(t *testing.T) {
	s := store_tester.New(t)
	s.CreateMemory(
		&store.SaveOption{
			Content:     "Always validate input at service boundaries",
			Description: "validate at boundaries; trust internal callers",
			Type:        "feedback",
			Tags:        []string{"always"},
		},
	)
	s.CreateMemory(
		&store.SaveOption{
			Content:     "Build regression tests to verify fix correctness",
			Description: "regression tests over manual verification",
			Type:        "feedback",
		},
	)
	results := s.SearchMemories("validate input", 10, "", "")
	assert.Greater(t, 0, float64(len(results)))
	assert.String(
		t,
		"validate at boundaries; trust internal callers",
		results[0].Description,
	)
}
