package notation

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ToUnstructured(input []byte) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	errors.PanicOnError(u.UnmarshalJSON(input))

	return u
}
