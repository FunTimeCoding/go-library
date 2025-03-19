package alert_filter

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_parameter"
	"testing"
)

func TestFilter(t *testing.T) {
	// TODO: Test cases
	p := advanced_parameter.New()
	assert.Any(
		t,
		[]*alert.Alert{},
		New(p).Run([]*alert.Alert{}),
	)
}
