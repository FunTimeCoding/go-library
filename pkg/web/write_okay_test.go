package web

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/web/spy_writer"
	"testing"
)

func TestWriteOkay(t *testing.T) {
	w := spy_writer.New()
	WriteOkay(w, strings.Alfa)
	assert.Any(t, []byte("Alfa"), w.Written)
	assert.Integer(t, 200, w.StatusCode)
}
