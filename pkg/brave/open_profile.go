package brave

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"os/exec"
	"path/filepath"
)

func OpenProfile(profile string) {
	errors.PanicOnError(
		exec.Command(
			constant.BravePath,
			fmt.Sprintf(
				"--user-data-dir=%s",
				filepath.Join(system.Home(), constant.BraveSettings),
			),
			fmt.Sprintf("--profile-directory=%s", profile),
		).Start(),
	)
}
