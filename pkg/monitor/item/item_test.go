package item

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/monitor/collector"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestItem(t *testing.T) {
	c := collector.New("example", "example", "examples", 0, nil)
	assert.String(
		t,
		"example-1",
		New(
			c,
			c.IntegerIdentifier(1),
			constant.Critical,
			strings.Alfa,
			"https://example.org/1",
			nil,
		).Identifier,
	)
}
