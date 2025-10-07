package preseed

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web"
)

func Download(
	release string,
	configurationPath string,
) {
	if system.FileExists(configurationPath) {
		fmt.Printf("Configuration exists: %s\n", configurationPath)

		return
	}

	l := Link(release)
	fmt.Printf("Locator: %s\n", l)
	web.Download(l, configurationPath)
}
