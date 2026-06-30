package result

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
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
		New(upper.Alfa, upper.Bravo, 1, nil),
	)
}
