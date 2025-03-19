package name_aliaser

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"testing"
)

func TestAliaser(t *testing.T) {
	a := New()
	a.Add("BigApple", "Apple")
	a.Add("ComplicatedOrange", "Orange")
	assert.String(t, "Apple", a.Alias("BigApple"))
	assert.String(t, "Orange", a.Alias("ComplicatedOrange"))
	assert.String(t, "Banana", a.Alias("Banana"))
	assert.Any(
		t,
		[]*alert.Alert{
			{Name: "Apple"},
			{Name: "Orange"},
		},
		a.Run(
			[]*alert.Alert{
				{Name: "BigApple"},
				{Name: "ComplicatedOrange"},
			},
		),
	)
}
