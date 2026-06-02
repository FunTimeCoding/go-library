//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
)

func TestSeedsOrderedByPosition(t *testing.T) {
	s := store_tester.New(t)
	s.Store.UpsertSeed("charlie", "charlie.md", "hash-c", "c")
	s.Store.UpsertSeed("alpha", "alpha.md", "hash-a", "a")
	s.Store.UpsertSeed("beta", "beta.md", "hash-b", "b")
	seeds := s.Store.Seeds()
	assert.String(t, "charlie", seeds[0].Name)
	assert.String(t, "alpha", seeds[1].Name)
	assert.String(t, "beta", seeds[2].Name)
}

func TestSeedsEmptyStore(t *testing.T) {
	s := store_tester.New(t)
	seeds := s.Store.Seeds()
	assert.Count(t, 0, seeds)
}
