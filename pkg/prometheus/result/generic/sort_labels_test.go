package generic

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"testing"
)

func TestSortLabels(t *testing.T) {
	assert.Strings(
		t,
		[]string{"namespace", "pod", "container"},
		sortLabels(
			[]string{
				constant.ContainerLabel,
				constant.PodLabel,
				constant.NamespaceLabel,
			},
		),
	)
}
