package brave

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"path/filepath"
)

func OpenProfileLink(
	link string,
	profile string,
) {
	run.New().Launch(
		constant.BravePath,
		fmt.Sprintf(
			"--user-data-dir=%s",
			filepath.Join(system.Home(), constant.BraveSettings),
		),
		fmt.Sprintf("--profile-directory=%s", profile),
		"--new-window", link,
	)
}
