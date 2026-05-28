//go:build local

package server

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestRestartProcess(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\n", "")
	result := s.Send("restart", "alpha")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, "*alpha", strings.TrimSpace(status))
}

func TestStopProcess(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\n", "")
	result := s.Send("stop", "alpha")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, " alpha", status)
}

func TestStartStoppedProcess(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\n", "")
	s.Send("stop", "alpha")
	time.Sleep(200 * time.Millisecond)
	result := s.Send("start", "alpha")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, "*alpha", strings.TrimSpace(status))
}
