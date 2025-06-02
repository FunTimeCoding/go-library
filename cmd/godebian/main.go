package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/debian"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	maintainerNameArgument  = "maintainer-name"
	maintainerEmailArgument = "maintainer-email"
	systemdUnitFlag         = "systemd-unit"
)

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
		maintainerNameArgument,
		"",
		"Maintainer name: AN Other",
	)
	pflag.String(
		maintainerEmailArgument,
		"",
		"Maintainer email: another@example.org",
	)
	var createUnit bool
	pflag.BoolVar(
		&createUnit,
		systemdUnitFlag,
		false,
		"Create a systemd unit",
	)
	argument.ParseBind()
	executable := viper.GetString(argument.Executable)
	version := semver.Trim(viper.GetString(argument.Version))
	maintainerName := viper.GetString(maintainerNameArgument)
	maintainerMail := viper.GetString(maintainerEmailArgument)
	architecture := constant.AMD64

	packageName := debian.PackageVersion(
		executable,
		version,
		1,
		architecture,
	)
	packageDirectory := system.Join(system.WorkingDirectory(), packageName)

	debianDirectory := system.Join(
		packageDirectory,
		debian.PackageConfigurationDirectory,
	)
	system.MakeDirectory(debianDirectory)
	system.SaveFile(
		system.Join(debianDirectory, debian.ControlFile),
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
			constant.Library,
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
		constant.Resources,
		constant.Local,
		constant.Binary,
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
