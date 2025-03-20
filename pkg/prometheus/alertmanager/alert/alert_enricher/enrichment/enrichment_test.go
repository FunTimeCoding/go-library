package enrichment

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestEnrichment(t *testing.T) {
	assert.True(t, New(strings.Alfa, strings.Bravo, strings.Charlie) != nil)
}
