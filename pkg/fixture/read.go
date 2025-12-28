package fixture

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Read(path ...string) string {
	f := File(path...)
	defer errors.PanicClose(f)

	return string(system.ReadAll(f))
}
