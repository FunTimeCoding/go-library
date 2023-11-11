package report

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestReport(t *testing.T) {
	root := New("Example root", NoLimit)
	root.String("String", "example")
	root.Integer("Integer", 2)
	firstSection := root.Nest("Example first section", NoLimit)
	firstSection.String("String", "example")
	firstSection.Float("Float without unit", 1, "")
	firstSection.Float("Float with unit", 2, "L")
	firstSection.Integer("Integer", 1)
	firstSection.Percent("Percent", 50)
	secondSection := firstSection.Nest("Example second section", NoLimit)
	secondSection.String("String", "example")
	other := New("Example other", NoLimit)
	other.String("String", "other")
	root.AppendSection(other)

	assert.String(
		t,
		"Example root"+
			"\n  String: example"+
			"\n  Integer: 2"+
			"\n  Example first section"+
			"\n    String: example"+
			"\n    Float without unit: 1.0"+
			"\n    Float with unit: 2.0 L"+
			"\n    Integer: 1"+
			"\n    Percent: 50%"+
			"\n    Example second section"+
			"\n      String: example"+
			"\n  Example other"+
			"\n    String: other",
		root.Render(),
	)
}

func TestReportLimit(t *testing.T) {
	root := New("Example root", 67)
	assert.Integer(t, 12, root.Length())
	root.String("String", "example")
	assert.Integer(t, 30, root.Length())

	firstSection := root.Nest("Example section", NoLimit)
	assert.Integer(t, 15, firstSection.Length())
	firstSection.String("String", "example")
	assert.Integer(t, 37, firstSection.Length())

	secondSection := root.Nest("Too long section", NoLimit)
	assert.Integer(t, 16, secondSection.Length())

	assert.String(
		t,
		"Example root"+
			"\n  String: example"+
			"\n  Example section"+
			"\n    String: example",
		root.Render(),
	)
}
