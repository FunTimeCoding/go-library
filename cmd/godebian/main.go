package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/debian"
	debianConstant "github.com/funtimecoding/go-library/pkg/debian/constant"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
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
	packageDirectory := join.Absolute(system.WorkingDirectory(), packageName)

	debianDirectory := join.Absolute(
		packageDirectory,
		debianConstant.PackageConfigurationDirectory,
	)
	system.MakeDirectory(debianDirectory)
	system.SaveFile(
		join.Absolute(debianDirectory, debianConstant.ControlFile),
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
		unit := join.Absolute(
			packageDirectory,
			constant.Library,
			"systemd",
			"system",
		)
		system.MakeDirectory(unit)
		system.SaveFile(
			join.Absolute(unit, fmt.Sprintf("%s.service", executable)),
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

	bin := join.Absolute(
		packageDirectory,
		constant.Resources,
		constant.Local,
		constant.Binary,
	)
	system.MakeDirectory(bin)
	system.Move(executable, join.Absolute(bin, executable))

	fmt.Println(
		system.Run(
			"dpkg-deb",
			"--build",
			"--root-owner-group",
			packageName,
		),
	)
}
