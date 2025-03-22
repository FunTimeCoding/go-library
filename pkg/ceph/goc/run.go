package goc

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Run(selected string) {
	// TODO: Swap Ceph configurations like goam
	file := environment.GetDefault(
		ConfigurationEnvironment,
		"",
	)
	arguments := environment.GetDefault(
		ArgumentEnvironment,
		"",
	)
	fmt.Printf("File: %s\n", file)
	fmt.Printf("Argument: %s\n", arguments)
	base := system.Join(
		system.Home(),
		constant.ConfigurationPath,
		CephPath,
	)
	fmt.Printf("Base: %s\n", base)
	active := file
	directories := system.Directories(base)

	if selected == "" {
		for _, d := range directories {
			if d == active {
				fmt.Printf("* %s\n", d)
			} else {
				fmt.Printf("  %s\n", d)
			}
		}

		return
	}

	if false {
		environment.Set("", "")
	}
}
