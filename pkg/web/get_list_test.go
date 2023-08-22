package web

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"net/http"
	"testing"
)

func TestGetList(t *testing.T) {
	r, e := http.NewRequest(
		http.MethodGet, "http://localhost?a=1,2,3", nil,
	)
	assert2.FatalOnError(t, e)
	assert2.Any(t, []string{"1", "2", "3"}, GetList(r, "a"))
}
