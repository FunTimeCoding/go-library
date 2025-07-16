package system

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func Move(
	from string,
	to string,
) {
	if e := os.Rename(from, to); e != nil {
		if e.Error() == "invalid cross-device link" {
			MoveCopy(from, to)
		} else {
			errors.PanicOnError(e)
		}
	}
}
