package markup

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"sigs.k8s.io/yaml"
)

func Clean(input []byte) []byte {
	var result map[string]any
	errors.PanicOnError(yaml.Unmarshal(input, &result))

	return Encode(RemoveEmpty(result).(map[string]any))
}
