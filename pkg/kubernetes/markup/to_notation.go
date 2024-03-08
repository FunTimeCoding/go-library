package markup

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func ToNotation(filePath string) []byte {
	f := system.Open(filePath)
	defer func() {
		errors.LogOnError(f.Close())
	}()
	var a any
	errors.PanicOnError(yaml.NewYAMLToJSONDecoder(f).Decode(&a))

	return notation.Marshall(a)
}
