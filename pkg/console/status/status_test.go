package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"testing"
)

type ExampleRaw struct {
	String string
}

type Alfa struct {
	Identifier  int
	Name        string
	Description string

	Raw *ExampleRaw
}

func (a *Alfa) Format(s *format.Settings) string {
	f := New(s).Integer(a.Identifier).String(a.Name, a.Description)

	if s.ShowExtended {
		f.Line("  line1")
		f.Line("  line2")
	}

	f.Raw(a.Raw)

	return f.Format()
}

func NewAlfa(
	identifier int,
	name string,
	description string,
) *Alfa {
	return &Alfa{
		Identifier:  identifier,
		Name:        name,
		Description: description,
	}
}

type Bravo struct {
	Identifier int
	Name       string

	Raw *ExampleRaw
}

func (b *Bravo) Format(s *format.Settings) string {
	f := New(s).Integer(b.Identifier).String(b.Name)

	if s.ShowExtended {
		f.TagLine(tag.Usage, "  line1")
		f.TagLine(tag.Usage, "  line2")
	}

	f.Raw(b.Raw)

	return f.Format()
}

func NewBravo(
	identifier int,
	name string,
	rawName string,
) *Bravo {
	return &Bravo{
		Identifier: identifier,
		Name:       name,
		Raw:        &ExampleRaw{String: rawName},
	}
}

func TestStatus(t *testing.T) {
	e := NewAlfa(1, "a", "b")
	assert.String(
		t,
		"1 | a | b\n  line1\n  line2",
		e.Format(format.New().Extended()),
	)
}

func TestStatusNested(t *testing.T) {
	o1 := NewAlfa(1, "a", "b")
	o2 := NewAlfa(2, "c", "d")
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

func TestTagLine(t *testing.T) {
	e := NewBravo(1, "a", "b")

	assert.String(
		t,
		"1 | a\n  line1\n  line2\n  Raw: &{String:b}",
		e.Format(format.New().Extended().Tag(tag.Usage).Raw()),
	)
}
