//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"testing"
)

func TestRelations(t *testing.T) {
	s := store_tester.New(t)
	id1 := s.CreateMemory(
		&store.SaveOption{
			Content:     "memory one",
			Description: "first",
			Type:        "feedback",
		},
	)
	id2 := s.CreateMemory(
		&store.SaveOption{
			Content:     "memory two",
			Description: "second",
			Type:        "feedback",
		},
	)
	s.CreateRelation(id1, id2)
	assert.Count(t, 1, s.ListRelations(id1))
	assert.Count(t, 1, s.ListRelations(id2))
}
