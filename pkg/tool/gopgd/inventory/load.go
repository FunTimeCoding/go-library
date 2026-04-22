package inventory

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"gopkg.in/yaml.v3"
)

func Load(path string) *Inventory {
	result := &Inventory{}
	errors.PanicOnError(yaml.Unmarshal(system.ReadBytesUnsafe(path), result))

	return result
}
