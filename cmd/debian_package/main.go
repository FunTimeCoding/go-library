package main

import (
	"fmt"
	debianHelper "github.com/funtimecoding/go-library/pkg/debian"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/argument"
	viperHelper "github.com/funtimecoding/go-library/pkg/viper"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const MaintainerNameArgument = "maintainer-name"
const MaintainerEmailArgument = "maintainer-email"
const SystemdUnitFlag = "systemd-unit"

func main() {
	pflag.String(
		argument.Executable,
		"",
		"Executable to package: go-example",
	)
	pflag.String(
		argument.Version,
		"",
		"Package semantic version: 1.0.0, v-prefix gets trimmed",
	)
	pflag.String(
		MaintainerNameArgument,
		"",
		"Maintainer name: AN Other",
	)
	pflag.String(
		MaintainerEmailArgument,
		"",
		"Maintainer email: another@example.org",
	)
	var createUnit bool
	pflag.BoolVar(
		&createUnit,
		SystemdUnitFlag,
		false,
		"Create a systemd unit",
	)
	viperHelper.ParseAndBind()
	executable := viper.GetString(argument.Executable)
	version := semver.Trim(viper.GetString(argument.Version))
	maintainerName := viper.GetString(MaintainerNameArgument)
	maintainerMail := viper.GetString(MaintainerEmailArgument)
	architecture := system.AMD64

	packageName := debianHelper.PackageVersion(
		executable,
		version,
		1,
		architecture,
	)
	packageDirectory := system.Join(system.WorkingDirectory(), packageName)

	debianDirectory := system.Join(
		packageDirectory,
		debianHelper.PackageConfigurationDirectory,
	)
	system.MakeDirectory(debianDirectory)
	system.SaveFile(
		system.Join(debianDirectory, debianHelper.ControlFile),
		fmt.Sprintf(
			"Package: %s"+
				"\nVersion: %s"+
				"\nArchitecture: %s"+
				"\nMaintainer: %s <%s>"+
				"\nDescription: Short stub description."+
				"\n Long stub description."+
				"\n",
			executable,
			version,
			architecture,
			maintainerName,
			maintainerMail,
		),
	)

	if createUnit {
		unit := system.Join(
			packageDirectory,
			system.Library,
			"systemd",
			"system",
		)
		system.MakeDirectory(unit)
		system.SaveFile(
			system.Join(unit, fmt.Sprintf("%s.service", executable)),
			fmt.Sprintf(
				"[Unit]"+
					"\nDescription=%s stub description"+
					"\nAfter=network.target"+
					"\n"+
					"\n[Service]"+
					"\nType=simple"+
					"\nExecStart=/usr/local/bin/%s"+
					"\n"+
					"\n[Install]"+
					"\nWantedBy=multi-user.target"+
					"\n",
				executable,
				executable,
			),
		)
	}

	bin := system.Join(
		packageDirectory,
		system.Resources,
		system.Local,
		system.Binary,
	)
	system.MakeDirectory(bin)
	system.Move(executable, system.Join(bin, executable))

	fmt.Println(
		system.Run(
			"dpkg-deb",
			"--build",
			"--root-owner-group",
			packageName,
		),
	)
}
