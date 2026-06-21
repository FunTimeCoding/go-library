//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store/save_option"
	"testing"
)

func TestRelations(t *testing.T) {
	s := store_tester.New(t)
	o := save_option.New()
	o.Content = "memory one"
	o.Description = "first"
	o.Type = "feedback"
	id1 := s.CreateMemory(o)
	p := save_option.New()
	p.Content = "memory two"
	p.Description = "second"
	p.Type = "feedback"
	id2 := s.CreateMemory(p)
	s.CreateRelation(id1, id2)
	assert.Count(t, 1, s.ListRelations(id1))
	assert.Count(t, 1, s.ListRelations(id2))
}
