//go:build local

package model_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestListDocuments(t *testing.T) {
	s := indexFixtures(t)
	result := s.ListDocuments("test")
	assert.StringContains(t, "alpha.md", result)
	assert.StringContains(t, "tools/gamma.md", result)
	assert.StringContains(t, "archive/epsilon.md", result)
}
