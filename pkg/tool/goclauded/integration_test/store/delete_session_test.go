//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestDeleteSession(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.True(t, s.GetSession("session-1") != nil)
	assert.FatalOnError(t, s.Store.DeleteSession("session-1"))
	assert.True(t, s.GetSession("session-1") == nil)
}
