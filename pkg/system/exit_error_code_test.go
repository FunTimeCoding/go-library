package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"os/exec"
	"testing"
)

func TestExitErrorCode(t *testing.T) {
	assert.Integer(
		t,
		-1,
		ExitErrorCode(error(&exec.ExitError{})),
	)
}
