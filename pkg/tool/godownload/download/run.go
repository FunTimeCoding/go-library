package download

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/godownload/download/option"
	library "gitlab.com/gitlab-org/api/client-go"
	"os"
)

func Run(o *option.Download) {
	validate(o)
	g := gitlab.New(o.Host, o.Token)
	project := common.FindProjectOrExit(g, o.Owner, o.Repository)
	filtered := packages.FilterByName(
		g.Packages(project.Identifier, false),
		o.Package,
	)
	var p *library.Package

	if o.PackageVersion == constant.LatestVersion {
		p = packages.FindLatest(filtered)
	} else {
		p = packages.FindVersionOrLatest(filtered, o.PackageVersion)
	}

	if p == nil {
		fmt.Printf("package not found: %s\n", o.Package)

		os.Exit(1)
	}

	if o.PackageVersion != constant.LatestVersion &&
		p.Version != o.PackageVersion {
		errors.Warning(
			"version %s not found, using latest\n",
			o.PackageVersion,
		)
	}

	f := packages.SystemMatchingFile(
		g.PackageFiles(project.Identifier, p.ID),
	)

	if f == nil {
		fmt.Printf("architecture and system match not found")

		os.Exit(1)
	}

	if !system.FileExists(o.Package) {
		if !system.FileExists(f.FileName) {
			packages.Download(
				packages.Link(o.Host, project, p, f, o.Verbose),
				o.Token,
				f.FileName,
			)
		}

		system.ExtractZip(f.FileName, system.WorkingDirectory())
		system.DeleteFile(f.FileName)
	} else {
		fmt.Printf("file already exists: %s\n", o.Package)
	}

	system.Executable(o.Package)

	if o.Output != DefaultOutput {
		system.Move(o.Package, join.Absolute(o.Output, o.Package))
	}
}
