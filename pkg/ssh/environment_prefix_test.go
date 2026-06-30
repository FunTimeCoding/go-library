package ssh

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ssh/command"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestEnvironmentPrefix(t *testing.T) {
	assert.String(
		t,
		"",
		environmentPrefix(command.New(upper.Alfa)),
	)
}
