package writer_mock

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"net/http"
	"testing"
)

func TestMock(t *testing.T) {
	mock := New()
	mock.Header().Set("a", "b")
	assert.Any(
		t,
		http.Header{"A": {"b"}},
		mock.Headers,
	)
}
