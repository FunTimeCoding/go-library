package item

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestItem(t *testing.T) {
	assert.String(
		t,
		"example-1",
		New(
			fmt.Sprintf("%s-%d", constant.ExamplePrefix, 1),
			constant.ErrorType,
			strings.Alfa,
			"https://example.org/1",
		).Identifier,
	)
}
