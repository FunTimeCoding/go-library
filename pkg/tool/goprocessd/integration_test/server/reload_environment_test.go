//go:build local

package server

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestReloadEnvironment(t *testing.T) {
	s := tester.New(
		t,
		"alpha: sh -c 'echo $TEST_VALUE && sleep 60'\n",
		"export TEST_VALUE=original\n",
	)
	time.Sleep(200 * time.Millisecond)
	log := s.Send("log", "alpha")
	assert.True(t, strings.Contains(log, "original"))
	s.WriteEnvrc("export TEST_VALUE=updated\n")
	s.Send("reload-environment")
	s.Send("restart", "alpha")
	time.Sleep(200 * time.Millisecond)
	log = s.Send("log", "alpha")
	assert.True(t, strings.Contains(log, "updated"))
}
