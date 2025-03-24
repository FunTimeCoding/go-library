package environment

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func SetITerm(
	k string,
	v string,
) {
	errors.PanicOnError(
		exec.Command(
			"osascript",
			"-e",
			fmt.Sprintf(
				`tell application "iTerm2" to tell current session of current window to write text "export %s='%s'"`,
				k,
				v,
			),
		).Run(),
	)
}
