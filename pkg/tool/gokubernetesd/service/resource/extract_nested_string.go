package resource

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ExtractNestedString(
	object map[string]any,
	fields ...string,
) string {
	v, _, e := unstructured.NestedString(object, fields...)
	errors.PanicOnError(e)

	return v
}
