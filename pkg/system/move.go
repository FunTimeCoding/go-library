package system

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"strings"
)

func Move(
	from string,
	to string,
) {
	if e := os.Rename(from, to); e != nil {

		if strings.Contains(e.Error(), "invalid cross-device link") {
			fmt.Printf("Error: >%s<\n", e.Error())
			fmt.Printf("Try MoveCopy from %s to %s\n", from, to)
			MoveCopy(from, to)
		} else {
			fmt.Printf("Other error: >%s<\n", e.Error())
			errors.PanicOnError(e)
		}
	}
}
