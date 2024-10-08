package status

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"testing"
)

type ExampleRaw struct {
	String string
}

type Example struct {
	Raw *ExampleRaw
}

func (e *Example) Format(s *format.Settings) string {
	f := New(s).String("a", "b")

	if s.ShowExtended {
		f.Line("  extended")
	}

	f.Raw(e.Raw)

	return f.Format()
}

func TestStatus(t *testing.T) {
	e := &Example{}
	f := format.New().Extended()

	assert.String(
		t,
		"a | b\n  extended",
		e.Format(f),
	)

	// TODO: Test with nested format and indentation
}
