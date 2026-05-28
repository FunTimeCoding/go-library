//go:build local

package server

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestReloadProcfileAddsNewEntry(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\n", "")
	s.WriteProcfile("alpha: sleep 60\nbravo: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	lines := strings.Split(status, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "*alpha", lines[0])
	assert.String(t, "*bravo", lines[1])
}

func TestReloadProcfileRemovesEntry(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\nbravo: sleep 60\n", "")
	s.WriteProcfile("alpha: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	list := s.Send("list")
	assert.String(t, "alpha", strings.TrimSpace(list))
}

func TestReloadProcfileChangedCommand(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\n", "")
	s.WriteProcfile("alpha: sleep 120\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	time.Sleep(500 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, "*alpha", strings.TrimSpace(status))
}
