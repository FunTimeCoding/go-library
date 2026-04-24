package environment

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func SetDarwin(
	k string,
	v string,
) {
	run.New().Start(
		"osascript",
		"-e",
		fmt.Sprintf(
			`tell application "iTerm2" to tell current session of current window to write text "export %s='%s'"`,
			k,
			v,
		),
	)
}
