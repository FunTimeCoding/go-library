//go:build local

package server

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
)

func TestStatusShowsRunning(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\nbravo: sleep 60\n", "")
	result := s.Send("status")
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "*alpha", lines[0])
	assert.String(t, "*bravo", lines[1])
}

func TestListShowsAllProcesses(t *testing.T) {
	s := tester.New(t, "alpha: sleep 60\nbravo: sleep 60\n", "")
	result := s.Send("list")
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "alpha", lines[0])
	assert.String(t, "bravo", lines[1])
}
