//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"testing"
)

func TestListTags(t *testing.T) {
	s := store_tester.New(t)
	s.CreateMemory(
		&store.SaveOption{
			Content:     "a",
			Description: "desc a",
			Type:        "feedback",
			Tags:        []string{"always", "deployment"},
		},
	)
	s.CreateMemory(
		&store.SaveOption{
			Content:     "b",
			Description: "desc b",
			Type:        "feedback",
			Tags:        []string{"always"},
		},
	)
	tags := s.ListTags()
	assert.Count(t, 2, tags)
	assert.String(t, "always", tags[0].Tag)
	assert.Integer(t, 2, tags[0].Count)
}
