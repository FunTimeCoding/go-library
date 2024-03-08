package markup

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"sigs.k8s.io/yaml"
)

func Encode(input any) []byte {
	result, e := yaml.Marshal(input)
	errors.PanicOnError(e)

	return result
}
