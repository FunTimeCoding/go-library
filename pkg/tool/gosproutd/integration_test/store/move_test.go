//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
)

func threeSeeds(s *store_tester.Tester) {
	s.Store.UpsertSeed("alpha", "alpha.md", "hash-a", "a")
	s.Store.UpsertSeed("beta", "beta.md", "hash-b", "b")
	s.Store.UpsertSeed("charlie", "charlie.md", "hash-c", "c")
}

func TestMoveUpSwapsWithAbove(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveUp(seeds[1].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "beta", after[0].Name)
	assert.String(t, "alpha", after[1].Name)
	assert.String(t, "charlie", after[2].Name)
}

func TestMoveUpAtTopIsNoOp(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveUp(seeds[0].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "alpha", after[0].Name)
}

func TestMoveDownSwapsWithBelow(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveDown(seeds[0].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "beta", after[0].Name)
	assert.String(t, "alpha", after[1].Name)
	assert.String(t, "charlie", after[2].Name)
}

func TestMoveDownAtBottomIsNoOp(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.MoveDown(seeds[2].Identifier)
	after := s.Store.Seeds()
	assert.String(t, "charlie", after[2].Name)
}
