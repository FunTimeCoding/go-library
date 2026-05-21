//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestDeleteSession(t *testing.T) {
	s := store_tester.New(t)
	s.Store.EnsureSession("session-1")
	assert.True(t, s.Store.GetSession("session-1") != nil)
	s.Store.DeleteSession("session-1")
	assert.True(t, s.Store.GetSession("session-1") == nil)
}
