package go_mod

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestReplaceReplaces(t *testing.T) {
	assert.String(
		t,
		"replace (\nb\n)\n",
		ReplaceReplaces(
			"replace (\na\n)\n",
			"replace (\nb\n)\n",
		),
	)
}
