package message

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/query_result"
	"testing"
)

func TestMessage(t *testing.T) {
	r, v := NewSlice(query_result.New())
	assert.NotNil(t, r)
	assert.NotNil(t, v)
}
