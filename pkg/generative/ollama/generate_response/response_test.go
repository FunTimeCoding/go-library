package generate_response

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/ollama/ollama/api"
	"testing"
)

func TestResponse(t *testing.T) {
	assert.True(t, New(&api.GenerateResponse{}) != nil)
}
