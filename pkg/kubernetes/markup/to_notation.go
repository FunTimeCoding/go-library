package markup

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func ToNotation(filePath string) []byte {
	f := system.Open(filePath)
	defer errors.LogClose(f)
	var a any
	errors.PanicOnError(yaml.NewYAMLToJSONDecoder(f).Decode(&a))

	return notation.Marshall(a)
}
