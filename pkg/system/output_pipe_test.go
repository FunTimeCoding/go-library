package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"testing"
)

func TestOutputStream(t *testing.T) {
	s := run.New().Stream("echo", "hello")
	assert.String(t, "hello\n", string(ReadAll(s.Reader())))
	errors.PanicOnError(s.Wait())
}
