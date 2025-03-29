package result

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestResult(t *testing.T) {
	assert.Any(
		t,
		&Result{
			OutputString: "Alfa",
			ErrorString:  "Bravo",
			Exit:         1,
			Error:        nil,
		},
		New(strings.Alfa, strings.Bravo, 1, nil),
	)
}
