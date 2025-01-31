package ssh

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ssh/command"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestEnvironmentPrefix(t *testing.T) {
	assert.String(
		t,
		"",
		environmentPrefix(command.New(strings.Alfa)),
	)
}
