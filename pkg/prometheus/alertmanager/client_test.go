package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/web"
	"testing"
)

func TestClient(t *testing.T) {
	assert.True(
		t,
		New(web.Localhost, strings.Alfa, strings.Bravo, nil) != nil,
	)
}
