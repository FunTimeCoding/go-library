package web

import (
	"github.com/funtimecoding/go-library/assert"
	"net/http"
	"testing"
)

func TestGetList(t *testing.T) {
	r, fail := http.NewRequest(
		http.MethodGet, "http://localhost?a=1,2,3", nil,
	)
	assert.FatalOnError(t, fail)
	assert.Any(t, []string{"1", "2", "3"}, GetList(r, "a"))
}