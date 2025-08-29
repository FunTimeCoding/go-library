package source

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/monitor/helper"
	"testing"
)

func TestSource(t *testing.T) {
	assert.NotNil(
		t,
		New(
			0,
			0,
			0,
			0,
			0,
			helper.SeverityWeights(0, 0, 0),
		),
	)
}
