package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
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

func (a *Alfa) Format(f *option.Format) string {
	s := New(f).Integer(a.Identifier).String(a.Name, a.Description)

	if f.ShowExtended {
		s.Line("  line1")
		s.Line("  line2")
	}

	s.Raw(a.Raw)

	return s.Format()
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

func (b *Bravo) Format(f *option.Format) string {
	s := New(f).Integer(b.Identifier).String(b.Name)

	if f.ShowExtended {
		s.TagLine(tag.Usage, "  line1")
		s.TagLine(tag.Usage, "  line2")
	}

	s.Raw(b.Raw)

	return s.Format()
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
	apple := NewAlfa(1, "a", "b")
	assert.String(
		t,
		"1 | a | b\n  line1\n  line2",
		apple.Format(option.New().Extended()),
	)
}

func TestStatusNested(t *testing.T) {
	apple := NewAlfa(1, "a", "b")
	orange := NewAlfa(2, "c", "d")
	f := option.New().Extended()

	appleOutput := fmt.Sprintf("%s\n", apple.Format(f))
	assert.String(t, "1 | a | b\n  line1\n  line2\n", appleOutput)

	orangeOutput := fmt.Sprintf("%s\n", orange.Format(f.Indent(1)))
	assert.String(
		t,
		"  2 | c | d\n    line1\n    line2\n",
		orangeOutput,
	)

	if false {
		assert.String(
			t,
			"1 | a | b\n  line1\n  line2\n  2 | c | d\n    line1\n    line2\n",
			fmt.Sprintf("%s%s", appleOutput, orangeOutput),
		)
	}
}

func TestTagLine(t *testing.T) {
	apple := NewBravo(1, "a", "b")
	assert.String(
		t,
		"1 | a\n  line1\n  line2\n  Raw: &{String:b}",
		apple.Format(option.New().Extended().Tag(tag.Usage).Raw()),
	)
}
