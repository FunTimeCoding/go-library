package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(
		t,
		New(constant.Localhost, strings.Alfa, strings.Bravo, nil),
	)
}
