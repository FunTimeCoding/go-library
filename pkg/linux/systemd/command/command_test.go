package command

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"systemctl",
			"list-units",
			"--output",
			"json",
			"--state",
			"failed",
		},
		Failed(),
	)
}
