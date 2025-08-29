package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestFilterLabel(t *testing.T) {
	o := fixtureOption()
	r1 := Raw("TEST-1")
	r1.Fields.Labels = []string{strings.Alfa}
	r2 := Raw("TEST-2")
	r2.Fields.Labels = []string{strings.Bravo}
	actual := FilterLabel([]*Issue{New(r1, o), New(r2, o)}, strings.Bravo)
	assert.Count(t, 1, actual)
	assert.String(t, "TEST-1", actual[0].Key)
}
