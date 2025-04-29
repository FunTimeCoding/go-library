package issue_enricher

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestEnricher(t *testing.T) {
	assert.True(t, New(nil, nil, []string{}) != nil)
}
