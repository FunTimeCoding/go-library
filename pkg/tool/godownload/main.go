package godownload

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

const (
	LatestVersion = "latest"
	DefaultOutput = ""
)

func Main() {
	common.Arguments()
	pflag.String(argument.Package, "", "Package to download")
	pflag.String(
		argument.Version,
		LatestVersion,
		"Version to download, falls back to latest if not found",
	)
	pflag.String(
		argument.Output,
		DefaultOutput,
		"Output directory for executable",
	)
	monitor.VerboseArgument()
	common.ValidateArguments()

	argument.RequiredStringFlag(argument.Package, 1)
	argument.RequiredStringFlag(argument.Version, 1)

	verbose := viper.GetBool(argument.Verbose)
	host := viper.GetString(argument.Host)
	token := viper.GetString(argument.Token)
	client := gitlab.New(host, token, []int{})
	project := common.FindProjectOrExit(
		client,
		viper.GetString(argument.Owner),
		viper.GetString(argument.Repository),
	)
	packageName := viper.GetString(argument.Package)
	version := viper.GetString(argument.Version)
	p := packages.FindVersionOrLatest(
		packages.FilterByName(
			client.Packages(project.Identifier, false),
			packageName,
		),
		version,
	)

	if p == nil {
		fmt.Printf("package not found: %s\n", packageName)

		os.Exit(1)
	}

	if version != LatestVersion && p.Version != version {
		fmt.Printf(
			"warning: version %s not found, using latest\n",
			version,
		)
	}

	f := packages.SystemMatchingFile(
		client.PackageFiles(project.Identifier, p.ID),
	)

	if f == nil {
		fmt.Printf("architecture and system match not found")

		os.Exit(1)
	}

	if !system.FileExists(packageName) {
		if !system.FileExists(f.FileName) {
			packages.Download(
				packages.Link(host, project, p, f, verbose),
				token,
				f.FileName,
			)
		}

		system.ExtractZip(f.FileName, system.WorkingDirectory())
		system.DeleteFile(f.FileName)
	} else {
		fmt.Printf("file already exists: %s\n", packageName)
	}

	system.Executable(packageName)
	output := viper.GetString(argument.Output)

	if output != DefaultOutput {
		system.Move(packageName, system.Join(output, packageName))
	}
}
