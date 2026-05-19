//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/store_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	"testing"
)

func TestForgetMemory(t *testing.T) {
	s := store_tester.New(t)
	identifier := s.CreateMemory(
		&store.SaveOption{
			Content:     "to forget",
			Description: "desc",
			Type:        "feedback",
		},
	)
	s.ForgetMemory(identifier, "test")
	assert.Count(t, 0, s.ListMemories("", "", true))
	assert.Count(t, 1, s.ListMemories("", "", false))
}
