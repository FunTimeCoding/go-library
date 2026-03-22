package brave

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/brave/helper"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func OpenProfile(profile string) {
	run.New().Launch(
		constant.BravePath,
		fmt.Sprintf("--user-data-dir=%s", helper.SettingsPath()),
		fmt.Sprintf("--profile-directory=%s", profile),
	)
}
