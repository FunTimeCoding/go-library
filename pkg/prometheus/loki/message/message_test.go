package message

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"
	"testing"
)

func TestMessage(t *testing.T) {
	r, v := NewSlice(response.Data{})
	assert.NotNil(t, r)
	assert.NotNil(t, v)
}
