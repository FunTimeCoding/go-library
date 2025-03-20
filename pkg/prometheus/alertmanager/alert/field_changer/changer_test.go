package field_changer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"testing"
)

func TestChanger(t *testing.T) {
	a1 := New()
	a1.AddName("BigApple", "Apple")
	a1.AddName("ComplicatedOrange", "Orange")
	assert.String(t, "Apple", a1.Name("BigApple"))
	assert.String(t, "Orange", a1.Name("ComplicatedOrange"))
	assert.String(t, "Banana", a1.Name("Banana"))
	assert.Any(
		t,
		[]*alert.Alert{
			{Name: "Apple"},
			{Name: "Orange"},
		},
		a1.Run(
			[]*alert.Alert{
				{Name: "BigApple"},
				{Name: "ComplicatedOrange"},
			},
		),
	)

	a2 := New()
	a2.AddName("BigApple", "Apple")
	a2.AddSeverity(
		"Apple",
		constant.WarningSeverity,
		constant.CriticalSeverity,
	)
	assert.Any(
		t,
		[]*alert.Alert{
			{Name: "Apple", Severity: "critical"},
		},
		a2.Run(
			[]*alert.Alert{
				{Name: "BigApple", Severity: constant.WarningSeverity},
			},
		),
	)
}
