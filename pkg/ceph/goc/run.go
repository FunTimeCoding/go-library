package goc

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"slices"
)

func Run(
	selected string,
	verbose bool,
) {
	file := environment.Default(ConfigurationEnvironment, "")

	if verbose {
		fmt.Printf("File: %s\n", file)
	}

	base := system.Join(system.Home(), constant.ConfigurationPath, CephPath)

	if verbose {
		fmt.Printf("Base: %s\n", base)
	}

	active := split.Slash(file)[5]

	if verbose {
		fmt.Printf("Active: %s\n", active)
	}

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

	if !slices.Contains(directories, selected) {
		fmt.Printf("Unexpected: %s\n", selected)

		return
	}

	name := configurationName(base, selected)
	newConfiguration := system.Join(base, selected, clientConfiguration)
	newArgument := fmt.Sprintf(
		"-n %s --keyring=%s",
		fmt.Sprintf("client.%s", name),
		system.Join(
			base,
			selected,
			fmt.Sprintf("ceph.client.%s.keyring", name),
		),
	)

	if verbose {
		fmt.Printf("newConfiguration: %s\n", newConfiguration)
		fmt.Printf("newArgument: %s\n", newArgument)
	}

	environment.SetITerm(ConfigurationEnvironment, newConfiguration)
	environment.SetITerm(ArgumentEnvironment, newArgument)
}
