package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	assert.Any(t, &http.Client{}, Client(true))
}
