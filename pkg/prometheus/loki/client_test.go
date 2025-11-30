package loki

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(
		t,
		New(strings.Alfa, strings.Bravo, strings.Charlie, false),
	)
}
