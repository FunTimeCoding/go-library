package brave

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"os/exec"
	"path/filepath"
)

func OpenProfileLink(
	link string,
	profile string,
) {
	errors.PanicOnError(
		exec.Command(
			constant.BraveMacPath,
			fmt.Sprintf(
				"--user-data-dir=%s",
				filepath.Join(system.Home(), constant.BraveMacSettings),
			),
			fmt.Sprintf("--profile-directory=%s", profile),
			"--new-window", link,
		).Start(),
	)
}
