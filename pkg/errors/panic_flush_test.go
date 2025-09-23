package errors

import (
	"github.com/funtimecoding/go-library/pkg/errors/mock_flusher"
	"testing"
)

func TestPanicFlush(t *testing.T) {
	PanicFlush(mock_flusher.New())
}
