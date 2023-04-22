package web

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-library/web/writer_mock"
	"net/http"
	"testing"
)

func TestTextHeader(t *testing.T) {
	mock := writer_mock.New()
	TextHeader(mock)
	assert.Any(
		t,
		http.Header{"Content-Type": {"text/plain"}},
		mock.Headers,
	)
}

func TestObjectHeader(t *testing.T) {
	mock := writer_mock.New()
	ObjectHeader(mock)
	assert.Any(
		t,
		http.Header{"Content-Type": {"application/json"}},
		mock.Headers,
	)
}
