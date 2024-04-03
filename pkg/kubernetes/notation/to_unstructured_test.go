package notation

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestToUnstructured(t *testing.T) {
	assert.Any(
		t,
		&unstructured.Unstructured{Object: map[string]any{"kind": "Pod"}},
		ToUnstructured([]byte(`{"kind":"Pod"}`)),
	)
}
