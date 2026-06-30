package enrichment

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestEnrichment(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, upper.Bravo, upper.Charlie))
}
