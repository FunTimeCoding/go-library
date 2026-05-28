//go:build local

package server

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestLogCapturesOutput(t *testing.T) {
	s := tester.New(
		t,
		"alpha: sh -c 'echo hello-from-alpha && sleep 60'\n",
		"",
	)
	time.Sleep(200 * time.Millisecond)
	log := s.Send("log", "alpha")
	assert.True(t, strings.Contains(log, "hello-from-alpha"))
}

func TestLogUnknownProcess(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\n", "")
	result := s.Send("log", "nonexistent")
	assert.True(t, strings.HasPrefix(result, "error:"))
}
