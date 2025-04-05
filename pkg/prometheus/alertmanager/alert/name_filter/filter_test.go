package name_filter

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestFilter(t *testing.T) {
	fixture := []*alert.Alert{{Name: strings.Alfa}, {Name: strings.Bravo}}

	f1 := New(false)
	f1.Keep(strings.Alfa)
	assertHasOnly(t, f1.Run(fixture), "Alfa")

	f2 := New(true)
	f2.Drop(strings.Alfa)
	assertHasOnly(t, f2.Run(fixture), "Bravo")
}

func assertHasOnly(
	t *testing.T,
	v []*alert.Alert,
	name string,
) {
	t.Helper()
	assert.Count(t, 1, v)
	assert.String(t, name, v[0].Name)
}
