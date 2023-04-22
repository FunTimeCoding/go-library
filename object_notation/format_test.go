package object_notation

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFormatTypes(t *testing.T) {
	assert.Any(
		t,
		"{"+
			"\n    \"String\": \"a\","+
			"\n    \"Integer\": 1,"+
			"\n    \"Float\": 1.5,"+
			"\n    \"Boolean\": true"+
			"\n}",
		Format(
			TypesFixture{
				String:  "a",
				Integer: 1,
				Float:   1.5,
				Boolean: true,
			},
		),
	)
}

func TestFormatStringWithVector(t *testing.T) {
	assert.Any(
		t,
		"{"+
			"\n    \"String\": \"1,<1.0, 1.0, 1.0>,2\""+
			"\n}",
		Format(
			StringFixture{
				String: "1,<1.0, 1.0, 1.0>,2",
			},
		),
	)
}

type StringFixture struct {
	String string
}

type TypesFixture struct {
	String  string
	Integer int
	Float   float64
	Boolean bool
}
