package pod

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestPod(t *testing.T) {
	assert.NotNil(
		t,
		New(
			&core.Pod{ObjectMeta: meta.ObjectMeta{Name: upper.Alfa}},
			"",
		),
	)
}
