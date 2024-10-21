package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"testing"
)

type ExampleRaw struct {
	String string
}

type Example struct {
	Identifier  int
	Name        string
	Description string

	Raw *ExampleRaw
}

func (e *Example) Format(s *format.Settings) string {
	f := New(s).Integer(e.Identifier).String(e.Name, e.Description)

	if s.ShowExtended {
		f.Line("  line1")
		f.Line("  line2")
	}

	f.Raw(e.Raw)

	return f.Format()
}

func NewExample(
	identifier int,
	name string,
	description string,
) *Example {
	return &Example{
		Identifier:  identifier,
		Name:        name,
		Description: description,
	}
}

func TestStatus(t *testing.T) {
	e := NewExample(1, "a", "b")
	assert.String(
		t,
		"1 | a | b\n  line1\n  line2",
		e.Format(format.New().Extended()),
	)
}

func TestStatusNested(t *testing.T) {
	o1 := NewExample(1, "a", "b")
	o2 := NewExample(2, "c", "d")
	f := format.New().Extended()

	o1Output := fmt.Sprintf("%s\n", o1.Format(f))
	assert.String(t, "1 | a | b\n  line1\n  line2\n", o1Output)

	o2Output := fmt.Sprintf("%s\n", o2.Format(f.Indent(1)))
	assert.String(t, "  2 | c | d\n    line1\n    line2\n", o2Output)

	if false {
		assert.String(
			t,
			"1 | a | b\n  line1\n  line2\n  2 | c | d\n    line1\n    line2\n",
			fmt.Sprintf("%s%s", o1Output, o2Output),
		)
	}
}
