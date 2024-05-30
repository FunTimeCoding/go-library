package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"strings"
	"testing"
)

func TestLint(t *testing.T) {
	assert.String(
		t,
		`import "example.org/example/fmt"`,
		Lint(
			strings.NewReader(
				"package example\n\nimport (\n\t\"example.org/example/fmt\"\n)\n\nfunc Example() {\n\tfmt.Println(\"Hello friend.\")\n}\n",
			),
		),
	)
}
