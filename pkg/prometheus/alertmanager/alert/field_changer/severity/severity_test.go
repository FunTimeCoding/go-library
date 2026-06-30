package severity

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestSeverity(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, upper.Bravo, upper.Charlie))
}
