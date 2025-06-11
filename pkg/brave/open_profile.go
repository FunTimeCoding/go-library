package brave

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/brave/helper"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os/exec"
)

func OpenProfile(profile string) {
	errors.PanicOnError(
		exec.Command(
			constant.BravePath,
			fmt.Sprintf("--user-data-dir=%s", helper.SettingsPath()),
			fmt.Sprintf("--profile-directory=%s", profile),
		).Start(),
	)
}
