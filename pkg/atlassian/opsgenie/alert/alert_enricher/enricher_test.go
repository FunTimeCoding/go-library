package alert_enricher

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestEnricher(t *testing.T) {
	assert.NotNil(t, New())
}
