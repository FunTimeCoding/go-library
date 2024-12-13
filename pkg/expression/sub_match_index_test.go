package expression

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSubMatchIndex(t *testing.T) {
	assert.String(
		t,
		"MyPrefix: 123",
		SubMatchIndex(
			`MyPrefix: (.*)`,
			"Some line\n"+
				"MyPrefix: 123\n"+
				"Some other line\n",
			0,
		),
	)
}
