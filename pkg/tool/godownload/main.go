package godownload

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

const (
	LatestVersion = "latest"
	DefaultOutput = ""
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	common.Arguments()
	pflag.String(argument.Package, "", "Package to download")
	pflag.String(
		argument.PackageVersion,
		LatestVersion,
		"Version to download, falls back to latest if not found",
	)
	pflag.String(
		argument.Output,
		DefaultOutput,
		"Output directory for executable",
	)
	monitor.VerboseArgument()
	monitor.VersionArgument()
	argument.ParseBind()
	monitor.VersionExit(version, gitHash, buildDate)
	common.ValidateArguments()

	argument.RequiredStringFlag(argument.Package)
	argument.RequiredStringFlag(argument.PackageVersion)

	if false {
		// If both environment and argument are passed, use argument
		// Fall back to environment
		argument.RequiredPositional(0, "package")
		environment.GetDefault("CI_SERVER_FQDN", "")
		environment.GetDefault("TOKEN", "")
		environment.GetDefault("OWNER", "")
		environment.GetDefault("REPOSITORY", "")
		environment.GetDefault("OUTPUT", "")
	}

	verbose := viper.GetBool(argument.Verbose)
	host := viper.GetString(argument.Host)
	token := viper.GetString(argument.Token)
	client := gitlab.New(host, token)
	project := common.FindProjectOrExit(
		client,
		viper.GetString(argument.Owner),
		viper.GetString(argument.Repository),
	)
	packageName := viper.GetString(argument.Package)
	packageVersion := viper.GetString(argument.PackageVersion)
	p := packages.FindVersionOrLatest(
		packages.FilterByName(
			client.Packages(project.Identifier, false),
			packageName,
		),
		packageVersion,
	)

	if p == nil {
		fmt.Printf("package not found: %s\n", packageName)

		os.Exit(1)
	}

	if packageVersion != LatestVersion && p.Version != packageVersion {
		fmt.Printf(
			"warning: version %s not found, using latest\n",
			packageVersion,
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
